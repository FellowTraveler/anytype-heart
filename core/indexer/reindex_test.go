package indexer

import (
	"context"
	"testing"

	"github.com/anyproto/any-sync/commonspace/spacestorage"
	"github.com/anyproto/any-sync/commonspace/spacestorage/mock_spacestorage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/anyproto/anytype-heart/core/block/editor"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock/smarttest"
	"github.com/anyproto/anytype-heart/core/block/object/objectcache/mock_objectcache"
	"github.com/anyproto/anytype-heart/core/block/source"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	coresb "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/database"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/addr"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore/mock_objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space/clientspace"
	mock_space "github.com/anyproto/anytype-heart/space/clientspace/mock_clientspace"
	"github.com/anyproto/anytype-heart/space/spacecore/storage/mock_storage"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

func TestReindexMarketplaceSpace(t *testing.T) {
	spaceId := "market"
	getMockSpace := func(fx *IndexerFixture) *clientspace.VirtualSpace {
		virtualSpace := clientspace.NewVirtualSpace(spaceId, clientspace.VirtualSpaceDeps{
			Indexer: fx,
		})
		mockCache := mock_objectcache.NewMockCache(t)
		smartTest := smarttest.New(addr.MissingObject)
		smartTest.SetSpace(virtualSpace)

		smartTest.SetType(coresb.SmartBlockTypePage)
		smartTest.SetSpaceId("spaceId")
		mockCache.EXPECT().GetObject(context.Background(), addr.MissingObject).Return(editor.NewMissingObject(smartTest), nil)
		mockCache.EXPECT().GetObject(context.Background(), addr.AnytypeProfileId).Return(smartTest, nil)
		virtualSpace.Cache = mockCache

		return virtualSpace
	}

	t.Run("reindex missing object", func(t *testing.T) {
		// given
		indexerFx := NewIndexerFixture(t)
		checksums := indexerFx.getLatestChecksums()
		err := indexerFx.store.SaveChecksums(spaceId, &checksums)
		assert.Nil(t, err)

		virtualSpace := getMockSpace(indexerFx)

		storage := mock_storage.NewMockClientStorage(t)
		storage.EXPECT().BindSpaceID(mock.Anything, mock.Anything).Return(nil)
		indexerFx.storageService = storage

		// when
		err = indexerFx.ReindexMarketplaceSpace(virtualSpace)

		// then
		details, err := indexerFx.store.GetDetails(addr.MissingObject)
		assert.Nil(t, err)
		assert.NotNil(t, details)
	})

	t.Run("do not reindex links in marketplace", func(t *testing.T) {
		// given
		fx := NewIndexerFixture(t)

		favs := []string{"fav1", "fav2"}
		trash := []string{"trash1", "trash2"}
		err := fx.store.UpdateObjectLinks("home", favs)
		require.NoError(t, err)
		err = fx.store.UpdateObjectLinks("bin", trash)
		require.NoError(t, err)

		homeLinks, err := fx.store.GetOutboundLinksByID("home")
		require.Equal(t, favs, homeLinks)

		archiveLinks, err := fx.store.GetOutboundLinksByID("bin")
		require.Equal(t, trash, archiveLinks)

		checksums := fx.getLatestChecksums()
		checksums.LinksErase = checksums.LinksErase - 1

		err = fx.objectStore.SaveChecksums(spaceId, &checksums)
		require.NoError(t, err)

		storage := mock_storage.NewMockClientStorage(t)
		storage.EXPECT().BindSpaceID(mock.Anything, mock.Anything).Return(nil)
		fx.storageService = storage

		// when
		err = fx.ReindexMarketplaceSpace(getMockSpace(fx))
		assert.NoError(t, err)

		// then
		homeLinks, err = fx.store.GetOutboundLinksByID("home")
		assert.NoError(t, err)
		assert.Equal(t, favs, homeLinks)

		archiveLinks, err = fx.store.GetOutboundLinksByID("bin")
		assert.NoError(t, err)
		assert.Equal(t, trash, archiveLinks)

		storeChecksums, err := fx.store.GetChecksums(spaceId)
		assert.Equal(t, ForceLinksReindexCounter, storeChecksums.LinksErase)
	})
}

func TestReindexDeletedObjects(t *testing.T) {
	const (
		spaceId1 = "spaceId1"
		spaceId2 = "spaceId2"
		spaceId3 = "spaceId3"
	)
	fx := NewIndexerFixture(t)

	fx.objectStore.AddObjects(t, []objectstore.TestObject{
		{
			bundle.RelationKeyId:        pbtypes.String("1"),
			bundle.RelationKeyIsDeleted: pbtypes.Bool(true),
		},
		{
			bundle.RelationKeyId:        pbtypes.String("2"),
			bundle.RelationKeyIsDeleted: pbtypes.Bool(true),
		},
		{
			bundle.RelationKeyId:        pbtypes.String("3"),
			bundle.RelationKeyIsDeleted: pbtypes.Bool(true),
			bundle.RelationKeySpaceId:   pbtypes.String(spaceId3),
		},
		{
			bundle.RelationKeyId: pbtypes.String("4"),
		},
	})

	checksums := fx.getLatestChecksums()
	checksums.AreDeletedObjectsReindexed = false

	err := fx.objectStore.SaveChecksums(spaceId1, &checksums)
	require.NoError(t, err)
	err = fx.objectStore.SaveChecksums(spaceId2, &checksums)
	require.NoError(t, err)

	t.Run("reindex first space", func(t *testing.T) {
		storage1 := mock_spacestorage.NewMockSpaceStorage(gomock.NewController(t))
		storage1.EXPECT().TreeDeletedStatus("1").Return(spacestorage.TreeDeletedStatusDeleted, nil)
		storage1.EXPECT().TreeDeletedStatus("2").Return("", nil)
		space1 := mock_space.NewMockSpace(t)
		space1.EXPECT().Id().Return(spaceId1)
		space1.EXPECT().Storage().Return(storage1)
		space1.EXPECT().StoredIds().Return([]string{})

		store := mock_objectstore.NewMockObjectStore(t)
		store.EXPECT().ListIds().Return([]string{}, nil).Times(8)
		fx.sourceFx.EXPECT().IDsListerBySmartblockType(mock.Anything, mock.Anything).Return(store, nil).Times(8)

		err = fx.ReindexSpace(space1)
		require.NoError(t, err)

		sums, err := fx.objectStore.GetChecksums(spaceId1)
		require.NoError(t, err)

		assert.True(t, sums.AreDeletedObjectsReindexed)
	})

	t.Run("reindex second space", func(t *testing.T) {
		storage2 := mock_spacestorage.NewMockSpaceStorage(gomock.NewController(t))
		storage2.EXPECT().TreeDeletedStatus("2").Return(spacestorage.TreeDeletedStatusDeleted, nil)
		space2 := mock_space.NewMockSpace(t)
		space2.EXPECT().Id().Return(spaceId2)
		space2.EXPECT().Storage().Return(storage2)
		space2.EXPECT().StoredIds().Return([]string{})
		store := mock_objectstore.NewMockObjectStore(t)
		store.EXPECT().ListIds().Return([]string{}, nil).Times(8)
		fx.sourceFx.EXPECT().IDsListerBySmartblockType(mock.Anything, mock.Anything).Return(store, nil).Times(8)

		err = fx.ReindexSpace(space2)
		require.NoError(t, err)

		sums, err := fx.objectStore.GetChecksums(spaceId2)
		require.NoError(t, err)

		assert.True(t, sums.AreDeletedObjectsReindexed)
	})

	got := fx.queryDeletedObjectIds(t, spaceId1)
	assert.Equal(t, []string{"1"}, got)

	got = fx.queryDeletedObjectIds(t, spaceId2)
	assert.Equal(t, []string{"2"}, got)

	got = fx.queryDeletedObjectIds(t, spaceId3)
	assert.Equal(t, []string{"3"}, got)
}

func TestIndexer_ReindexSpace_EraseLinks(t *testing.T) {
	const (
		spaceId1 = "space1"
		spaceId2 = "space2"
	)
	fx := NewIndexerFixture(t)

	fx.sourceFx.EXPECT().IDsListerBySmartblockType(mock.Anything, mock.Anything).RunAndReturn(
		func(_ string, sbt coresb.SmartBlockType) (source.IDsLister, error) {
			switch sbt {
			case coresb.SmartBlockTypeHome:
				return idsLister{Ids: []string{"home"}}, nil
			case coresb.SmartBlockTypeArchive:
				return idsLister{Ids: []string{"bin"}}, nil
			default:
				return idsLister{Ids: []string{}}, nil
			}
		},
	)

	fx.objectStore.AddObjects(t, []objectstore.TestObject{
		{
			bundle.RelationKeyId:      pbtypes.String("fav1"),
			bundle.RelationKeySpaceId: pbtypes.String(spaceId1),
		},
		{
			bundle.RelationKeyId:      pbtypes.String("fav2"),
			bundle.RelationKeySpaceId: pbtypes.String(spaceId1),
		},
		{
			bundle.RelationKeyId:      pbtypes.String("trash1"),
			bundle.RelationKeySpaceId: pbtypes.String(spaceId1),
		},
		{
			bundle.RelationKeyId:      pbtypes.String("trash2"),
			bundle.RelationKeySpaceId: pbtypes.String(spaceId1),
		},
		{
			bundle.RelationKeyId:      pbtypes.String("obj1"),
			bundle.RelationKeySpaceId: pbtypes.String(spaceId2),
		},
		{
			bundle.RelationKeyId:      pbtypes.String("obj2"),
			bundle.RelationKeySpaceId: pbtypes.String(spaceId2),
		},
		{
			bundle.RelationKeyId:      pbtypes.String("obj3"),
			bundle.RelationKeySpaceId: pbtypes.String(spaceId2),
		},
	})

	checksums := fx.getLatestChecksums()
	checksums.LinksErase = checksums.LinksErase - 1

	err := fx.objectStore.SaveChecksums(spaceId1, &checksums)
	require.NoError(t, err)
	err = fx.objectStore.SaveChecksums(spaceId2, &checksums)
	require.NoError(t, err)

	t.Run("links from archive and home are deleted", func(t *testing.T) {
		// given
		favs := []string{"fav1", "fav2"}
		trash := []string{"trash1", "trash2"}
		err = fx.store.UpdateObjectLinks("home", favs)
		require.NoError(t, err)
		err = fx.store.UpdateObjectLinks("bin", trash)
		require.NoError(t, err)

		homeLinks, err := fx.store.GetOutboundLinksByID("home")
		require.Equal(t, favs, homeLinks)

		archiveLinks, err := fx.store.GetOutboundLinksByID("bin")
		require.Equal(t, trash, archiveLinks)

		space1 := mock_space.NewMockSpace(t)
		space1.EXPECT().Id().Return(spaceId1)
		space1.EXPECT().StoredIds().Return([]string{}).Maybe()
		// store := mock_objectstore.NewMockObjectStore(t)
		// store.EXPECT().ListIds().Return([]string{}, nil).Times(8)

		// when
		err = fx.ReindexSpace(space1)
		assert.NoError(t, err)

		// then
		homeLinks, err = fx.store.GetOutboundLinksByID("home")
		assert.NoError(t, err)
		assert.Empty(t, homeLinks)

		archiveLinks, err = fx.store.GetOutboundLinksByID("bin")
		assert.NoError(t, err)
		assert.Empty(t, archiveLinks)

		storeChecksums, err := fx.store.GetChecksums(spaceId1)
		assert.Equal(t, ForceLinksReindexCounter, storeChecksums.LinksErase)
	})

	t.Run("links from plain objects are deleted as well", func(t *testing.T) {
		// given
		obj1links := []string{"obj2", "obj3"}
		obj2links := []string{"obj1"}
		obj3links := []string{"obj2"}
		err = fx.store.UpdateObjectLinks("obj1", obj1links)
		require.NoError(t, err)
		err = fx.store.UpdateObjectLinks("obj2", obj2links)
		require.NoError(t, err)
		err = fx.store.UpdateObjectLinks("obj3", obj3links)
		require.NoError(t, err)

		storedObj1links, err := fx.store.GetOutboundLinksByID("obj1")
		require.Equal(t, obj1links, storedObj1links)
		storedObj2links, err := fx.store.GetOutboundLinksByID("obj2")
		require.Equal(t, obj2links, storedObj2links)
		storedObj3links, err := fx.store.GetOutboundLinksByID("obj3")
		require.Equal(t, obj3links, storedObj3links)

		space1 := mock_space.NewMockSpace(t)
		space1.EXPECT().Id().Return(spaceId2)
		space1.EXPECT().StoredIds().Return([]string{}).Maybe()

		// when
		err = fx.ReindexSpace(space1)
		assert.NoError(t, err)

		// then
		storedObj1links, err = fx.store.GetOutboundLinksByID("obj1")
		assert.NoError(t, err)
		assert.Empty(t, storedObj1links)
		storedObj2links, err = fx.store.GetOutboundLinksByID("obj2")
		assert.NoError(t, err)
		assert.Empty(t, storedObj2links)
		storedObj3links, err = fx.store.GetOutboundLinksByID("obj3")
		assert.NoError(t, err)
		assert.Empty(t, storedObj3links)

		storeChecksums, err := fx.store.GetChecksums(spaceId2)
		assert.NoError(t, err)
		assert.Equal(t, ForceLinksReindexCounter, storeChecksums.LinksErase)
	})
}

func (fx *IndexerFixture) queryDeletedObjectIds(t *testing.T, spaceId string) []string {
	ids, _, err := fx.objectStore.QueryObjectIDs(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeySpaceId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(spaceId),
			},
			{
				RelationKey: bundle.RelationKeyIsDeleted.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.Bool(true),
			},
		},
	})
	require.NoError(t, err)
	return ids
}

type idsLister struct {
	Ids []string
}

func (l idsLister) ListIds() ([]string, error) {
	return l.Ids, nil
}
