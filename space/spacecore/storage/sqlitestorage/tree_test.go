package sqlitestorage

import (
	"testing"

	"github.com/anyproto/any-sync/commonspace/object/tree/treechangeproto"
	"github.com/anyproto/any-sync/commonspace/object/tree/treestorage"
	"github.com/stretchr/testify/require"
)

func TestTreeStorage_Create(t *testing.T) {
	fx := newFixture(t)
	defer fx.finish(t)

	spacePayload := spaceTestPayload()
	ss, err := createSpaceStorage(fx.storageService, spacePayload)
	require.NoError(t, err)
	payload := treeTestPayload()
	store, err := ss.CreateTreeStorage(payload)
	require.NoError(t, err)
	testTreePayload(t, store, payload)

	t.Run("create same storage returns error", func(t *testing.T) {
		_, err := ss.CreateTreeStorage(payload)
		require.Error(t, err)
	})
}

func TestTreeStorage_Methods(t *testing.T) {
	fx := newFixture(t)
	defer fx.finish(t)

	spacePayload := spaceTestPayload()
	ss, err := createSpaceStorage(fx.storageService, spacePayload)
	require.NoError(t, err)
	payload := treeTestPayload()
	store, err := ss.CreateTreeStorage(payload)
	require.NoError(t, err)

	store, err = ss.TreeStorage(payload.RootRawChange.Id)
	require.NoError(t, err)
	testTreePayload(t, store, payload)

	t.Run("update heads", func(t *testing.T) {
		newHeads := []string{"a", "b"}
		require.NoError(t, store.SetHeads(newHeads))
		heads, err := store.Heads()
		require.NoError(t, err)
		require.Equal(t, newHeads, heads)
	})

	t.Run("add raw change, get change and has change", func(t *testing.T) {
		newChange := &treechangeproto.RawTreeChangeWithId{RawChange: []byte("ab"), Id: "newId"}
		require.NoError(t, store.AddRawChange(newChange))
		rawCh, err := store.GetRawChange(ctx, newChange.Id)
		require.NoError(t, err)
		require.Equal(t, newChange, rawCh)
		has, err := store.HasChange(ctx, newChange.Id)
		require.NoError(t, err)
		require.True(t, has)
	})

	t.Run("get and has for unknown change", func(t *testing.T) {
		incorrectId := "incorrectId"
		_, err := store.GetRawChange(ctx, incorrectId)
		require.Error(t, err)
		has, err := store.HasChange(ctx, incorrectId)
		require.NoError(t, err)
		require.False(t, has)
	})
}

func TestTreeStorage_Delete(t *testing.T) {
	fx := newFixture(t)
	defer fx.finish(t)

	spacePayload := spaceTestPayload()
	ss, err := createSpaceStorage(fx.storageService, spacePayload)
	require.NoError(t, err)
	payload := treeTestPayload()
	store, err := ss.CreateTreeStorage(payload)
	require.NoError(t, err)

	t.Run("add raw change, get change and has change", func(t *testing.T) {
		newChange := &treechangeproto.RawTreeChangeWithId{RawChange: []byte("ab"), Id: "newId"}
		require.NoError(t, store.AddRawChange(newChange))

		err = store.Delete()
		require.NoError(t, err)

		_, err = ss.TreeStorage(payload.RootRawChange.Id)
		require.Equal(t, err, treestorage.ErrUnknownTreeId)
	})
}

func treeTestPayload() treestorage.TreeStorageCreatePayload {
	rootRawChange := &treechangeproto.RawTreeChangeWithId{RawChange: []byte("some"), Id: "someRootId"}
	otherChange := &treechangeproto.RawTreeChangeWithId{RawChange: []byte("some other"), Id: "otherId"}
	changes := []*treechangeproto.RawTreeChangeWithId{rootRawChange, otherChange}
	return treestorage.TreeStorageCreatePayload{
		RootRawChange: rootRawChange,
		Changes:       changes,
		Heads:         []string{rootRawChange.Id},
	}
}

func testTreePayload(t *testing.T, store treestorage.TreeStorage, payload treestorage.TreeStorageCreatePayload) {
	require.Equal(t, payload.RootRawChange.Id, store.Id())

	root, err := store.Root()
	require.NoError(t, err)
	require.Equal(t, root, payload.RootRawChange)

	heads, err := store.Heads()
	require.NoError(t, err)
	require.Equal(t, payload.Heads, heads)

	for _, ch := range payload.Changes {
		dbCh, err := store.GetRawChange(ctx, ch.Id)
		require.NoError(t, err)
		require.Equal(t, ch, dbCh)
	}
	return
}
