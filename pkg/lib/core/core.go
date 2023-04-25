package core

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/anytypeio/any-sync/app"
	"github.com/anytypeio/any-sync/commonfile/fileservice"
	"github.com/anytypeio/any-sync/commonspace/object/treegetter"
	"github.com/libp2p/go-libp2p/core/peer"
	"go.uber.org/zap"

	"github.com/anytypeio/go-anytype-middleware/core/anytype/config"
	"github.com/anytypeio/go-anytype-middleware/core/configfetcher"
	"github.com/anytypeio/go-anytype-middleware/core/event"
	"github.com/anytypeio/go-anytype-middleware/core/wallet"
	"github.com/anytypeio/go-anytype-middleware/metrics"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/cafe"
	coresb "github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/datastore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/files"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/addr"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/filestore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/threads"
	"github.com/anytypeio/go-anytype-middleware/space"
)

var log = logging.Logger("anytype-core")

var ErrObjectDoesNotBelongToWorkspace = fmt.Errorf("object does not belong to workspace")

const (
	CName  = "anytype"
	tmpDir = "tmp"
)

type Service interface {
	Account() string // deprecated, use wallet component
	Device() string  // deprecated, use wallet component
	Stop() error
	IsStarted() bool
	SpaceService() space.Service

	EnsurePredefinedBlocks(ctx context.Context) error
	PredefinedBlocks() threads.DerivedSmartblockIds

	// FileOffload removes file blocks recursively, but leave details
	FileOffload(id string) (bytesRemoved uint64, err error)

	FileByHash(ctx context.Context, hash string) (File, error)
	FileAdd(ctx context.Context, opts ...files.AddOption) (File, error)
	FileGetKeys(hash string) (*files.FileKeys, error)
	FileStoreKeys(fileKeys ...files.FileKeys) error

	ImageByHash(ctx context.Context, hash string) (Image, error)
	ImageAdd(ctx context.Context, opts ...files.AddOption) (Image, error)

	GetAllWorkspaces() ([]string, error)
	GetWorkspaceIdForObject(objectId string) (string, error)

	ObjectStore() objectstore.ObjectStore // deprecated
	FileStore() filestore.FileStore       // deprecated
	ThreadsIds() ([]string, error)        // deprecated

	ObjectInfoWithLinks(id string) (*model.ObjectInfoWithLinks, error)

	ProfileInfo

	app.ComponentRunnable
	TempDir() string
}

var _ app.Component = (*Anytype)(nil)

var _ Service = (*Anytype)(nil)

type ObjectsDeriver interface {
	DeriveObject(ctx context.Context, tp coresb.SmartBlockType, newAccount bool) (id string, err error)
}

type Anytype struct {
	files        *files.Service
	cafe         cafe.Client
	objectStore  objectstore.ObjectStore
	fileStore    filestore.FileStore
	fetcher      configfetcher.ConfigFetcher
	sendEvent    func(event *pb.Event)
	deriver      ObjectsDeriver
	spaceService space.Service

	ds datastore.Datastore

	predefinedBlockIds threads.DerivedSmartblockIds

	replicationWG    sync.WaitGroup
	migrationOnce    sync.Once
	lock             sync.Mutex
	isStarted        bool // use under the lock
	shutdownStartsCh chan struct {
	} // closed when node shutdown starts

	subscribeOnce       sync.Once
	config              *config.Config
	wallet              wallet.Wallet
	tmpFolderAutocreate sync.Once
	tempDir             string
	commonFiles         fileservice.FileService
}

func (a *Anytype) ThreadsIds() ([]string, error) {
	return nil, nil
}

func New() *Anytype {
	return &Anytype{
		shutdownStartsCh: make(chan struct{}),
	}
}

func (a *Anytype) Init(ap *app.App) (err error) {
	a.wallet = ap.MustComponent(wallet.CName).(wallet.Wallet)
	a.config = ap.MustComponent(config.CName).(*config.Config)
	a.objectStore = ap.MustComponent(objectstore.CName).(objectstore.ObjectStore)
	a.fileStore = ap.MustComponent(filestore.CName).(filestore.FileStore)
	a.ds = ap.MustComponent(datastore.CName).(datastore.Datastore)
	a.cafe = ap.MustComponent(cafe.CName).(cafe.Client)
	a.files = ap.MustComponent(files.CName).(*files.Service)
	a.commonFiles = ap.MustComponent(fileservice.CName).(fileservice.FileService)
	a.sendEvent = ap.MustComponent(event.CName).(event.Sender).Send
	a.fetcher = ap.MustComponent(configfetcher.CName).(configfetcher.ConfigFetcher)
	a.deriver = ap.MustComponent(treegetter.CName).(ObjectsDeriver)
	a.spaceService = ap.MustComponent(space.CName).(space.Service)
	return
}

