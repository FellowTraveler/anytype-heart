package builtinobjects

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/anyproto/any-sync/app"

	"github.com/anyproto/anytype-heart/core/block"
	importer "github.com/anyproto/anytype-heart/core/block/import"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/system_object"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/core"
	"github.com/anyproto/anytype-heart/pkg/lib/database"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/constant"
	"github.com/anyproto/anytype-heart/util/pbtypes"

	_ "embed"
)

const (
	CName            = "builtinobjects"
	injectionTimeout = 30 * time.Second

	migrationUseCase       = -1
	migrationDashboardName = "bafyreiha2hjbrzmwo7rpiiechv45vv37d6g5aezyr5wihj3agwawu6zi3u"
)

//go:embed data/skip.zip
var skipZip []byte

//go:embed data/personal_projects.zip
var personalProjectsZip []byte

//go:embed data/knowledge_base.zip
var knowledgeBaseZip []byte

//go:embed data/notes_diary.zip
var notesDiaryZip []byte

//go:embed data/migration_dashboard.zip
var migrationDashboardZip []byte

//go:embed data/strategic_writing.zip
var strategicWritingZip []byte

var (
	log = logging.Logger("anytype-mw-builtinobjects")

	archives = map[pb.RpcObjectImportUseCaseRequestUseCase][]byte{
		pb.RpcObjectImportUseCaseRequest_SKIP:              skipZip,
		pb.RpcObjectImportUseCaseRequest_PERSONAL_PROJECTS: personalProjectsZip,
		pb.RpcObjectImportUseCaseRequest_KNOWLEDGE_BASE:    knowledgeBaseZip,
		pb.RpcObjectImportUseCaseRequest_NOTES_DIARY:       notesDiaryZip,
		pb.RpcObjectImportUseCaseRequest_STRATEGIC_WRITING: strategicWritingZip,
	}
)

type BuiltinObjects interface {
	app.Component

	CreateObjectsForUseCase(ctx context.Context, spaceID string, req pb.RpcObjectImportUseCaseRequestUseCase) (code pb.RpcObjectImportUseCaseResponseErrorCode, err error)
	InjectMigrationDashboard(spaceID string) error
}

type builtinObjects struct {
	service             *block.Service
	coreService         core.Service
	importer            importer.Importer
	store               objectstore.ObjectStore
	tempDirService      core.TempDirProvider
	systemObjectService system_object.Service
}

func New() BuiltinObjects {
	return &builtinObjects{}
}

func (b *builtinObjects) Init(a *app.App) (err error) {
	b.service = a.MustComponent(block.CName).(*block.Service)
	b.coreService = a.MustComponent(core.CName).(core.Service)
	b.importer = a.MustComponent(importer.CName).(importer.Importer)
	b.store = app.MustComponent[objectstore.ObjectStore](a)
	b.tempDirService = app.MustComponent[core.TempDirProvider](a)
	b.systemObjectService = app.MustComponent[system_object.Service](a)
	return
}

func (b *builtinObjects) Name() (name string) {
	return CName
}

func (b *builtinObjects) CreateObjectsForUseCase(
	ctx context.Context,
	spaceID string,
	useCase pb.RpcObjectImportUseCaseRequestUseCase,
) (code pb.RpcObjectImportUseCaseResponseErrorCode, err error) {
	start := time.Now()

	archive, found := archives[useCase]
	if !found {
		return pb.RpcObjectImportUseCaseResponseError_BAD_INPUT,
			fmt.Errorf("failed to import builtinObjects: invalid Use Case value: %v", useCase)
	}

	if err = b.inject(ctx, spaceID, useCase, archive); err != nil {
		return pb.RpcObjectImportUseCaseResponseError_UNKNOWN_ERROR,
			fmt.Errorf("failed to import builtinObjects for Use Case %s: %s",
				pb.RpcObjectImportUseCaseRequestUseCase_name[int32(useCase)], err.Error())
	}

	spent := time.Now().Sub(start)
	if spent > injectionTimeout {
		log.Debugf("built-in objects injection time exceeded timeout of %s and is %s", injectionTimeout.String(), spent.String())
	}

	return pb.RpcObjectImportUseCaseResponseError_NULL, nil
}

func (b *builtinObjects) InjectMigrationDashboard(spaceID string) error {
	return b.inject(context.Background(), spaceID, migrationUseCase, migrationDashboardZip)
}

