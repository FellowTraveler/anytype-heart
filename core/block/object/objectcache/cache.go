package objectcache

import (
	"context"
	"errors"
	"time"

	"github.com/anyproto/any-sync/accountservice"
	"github.com/anyproto/any-sync/app/ocache"
	"github.com/anyproto/any-sync/commonspace/object/tree/treestorage"

	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/object/payloadcreator"
	"github.com/anyproto/anytype-heart/core/block/source"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/space/spacecore"
)

var log = logging.Logger("anytype-mw-object-cache")

type ctxKey int

const (
	optsKey ctxKey = iota

	ObjectLoadTimeout = time.Minute * 3
)

type treeCreateCache struct {
	initFunc smartblock.InitFunc
}

type cacheOpts struct {
	spaceId      string
	createOption *treeCreateCache
	buildOption  source.BuildOptions
	putObject    smartblock.SmartBlock
}

const CName = "client.object.objectcache"

type InitFunc = func(id string) *smartblock.InitContext

type ObjectFactory interface {
	InitObject(id string, initCtx *smartblock.InitContext) (sb smartblock.SmartBlock, err error)
}

type Cache interface {
	payloadcreator.PayloadCreator

	CreateTreeObject(ctx context.Context, params TreeCreationParams) (sb smartblock.SmartBlock, err error)
	CreateTreeObjectWithPayload(ctx context.Context, payload treestorage.TreeStorageCreatePayload, initFunc InitFunc) (sb smartblock.SmartBlock, err error)
	DeriveTreeObject(ctx context.Context, params TreeDerivationParams) (sb smartblock.SmartBlock, err error)
	GetObject(ctx context.Context, id string) (sb smartblock.SmartBlock, err error)
	GetObjectWithTimeout(ctx context.Context, id string) (sb smartblock.SmartBlock, err error)
	DoLockedIfNotExists(objectID string, proc func() error) error
	Remove(ctx context.Context, objectID string) error
	CloseBlocks()
}

type personalIDProvider interface {
	PersonalSpaceID() string
}

type objectCache struct {
	personalSpaceId string
	objectFactory   ObjectFactory
	accountService  accountservice.Service
	cache           ocache.OCache
	closing         chan struct{}
	space           *spacecore.AnySpace
	keyConverter    smartblock.KeyToIDConverter
}

func New(
	space *spacecore.AnySpace,
	accountService accountservice.Service,
	objectFactory ObjectFactory,
	personalSpaceId string,
	keyConverter smartblock.KeyToIDConverter,
) Cache {
	c := &objectCache{
		personalSpaceId: personalSpaceId,
		space:           space,
		accountService:  accountService,
		objectFactory:   objectFactory,
		closing:         make(chan struct{}),
		keyConverter:    keyConverter,
	}
	c.cache = ocache.New(
		c.cacheLoad,
		// ocache.WithLogger(log.Desugar()),
		ocache.WithGCPeriod(time.Minute),
		// TODO: [MR] Get ttl from config
		ocache.WithTTL(time.Duration(60)*time.Second),
	)
	return c
}

func (c *objectCache) Close(_ context.Context) error {
	close(c.closing)
	return c.cache.Close()
}

func ContextWithCreateOption(ctx context.Context, initFunc smartblock.InitFunc) context.Context {
	return context.WithValue(ctx, optsKey, cacheOpts{
		createOption: &treeCreateCache{
			initFunc: initFunc,
		},
	})
}

func ContextWithBuildOptions(ctx context.Context, buildOpts source.BuildOptions) context.Context {
	return context.WithValue(ctx,
		optsKey,
		cacheOpts{
			buildOption: buildOpts,
		},
	)
}

func (c *objectCache) cacheLoad(ctx context.Context, id string) (value ocache.Object, err error) {
	opts := ctx.Value(optsKey).(cacheOpts)
	buildObject := func(id string) (sb smartblock.SmartBlock, err error) {
		initCtx := &smartblock.InitContext{
			Ctx:              ctx,
			BuildOpts:        opts.buildOption,
			SpaceID:          opts.spaceId,
			KeyToIDConverter: c.keyConverter,
		}
		return c.objectFactory.InitObject(id, initCtx)
	}
	createObject := func() (sb smartblock.SmartBlock, err error) {
		initCtx := opts.createOption.initFunc(id)
		initCtx.IsNewObject = true
		initCtx.Ctx = ctx
		initCtx.SpaceID = opts.spaceId
		initCtx.BuildOpts = opts.buildOption
		return c.objectFactory.InitObject(id, initCtx)
	}

	switch {
	case opts.createOption != nil:
		return createObject()
	case opts.putObject != nil:
		// putting object through cache
		return opts.putObject, nil
	default:
		break
	}

	return buildObject(id)
}

func (c *objectCache) GetObject(ctx context.Context, id string) (sb smartblock.SmartBlock, err error) {
	ctx = updateCacheOpts(ctx, func(opts cacheOpts) cacheOpts {
		opts.spaceId = c.space.Id()
		return opts
	})
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var (
		done    = make(chan struct{})
		closing bool
	)
	var start time.Time
	go func() {
		select {
		case <-done:
			cancel()
		case <-c.closing:
			start = time.Now()
			cancel()
			closing = true
		}
	}()
	v, err := c.cache.Get(ctx, id)
	close(done)
	if closing && errors.Is(err, context.Canceled) {
		log.With("close_delay", time.Since(start).Milliseconds()).With("objectID", id).Warnf("object was loading during closing")
	}
	if err != nil {
		return
	}
	return v.(smartblock.SmartBlock), nil
}

func (c *objectCache) Remove(ctx context.Context, objectID string) error {
	_, err := c.cache.Remove(ctx, objectID)
	return err
}

func (c *objectCache) GetObjectWithTimeout(ctx context.Context, id string) (sb smartblock.SmartBlock, err error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, ObjectLoadTimeout)
		defer cancel()
	}
	return c.GetObject(ctx, id)
}

func (c *objectCache) DoLockedIfNotExists(objectID string, proc func() error) error {
	return c.cache.DoLockedIfNotExists(objectID, proc)
}

func (c *objectCache) CloseBlocks() {
	c.cache.ForEach(func(v ocache.Object) (isContinue bool) {
		ob := v.(smartblock.SmartBlock)
		ob.Lock()
		ob.ObjectCloseAllSessions()
		ob.Unlock()
		return true
	})
}

func CacheOptsWithRemoteLoadDisabled(ctx context.Context) context.Context {
	return updateCacheOpts(ctx, func(opts cacheOpts) cacheOpts {
		opts.buildOption.DisableRemoteLoad = true
		return opts
	})
}

func updateCacheOpts(ctx context.Context, update func(opts cacheOpts) cacheOpts) context.Context {
	opts, ok := ctx.Value(optsKey).(cacheOpts)
	if !ok {
		opts = cacheOpts{}
	}
	return context.WithValue(ctx, optsKey, update(opts))
}
