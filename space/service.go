package space

import (
	"context"
	"sync"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/app/logger"
	"github.com/gogo/protobuf/types"

	"github.com/anyproto/anytype-heart/core/anytype/config"
	"github.com/anyproto/anytype-heart/core/block/object/objectcache"
	coresb "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/threads"
	"github.com/anyproto/anytype-heart/space/objectprovider"
	"github.com/anyproto/anytype-heart/space/spacecore"
	"github.com/anyproto/anytype-heart/space/spaceinfo"
)

const CName = "client.space"

var log = logger.NewNamed(CName)

func New() SpaceService {
	return &service{}
}

type spaceIndexer interface {
	ReindexCommonObjects() error
	ReindexSpace(spaceID string) error
}

type bundledObjectsInstaller interface {
	InstallBundledObjects(ctx context.Context, spaceID string, ids []string) ([]string, []*types.Struct, error)
}

type SpaceService interface {
	Create(ctx context.Context) (space Space, err error)
	Get(ctx context.Context, id string) (space Space, err error)

	DerivedIDs(ctx context.Context, spaceID string) (ids threads.DerivedSmartblockIds, err error)

	app.ComponentRunnable
}

type service struct {
	indexer     spaceIndexer
	spaceCore   spacecore.SpaceCoreService
	provider    objectprovider.ObjectProvider
	objectCache objectcache.Cache

	techSpace *techSpace

	personalSpaceID string

	newAccount bool

	loading map[string]*loadingSpace
	loaded  map[string]Space
	mu      sync.Mutex

	ctx       context.Context
	ctxCancel context.CancelFunc

	repKey uint64
}

func (s *service) Init(a *app.App) (err error) {
	s.indexer = app.MustComponent[spaceIndexer](a)
	s.spaceCore = app.MustComponent[spacecore.SpaceCoreService](a)
	s.objectCache = app.MustComponent[objectcache.Cache](a)
	installer := app.MustComponent[bundledObjectsInstaller](a)
	s.provider = objectprovider.NewObjectProvider(s.objectCache, installer)
	s.newAccount = a.MustComponent(config.CName).(*config.Config).NewAccount
	s.loading = map[string]*loadingSpace{}
	s.loaded = map[string]Space{}
	return nil
}

func (s *service) Name() (name string) {
	return CName
}

func (s *service) Run(ctx context.Context) (err error) {
	s.ctx, s.ctxCancel = context.WithCancel(context.Background())

	s.personalSpaceID, err = s.spaceCore.DeriveID(ctx, spacecore.SpaceType)
	if err != nil {
		return
	}

	techSpaceCore, err := s.spaceCore.Derive(ctx, spacecore.TechSpaceType)
	if err != nil {
		return
	}
	s.techSpace = newTechSpace(s, techSpaceCore)

	err = s.indexer.ReindexCommonObjects()
	if err != nil {
		return
	}

	if s.newAccount {
		return s.createPersonalSpace(ctx)
	}
	return s.loadPersonalSpace(ctx)
}

func (s *service) Create(ctx context.Context) (Space, error) {
	coreSpace, err := s.spaceCore.Create(ctx, s.repKey)
	if err != nil {
		return nil, err
	}
	return s.create(ctx, coreSpace)
}

func (s *service) Get(ctx context.Context, spaceID string) (sp Space, err error) {
	return s.waitLoad(ctx, spaceID)
}

func (s *service) open(ctx context.Context, spaceID string) (sp Space, err error) {
	coreSpace, err := s.spaceCore.Get(ctx, spaceID)
	if err != nil {
		return nil, err
	}
	sp = newSpace(s, coreSpace)
	return
}

func (s *service) createPersonalSpace(ctx context.Context) (err error) {
	coreSpace, err := s.spaceCore.Derive(ctx, spacecore.SpaceType)
	if err != nil {
		return
	}
	_, err = s.create(ctx, coreSpace)
	return
}

func (s *service) loadPersonalSpace(ctx context.Context) (err error) {
	if err = s.startLoad(ctx, s.personalSpaceID); err != nil {
		return
	}
	_, err = s.waitLoad(ctx, s.personalSpaceID)
	return err
}

func (s *service) IsPersonal(id string) bool {
	return s.personalSpaceID == id
}

func (s *service) DerivedIDs(ctx context.Context, spaceID string) (ids threads.DerivedSmartblockIds, err error) {
	var sbTypes []coresb.SmartBlockType
	if s.IsPersonal(spaceID) {
		sbTypes = threads.PersonalSpaceTypes
	} else {
		sbTypes = threads.SpaceTypes
	}
	return s.provider.DeriveObjectIDs(ctx, spaceID, sbTypes)
}

func (s *service) OnViewCreated(ctx context.Context, spaceID string) (info spaceinfo.SpaceInfo, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	info, _, err = s.createLoaderOrReturnInfo(ctx, spaceID)
	return
}

func (s *service) Close(ctx context.Context) (err error) {
	return nil
}