func (b *builtinObjects) inject(ctx context.Context, spaceID string, useCase pb.RpcObjectImportUseCaseRequestUseCase, archive []byte) (err error) {
	path := filepath.Join(b.tempDirService.TempDir(), time.Now().Format("tmp.20060102.150405.99")+".zip")
	if err = os.WriteFile(path, archive, 0644); err != nil {
		return fmt.Errorf("failed to save use case archive to temporary file: %s", err.Error())
	}

	if err = b.importArchive(ctx, spaceID, path); err != nil {
		return err
	}

	// TODO: GO-1387 Need to use profile.pb to handle dashboard injection during migration
	oldId := migrationDashboardName
	if useCase != migrationUseCase {
		oldId, err = b.getOldSpaceDashboardId(archive)
		if err != nil {
			log.Errorf("Failed to get old id of space dashboard object: %s", err.Error())
			return nil
		}
	}

	newId, err := b.getNewSpaceDashboardId(spaceID, oldId)
	if err != nil {
		log.Errorf("Failed to get new id of space dashboard object: %s", err.Error())
		return nil
	}

	b.handleSpaceDashboard(ctx, spaceID, newId)

	if useCase != pb.RpcObjectImportUseCaseRequest_SKIP {
		b.createNotesAndTaskTrackerWidgets(ctx, spaceID)
	}
	return
}

func (b *builtinObjects) importArchive(ctx context.Context, spaceID string, path string) (err error) {
	if err = b.importer.Import(ctx, &pb.RpcObjectImportRequest{
		SpaceId:               spaceID,
		UpdateExistingObjects: false,
		Type:                  pb.RpcObjectImportRequest_Pb,
		Mode:                  pb.RpcObjectImportRequest_ALL_OR_NOTHING,
		NoProgress:            true,
		IsMigration:           false,
		Params: &pb.RpcObjectImportRequestParamsOfPbParams{
			PbParams: &pb.RpcObjectImportRequestPbParams{
				Path:         []string{path},
				NoCollection: true,
			}},
	}); err != nil {
		return err
	}

	if err = os.Remove(path); err != nil {
		log.Errorf("failed to remove temporary file %s: %s", path, err.Error())
	}

	return nil
}

func (b *builtinObjects) getOldSpaceDashboardId(archive []byte) (id string, err error) {
	var (
		rd      io.ReadCloser
		openErr error
	)
	zr, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))
	if err != nil {
		return "", err
	}
	profileFound := false
	for _, zf := range zr.File {
		if zf.Name == constant.ProfileFile {
			profileFound = true
			rd, openErr = zf.Open()
			if openErr != nil {
				return "", openErr
			}
			break
		}
	}

	if !profileFound {
		return "", fmt.Errorf("no profile file included in archive")
	}

	defer rd.Close()
	data, err := io.ReadAll(rd)

	profile := &pb.Profile{}
	if err = profile.Unmarshal(data); err != nil {
		return "", err
	}
	return profile.SpaceDashboardId, nil
}

func (b *builtinObjects) getNewSpaceDashboardId(spaceID string, oldID string) (id string, err error) {
	ids, _, err := b.store.QueryObjectIDs(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				Condition:   model.BlockContentDataviewFilter_Equal,
				RelationKey: bundle.RelationKeyOldAnytypeID.String(),
				Value:       pbtypes.String(oldID),
			},
			{
				Condition:   model.BlockContentDataviewFilter_Equal,
				RelationKey: bundle.RelationKeySpaceId.String(),
				Value:       pbtypes.String(spaceID),
			},
		},
	}, nil)
	if err == nil && len(ids) > 0 {
		return ids[0], nil
	}
	return "", err
}

func (b *builtinObjects) handleSpaceDashboard(ctx context.Context, spaceID string, id string) {
	if err := b.service.SetDetails(nil, pb.RpcObjectSetDetailsRequest{
		ContextId: b.coreService.PredefinedObjects(spaceID).Workspace,
		Details: []*pb.RpcObjectSetDetailsDetail{
			{
				Key:   bundle.RelationKeySpaceDashboardId.String(),
				Value: pbtypes.String(id),
			},
		},
	}); err != nil {
		log.Errorf("Failed to set SpaceDashboardId relation to Account object: %s", err.Error())
	}
	b.createSpaceDashboardWidget(ctx, spaceID, id)
}

