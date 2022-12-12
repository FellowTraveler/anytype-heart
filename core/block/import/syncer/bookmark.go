package syncer

import (
	"fmt"

	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/core/session"
	"github.com/anytypeio/go-anytype-middleware/pb"
)

type BookmarkSyncer struct {
	service *block.Service
}

func NewBookmarkSyncer(service *block.Service) *BookmarkSyncer {
	return &BookmarkSyncer{service: service}
}

func (bs *BookmarkSyncer) Sync(ctx *session.Context, id string, b simple.Block) error {
	err := bs.service.BookmarkFetch(ctx, pb.RpcBlockBookmarkFetchRequest{
		ContextId: id,
		BlockId:   b.Model().GetBookmark().TargetObjectId,
		Url:       b.Model().GetBookmark().Url,
	})
	if err != nil {
		return fmt.Errorf("failed syncing bookmark: %s", err)
	}
	return nil
}
