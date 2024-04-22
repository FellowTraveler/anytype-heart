package device

import (
	"context"
	"os"
	"testing"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/commonspace/object/tree/treestorage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/anyproto/anytype-heart/core/block/editor"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock/smarttest"
	"github.com/anyproto/anytype-heart/core/block/object/objectcache/mock_objectcache"
	wallet2 "github.com/anyproto/anytype-heart/core/wallet"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space/clientspace"
	"github.com/anyproto/anytype-heart/space/mock_space"
	"github.com/anyproto/anytype-heart/tests/testutil"
)

func TestService_SaveDeviceInfo(t *testing.T) {
	deviceObjectId := "deviceObjectId"
	techSpaceId := "techSpaceId"
	t.Run("save device in object", func(t *testing.T) {
		// given
		testDevice := &model.DeviceInfo{
			Id:   "id",
			Name: "test",
		}

		devicesService := newFixture(t, deviceObjectId)
		virtualSpace := clientspace.NewVirtualSpace(techSpaceId, clientspace.VirtualSpaceDeps{})
		devicesService.mockSpaceService.EXPECT().Get(context.Background(), techSpaceId).Return(virtualSpace, nil)
		devicesService.mockSpaceService.EXPECT().TechSpaceId().Return(techSpaceId)

		deviceObject := &editor.Page{SmartBlock: smarttest.New(deviceObjectId)}
		mockCache := mock_objectcache.NewMockCache(t)
		mockCache.EXPECT().GetObject(context.Background(), deviceObjectId).Return(deviceObject, nil)

		virtualSpace.Cache = mockCache

		// when
		err := devicesService.SaveDeviceInfo(context.Background(), testDevice)

		// then
		assert.Nil(t, err)
		assert.NotNil(t, deviceObject.NewState().GetDevice("id"))
	})

	t.Run("save device in object, device exist", func(t *testing.T) {
		// given
		testDevice := &model.DeviceInfo{
			Id:   "id",
			Name: "test",
		}

		devicesService := newFixture(t, deviceObjectId)
		virtualSpace := clientspace.NewVirtualSpace(techSpaceId, clientspace.VirtualSpaceDeps{})
		devicesService.mockSpaceService.EXPECT().Get(context.Background(), techSpaceId).Return(virtualSpace, nil)
		devicesService.mockSpaceService.EXPECT().TechSpaceId().Return(techSpaceId)

		deviceObject := &editor.Page{SmartBlock: smarttest.New(deviceObjectId)}
		mockCache := mock_objectcache.NewMockCache(t)
		mockCache.EXPECT().GetObject(context.Background(), deviceObjectId).Return(deviceObject, nil)
		virtualSpace.Cache = mockCache

		testDevice1 := &model.DeviceInfo{
			Id:   "id",
			Name: "test1",
		}

		// when
		err := devicesService.SaveDeviceInfo(context.Background(), testDevice)
		err = devicesService.SaveDeviceInfo(context.Background(), testDevice1)

		// then
		assert.Nil(t, err)
		assert.NotNil(t, deviceObject.NewState().GetDevice("id"))
		assert.Equal(t, "test", deviceObject.NewState().GetDevice("id").Name)
	})
}

func TestService_UpdateName(t *testing.T) {
	deviceObjectId := "deviceObjectId"
	techSpaceId := "techSpaceId"
	t.Run("update name, device not exist", func(t *testing.T) {
		// given

		devicesService := newFixture(t, deviceObjectId)
		virtualSpace := clientspace.NewVirtualSpace(techSpaceId, clientspace.VirtualSpaceDeps{})
		devicesService.mockSpaceService.EXPECT().Get(context.Background(), techSpaceId).Return(virtualSpace, nil)
		devicesService.mockSpaceService.EXPECT().TechSpaceId().Return(techSpaceId)

		deviceObject := &editor.Page{SmartBlock: smarttest.New(deviceObjectId)}
		mockCache := mock_objectcache.NewMockCache(t)
		mockCache.EXPECT().GetObject(context.Background(), deviceObjectId).Return(deviceObject, nil)

		virtualSpace.Cache = mockCache

		// when
		err := devicesService.UpdateName(context.Background(), "id", "new name")

		// then
		assert.Nil(t, err)
		assert.NotNil(t, deviceObject.NewState().GetDevice("id"))
		assert.Equal(t, "new name", deviceObject.NewState().GetDevice("id").Name)
	})

	t.Run("update name, device exists", func(t *testing.T) {
		// given
		testDevice := &model.DeviceInfo{
			Id:   "id",
			Name: "test",
		}

		devicesService := newFixture(t, deviceObjectId)
		virtualSpace := clientspace.NewVirtualSpace(techSpaceId, clientspace.VirtualSpaceDeps{})
		devicesService.mockSpaceService.EXPECT().Get(context.Background(), techSpaceId).Return(virtualSpace, nil)
		devicesService.mockSpaceService.EXPECT().TechSpaceId().Return(techSpaceId)

		deviceObject := &editor.Page{SmartBlock: smarttest.New(deviceObjectId)}
		mockCache := mock_objectcache.NewMockCache(t)
		mockCache.EXPECT().GetObject(context.Background(), deviceObjectId).Return(deviceObject, nil)

		virtualSpace.Cache = mockCache
		err := devicesService.SaveDeviceInfo(context.Background(), testDevice)
		assert.Nil(t, err)

		// when
		err = devicesService.UpdateName(context.Background(), "id", "new name")

		// then
		assert.Nil(t, err)
		assert.NotNil(t, deviceObject.NewState().GetDevice("id"))
		assert.Equal(t, "new name", deviceObject.NewState().GetDevice("id").Name)
	})
}