func (a *Anytype) Name() string {
	return CName
}

func (a *Anytype) SpaceService() space.Service {
	return a.spaceService
}

// Deprecated, use wallet component directly
func (a *Anytype) Account() string {
	pk, _ := a.wallet.GetAccountPrivkey()
	if pk == nil {
		return ""
	}
	return pk.Address()
}

// Deprecated, use wallet component directly
func (a *Anytype) Device() string {
	pk, _ := a.wallet.GetDevicePrivkey()
	if pk == nil {
		return ""
	}
	return pk.Address()
}

func (a *Anytype) Run(ctx context.Context) (err error) {
	if err = a.RunMigrations(); err != nil {
		return
	}

	a.start()
	return nil
}

func (a *Anytype) IsStarted() bool {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.isStarted
}

func (a *Anytype) GetAllWorkspaces() ([]string, error) {
	return nil, nil
}

func (a *Anytype) GetWorkspaceIdForObject(objectId string) (string, error) {
	if strings.HasPrefix(objectId, "_") {
		return addr.AnytypeMarketplaceWorkspace, nil
	}
	if a.predefinedBlockIds.IsAccount(objectId) {
		return "", ErrObjectDoesNotBelongToWorkspace
	}
	return a.predefinedBlockIds.Account, nil
}

// PredefinedBlocks returns default blocks like home and archive
// ⚠️ Will return empty struct in case it runs before Anytype.Start()
func (a *Anytype) PredefinedBlocks() threads.DerivedSmartblockIds {
	return a.predefinedBlockIds
}

func (a *Anytype) HandlePeerFound(p peer.AddrInfo) {
	// TODO: [MR] mdns
}

func (a *Anytype) start() {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.isStarted {
		return
	}

	a.isStarted = true
}

func (a *Anytype) EnsurePredefinedBlocks(ctx context.Context) (err error) {
	sbTypes := []coresb.SmartBlockType{
		coresb.SmartBlockTypeWorkspace,
		coresb.SmartBlockTypeProfilePage,
		coresb.SmartBlockTypeArchive,
		coresb.SmartBlockTypeWidget,
		coresb.SmartBlockTypeHome,
	}
	for _, sbt := range sbTypes {
		var id string
		id, err = a.deriver.DeriveObject(ctx, sbt, a.config.NewAccount)
		if err != nil {
			log.With(zap.Error(err)).Debug("derived object with error")
			return
		}
		a.predefinedBlockIds.InsertId(sbt, id)
	}

	return nil
}

func (a *Anytype) Close(ctx context.Context) (err error) {
	metrics.SharedClient.Close()
	return a.Stop()
}

func (a *Anytype) Stop() error {
	fmt.Printf("stopping the library...\n")
	defer fmt.Println("library has been successfully stopped")
	a.lock.Lock()
	defer a.lock.Unlock()
	a.isStarted = false

	if a.shutdownStartsCh != nil {
		close(a.shutdownStartsCh)
	}

	// fixme useless!
	a.replicationWG.Wait()

	return nil
}

func (a *Anytype) TempDir() string {
	// it shouldn't be a case when it is called before wallet init, but just in case lets add the check here
	if a.wallet == nil || a.wallet.RootPath() == "" {
		return os.TempDir()
	}

	var err error
	// simultaneous calls to TempDir will wait for the once func to finish, so it will be fine
	a.tmpFolderAutocreate.Do(func() {
		path := filepath.Join(a.wallet.RootPath(), tmpDir)
		err = os.MkdirAll(path, 0700)
		if err != nil {
			log.Errorf("failed to make temp dir, use the default system one: %s", err.Error())
			a.tempDir = os.TempDir()
		} else {
			a.tempDir = path
		}
	})

	return a.tempDir
}
