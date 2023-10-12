package space

import (
	"testing"

	"github.com/anyproto/any-sync/commonspace/object/accountdata"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"

	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
)

func TestSpace_deriveAccountMetadata(t *testing.T) {
	randKeys, err := accountdata.NewRandom()
	require.NoError(t, err)
	symKey, err := deriveAccountEncKey(randKeys.SignKey)
	require.NoError(t, err)
	symKeyRaw, err := symKey.Raw()
	require.NoError(t, err)
	metadata1, err := deriveAccountMetadata(randKeys.SignKey)
	require.NoError(t, err)
	metadata2, err := deriveAccountMetadata(randKeys.SignKey)
	require.NoError(t, err)
	require.Equal(t, metadata1, metadata2)
	metadata := &model.Metadata{}
	err = proto.Unmarshal(metadata1, metadata)
	require.NoError(t, err)
	require.Equal(t, model.MetadataType_Identity, metadata.MetadataType)
	metadataPayload := &model.MetadataIdentityPayload{}
	err = proto.Unmarshal(metadata.MetadataPayload, metadataPayload)
	require.NoError(t, err)
	require.Equal(t, symKeyRaw, metadataPayload.ProfileSymKey)
}