func TestService_ListDevices(t *testing.T) {
	deviceObjectId := "deviceObjectId"
	techSpaceId := "techSpaceId"
	t.Run("list devices, no devices", func(t *testing.T) {
		// given

		devicesService := newFixture(t, deviceObjectId)
		virtualSpace := clientspace.NewVirtualSpace(techSpaceId, clientspace.VirtualSpaceDeps{})
		devicesService.mockSpaceService.EXPECT().Get(context.Background(), techSpaceId).Return(virtualSpace, nil)
		devicesService.mockSpaceService.EXPECT().TechSpaceId().Return(techSpaceId)

		deviceObject := &editor.Page{SmartBlock: smarttest.New(deviceObjectId)}
		mockCache := mock_objectcache.NewMockCache(t)
		mockCache.EXPECT().GetObject(context.Background(), deviceObjectId).Return(deviceObject, nil)

		virtualSpace.Cache = mockCache

		// when
		devicesList, err := devicesService.ListDevices(context.Background())

		// then
		assert.Nil(t, err)
		assert.Len(t, devicesList, 0)
	})

	t.Run("list devices", func(t *testing.T) {
		// given
		testDevice := &model.DeviceInfo{
			Id:   "id",
			Name: "test",
		}

		testDevice1 := &model.DeviceInfo{
			Id:   "id1",
			Name: "test1",
		}

		devicesService := newFixture(t, deviceObjectId)
		virtualSpace := clientspace.NewVirtualSpace(techSpaceId, clientspace.VirtualSpaceDeps{})
		devicesService.mockSpaceService.EXPECT().Get(context.Background(), techSpaceId).Return(virtualSpace, nil)
		devicesService.mockSpaceService.EXPECT().TechSpaceId().Return(techSpaceId)

		deviceObject := &editor.Page{SmartBlock: smarttest.New(deviceObjectId)}
		mockCache := mock_objectcache.NewMockCache(t)
		mockCache.EXPECT().GetObject(context.Background(), deviceObjectId).Return(deviceObject, nil)

		virtualSpace.Cache = mockCache

		err := devicesService.SaveDeviceInfo(context.Background(), testDevice)
		assert.Nil(t, err)
		err = devicesService.SaveDeviceInfo(context.Background(), testDevice1)
		assert.Nil(t, err)

		// when
		devicesList, err := devicesService.ListDevices(context.Background())

		// then
		assert.Nil(t, err)
		assert.Len(t, devicesList, 2)
		assert.Equal(t, devicesList[0].Id, "id")
		assert.Equal(t, devicesList[1].Id, "id1")
	})
}

func TestService_loadDevices(t *testing.T) {
	deviceObjectId := "deviceObjectId"
	techSpaceId := "techSpaceId"
	ctx := context.Background()
	t.Run("loadDevices, device object not exist", func(t *testing.T) {
		// given
		devicesService := newFixture(t, deviceObjectId)
		virtualSpace := clientspace.NewVirtualSpace(techSpaceId, clientspace.VirtualSpaceDeps{})
		devicesService.mockSpaceService.EXPECT().GetTechSpace(ctx).Return(virtualSpace, nil)

		deviceObject := &editor.Page{SmartBlock: smarttest.New(deviceObjectId)}
		mockCache := mock_objectcache.NewMockCache(t)
		mockCache.EXPECT().GetObject(ctx, deviceObjectId).Return(deviceObject, nil)
		mockCache.EXPECT().DeriveTreeObject(ctx, mock.Anything).Return(nil, treestorage.ErrTreeExists)
		mockCache.EXPECT().DeriveObjectID(ctx, mock.Anything).Return(deviceObjectId, nil)
		virtualSpace.Cache = mockCache

		// when
		devicesService.loadDevices(ctx)

		// then
		assert.NotNil(t, deviceObject.NewState().GetDevice(devicesService.wallet.GetDevicePrivkey().GetPublic().PeerId()))
	})

	t.Run("loadDevices, device object exist", func(t *testing.T) {
		// given
		devicesService := newFixture(t, deviceObjectId)
		virtualSpace := clientspace.NewVirtualSpace(techSpaceId, clientspace.VirtualSpaceDeps{})
		devicesService.mockSpaceService.EXPECT().GetTechSpace(ctx).Return(virtualSpace, nil)

		deviceObject := &editor.Page{SmartBlock: smarttest.New(deviceObjectId)}
		mockCache := mock_objectcache.NewMockCache(t)
		mockCache.EXPECT().DeriveTreeObject(ctx, mock.Anything).Return(deviceObject, nil)
		virtualSpace.Cache = mockCache

		// when
		devicesService.loadDevices(ctx)

		// then
		assert.NotNil(t, deviceObject.NewState().GetDevice(devicesService.wallet.GetDevicePrivkey().GetPublic().PeerId()))
	})
}

type deviceFixture struct {
	*devices

	mockSpaceService *mock_space.MockService
	mockCache        *mock_objectcache.MockCache
	wallet           wallet2.Wallet
}

func newFixture(t *testing.T, deviceObjectId string) *deviceFixture {
	mockSpaceService := mock_space.NewMockService(t)
	mockCache := mock_objectcache.NewMockCache(t)
	wallet := wallet2.NewWithRepoDirAndRandomKeys(os.TempDir())
	df := &deviceFixture{
		mockSpaceService: mockSpaceService,
		mockCache:        mockCache,
		wallet:           wallet,
		devices:          &devices{deviceObjectId: deviceObjectId},
	}

	a := &app.App{}

	a.Register(testutil.PrepareMock(context.Background(), a, mockSpaceService)).
		Register(wallet)

	err := wallet.Init(a)
	assert.Nil(t, err)
	err = df.Init(a)
	assert.Nil(t, err)
	return df
}