func (b *builtinObjects) createSpaceDashboardWidget(ctx context.Context, spaceID string, id string) {
	targetID, err := b.getWidgetBlockIdByNumber(spaceID, 0)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	if _, err = b.service.CreateWidgetBlock(nil, &pb.RpcBlockCreateWidgetRequest{
		ContextId:    b.coreService.PredefinedObjects(spaceID).Widgets,
		TargetId:     targetID,
		Position:     model.Block_Top,
		WidgetLayout: model.BlockContentWidget_Link,
		Block: &model.Block{
			Id:          "",
			ChildrenIds: nil,
			Content: &model.BlockContentOfLink{
				Link: &model.BlockContentLink{
					TargetBlockId: id,
					Style:         model.BlockContentLink_Page,
					IconSize:      model.BlockContentLink_SizeNone,
					CardStyle:     model.BlockContentLink_Inline,
					Description:   model.BlockContentLink_None,
				},
			},
		},
	}); err != nil {
		log.Errorf("Failed to link SpaceDashboard to Widget object: %s", err.Error())
	}
}

func (b *builtinObjects) createNotesAndTaskTrackerWidgets(ctx context.Context, spaceID string) {
	targetID, err := b.getWidgetBlockIdByNumber(spaceID, 1)
	if err != nil {
		log.Errorf("Failed to get id of second widget block: %s", err.Error())
		return
	}
	for _, objectTypeKey := range []bundle.TypeKey{bundle.TypeKeyNote, bundle.TypeKeyTask} {
		id, err := b.getSetIDByObjectTypeKey(spaceID, objectTypeKey)
		if err != nil {
			log.Errorf("Failed to get id of set by '%s' to create widget object: %s", objectTypeKey, err.Error())
			continue
		}
		if _, err = b.service.CreateWidgetBlock(nil, &pb.RpcBlockCreateWidgetRequest{
			ContextId:    b.coreService.PredefinedObjects(spaceID).Widgets,
			TargetId:     targetID,
			Position:     model.Block_Bottom,
			WidgetLayout: model.BlockContentWidget_CompactList,
			Block: &model.Block{
				Id:          "",
				ChildrenIds: nil,
				Content: &model.BlockContentOfLink{
					Link: &model.BlockContentLink{
						TargetBlockId: id,
						Style:         model.BlockContentLink_Page,
						IconSize:      model.BlockContentLink_SizeNone,
						CardStyle:     model.BlockContentLink_Inline,
						Description:   model.BlockContentLink_None,
					},
				},
			},
		}); err != nil {
			log.Errorf("Failed to make Widget block for set by '%s': %s", objectTypeKey, err.Error())
		}
	}
}

func (b *builtinObjects) getSetIDByObjectTypeKey(spaceID string, objectTypeKey bundle.TypeKey) (string, error) {
	objectTypeID, err := b.systemObjectService.GetTypeIdByKey(context.Background(), spaceID, objectTypeKey)
	if err != nil {
		return "", fmt.Errorf("get type id by key '%s': %s", objectTypeKey, err)
	}
	ids, _, err := b.store.QueryObjectIDs(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				Condition:   model.BlockContentDataviewFilter_Equal,
				RelationKey: bundle.RelationKeySetOf.String(),
				Value:       pbtypes.StringList([]string{objectTypeID}),
			},
		},
	}, nil)
	if err == nil && len(ids) > 0 {
		return ids[0], nil
	}
	if len(ids) == 0 {
		err = fmt.Errorf("no object found")
	}
	return "", err
}

func (b *builtinObjects) getWidgetBlockIdByNumber(spaceID string, index int) (string, error) {
	w, err := b.service.GetObject(context.Background(), domain.FullID{
		SpaceID:  spaceID,
		ObjectID: b.coreService.PredefinedObjects(spaceID).Widgets,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get Widget object: %s", err.Error())
	}
	root := w.Pick(w.RootId())
	if root == nil {
		return "", fmt.Errorf("failed to pick root block of Widget object")
	}
	if len(root.Model().ChildrenIds) < index+1 {
		return "", fmt.Errorf("failed to get %d block of Widget object as there olny %d of them", index+1, len(root.Model().ChildrenIds))
	}
	target := w.Pick(root.Model().ChildrenIds[index])
	if target == nil {
		return "", fmt.Errorf("failed to get id of first block of Widget object: %s", err.Error())
	}
	return target.Model().Id, nil
}
