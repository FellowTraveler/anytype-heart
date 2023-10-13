package editor

import (
	"context"
	"fmt"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/commonspace/object/tree/objecttree"
	"github.com/gogo/protobuf/types"

	"github.com/anyproto/anytype-heart/core/anytype/config"
	"github.com/anyproto/anytype-heart/core/block/editor/bookmark"
	"github.com/anyproto/anytype-heart/core/block/editor/converter"
	"github.com/anyproto/anytype-heart/core/block/editor/file"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/getblock"
	"github.com/anyproto/anytype-heart/core/block/migration"
	"github.com/anyproto/anytype-heart/core/block/restriction"
	"github.com/anyproto/anytype-heart/core/block/source"
	"github.com/anyproto/anytype-heart/core/event"
	"github.com/anyproto/anytype-heart/core/files"
	"github.com/anyproto/anytype-heart/pkg/lib/core"
	coresb "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
)

var log = logging.Logger("anytype-mw-editor")

type personalIDProvider interface {
	PersonalSpaceID() string
}

type bundledObjectsInstaller interface {
	InstallBundledObjects(ctx context.Context, spaceID string, ids []string) ([]string, []*types.Struct, error)
}

type ObjectFactory struct {
	anytype            core.Service
	bookmarkService    bookmark.BookmarkService
	fileBlockService   file.BlockService
	layoutConverter    converter.LayoutConverter
	objectStore        objectstore.ObjectStore
	sourceService      source.Service
	tempDirProvider    core.TempDirProvider
	fileService        files.Service
	config             *config.Config
	picker             getblock.ObjectGetter
	eventSender        event.Sender
	restrictionService restriction.Service
	indexer            smartblock.Indexer
	spaceService       spaceService
	objectDeriver      objectDeriver
}

func NewObjectFactory() *ObjectFactory {
	return &ObjectFactory{}
}

func (f *ObjectFactory) Init(a *app.App) (err error) {
	f.anytype = app.MustComponent[core.Service](a)
	f.bookmarkService = app.MustComponent[bookmark.BookmarkService](a)
	f.fileBlockService = app.MustComponent[file.BlockService](a)
	f.objectStore = app.MustComponent[objectstore.ObjectStore](a)
	f.restrictionService = app.MustComponent[restriction.Service](a)
	f.sourceService = app.MustComponent[source.Service](a)
	f.fileService = app.MustComponent[files.Service](a)
	f.config = app.MustComponent[*config.Config](a)
	f.tempDirProvider = app.MustComponent[core.TempDirProvider](a)
	f.layoutConverter = app.MustComponent[converter.LayoutConverter](a)
	f.picker = app.MustComponent[getblock.ObjectGetter](a)
	f.indexer = app.MustComponent[smartblock.Indexer](a)
	f.eventSender = app.MustComponent[event.Sender](a)
	f.objectDeriver = app.MustComponent[objectDeriver](a)
	f.spaceService = app.MustComponent[spaceService](a)

	return nil
}

const CName = "objectFactory"

func (f *ObjectFactory) Name() (name string) {
	return CName
}

func (f *ObjectFactory) InitObject(space smartblock.Space, id string, initCtx *smartblock.InitContext) (sb smartblock.SmartBlock, err error) {
	sc, err := f.sourceService.NewSource(initCtx.Ctx, space, id, initCtx.BuildOpts)
	if err != nil {
		return
	}

	var ot objecttree.ObjectTree
	if p, ok := sc.(source.ObjectTreeProvider); ok {
		ot = p.Tree()
	}
	defer func() {
		if err != nil && ot != nil {
			ot.Close()
		}
	}()

	sb, err = f.New(space, sc.Type())
	if err != nil {
		return nil, fmt.Errorf("new smartblock: %w", err)
	}

	if ot != nil {
		// using lock from object tree
		sb.SetLocker(ot)
	}

	// we probably don't need any locks here, because the object is initialized synchronously
	initCtx.Source = sc
	err = sb.Init(initCtx)
	if err != nil {
		return nil, fmt.Errorf("init smartblock: %w", err)
	}

	migration.RunMigrations(sb, initCtx)
	return sb, sb.Apply(initCtx.State, smartblock.NoHistory, smartblock.NoEvent, smartblock.NoRestrictions, smartblock.SkipIfNoChanges, smartblock.KeepInternalFlags)
}

func (f *ObjectFactory) produceSmartblock(space smartblock.Space) smartblock.SmartBlock {
	return smartblock.New(
		space,
		f.fileService,
		f.restrictionService,
		f.objectStore,
		f.indexer,
		f.eventSender,
	)
}

func (f *ObjectFactory) New(space smartblock.Space, sbType coresb.SmartBlockType) (smartblock.SmartBlock, error) {
	sb := f.produceSmartblock(space)
	switch sbType {
	case coresb.SmartBlockTypePage,
		coresb.SmartBlockTypeDate,
		coresb.SmartBlockTypeBundledRelation,
		coresb.SmartBlockTypeBundledObjectType,
		coresb.SmartBlockTypeObjectType,
		coresb.SmartBlockTypeRelation,
		coresb.SmartBlockTypeRelationOption:
		return f.newPage(sb), nil
	case coresb.SmartBlockTypeArchive:
		return NewArchive(sb, f.objectStore), nil
	case coresb.SmartBlockTypeHome:
		return NewDashboard(sb, f.objectStore, f.anytype, f.layoutConverter), nil
	case coresb.SmartBlockTypeProfilePage,
		coresb.SmartBlockTypeAnytypeProfile:
		return NewProfile(sb, f.objectStore, f.fileBlockService, f.anytype, f.picker, f.bookmarkService, f.tempDirProvider, f.layoutConverter, f.fileService, f.eventSender), nil
	case coresb.SmartBlockTypeFile:
		return NewFiles(sb), nil
	case coresb.SmartBlockTypeTemplate,
		coresb.SmartBlockTypeBundledTemplate:
		return f.newTemplate(sb), nil
	case coresb.SmartBlockTypeWorkspace:
		return NewWorkspace(sb, f.objectStore, f.spaceService, f.layoutConverter, f.config, f.eventSender, f.objectDeriver), nil
	case coresb.SmartBlockTypeSpaceView:
		return newSpaceView(
			sb,
			f.spaceService,
		), nil
	case coresb.SmartBlockTypeMissingObject:
		return NewMissingObject(sb), nil
	case coresb.SmartBlockTypeWidget:
		return NewWidgetObject(sb, f.objectStore, f.layoutConverter), nil
	case coresb.SmartBlockTypeSubObject:
		return nil, fmt.Errorf("subobject not supported via factory")
	default:
		return nil, fmt.Errorf("unexpected smartblock type: %v", sbType)
	}
}
