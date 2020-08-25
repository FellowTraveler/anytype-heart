package history

import (
	"fmt"

	"github.com/anytypeio/go-anytype-middleware/change"
	"github.com/anytypeio/go-anytype-middleware/core/anytype"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/pb"
)

func NewHistory(a anytype.Service, bs BlockService) History {
	return &history{
		a:  a,
		bs: bs,
	}
}

type History interface {
	Show(pageId, versionId string) (bs *pb.EventBlockShow, err error)
	Versions(pageId, lastVersionId string, limit int) (resp []*pb.RpcHistoryVersionsVersion, err error)
	SetVersion(pageId, versionId string) (err error)
}

type BlockService interface {
	ResetToState(pageId string, s *state.State) (err error)
}

type history struct {
	a  anytype.Service
	bs BlockService
}

func (h *history) Show(pageId, versionId string) (bs *pb.EventBlockShow, err error) {
	s, err := h.buildState(pageId, versionId)
	if err != nil {
		return
	}
	return &pb.EventBlockShow{
		RootId:  pageId,
		Blocks:  s.Blocks(),
		Details: nil, // TODO:
	}, nil
}

func (h *history) Versions(pageId, lastVersionId string, limit int) (resp []*pb.RpcHistoryVersionsVersion, err error) {
	if limit <= 0 {
		limit = 100
	}
	for len(resp) < limit {
		tree, e := h.buildTree(pageId, lastVersionId)
		if e != nil {
			return nil, e
		}
		if tree.Len() == 1 && tree.RootId() == lastVersionId {
			return
		}
		var data []*pb.RpcHistoryVersionsVersion
		tree.Iterate(tree.RootId(), func(c *change.Change) (isContinue bool) {
			if c.Id != lastVersionId {
				data = append(data, &pb.RpcHistoryVersionsVersion{
					Id:          c.Id,
					PreviousIds: c.PreviousIds,
					Time:        c.Timestamp,
				})
			}
			return true
		})
		resp = append(data, resp...)
		lastVersionId = tree.RootId()
	}
	return
}

func (h *history) SetVersion(pageId, versionId string) (err error) {
	s, err := h.buildState(pageId, versionId)
	if err != nil {
		return
	}
	return h.bs.ResetToState(pageId, s)
}

func (h *history) buildTree(pageId, versionId string) (*change.Tree, error) {
	sb, err := h.a.GetBlock(pageId)
	if err != nil {
		err = fmt.Errorf("history: anytype.GetBlock error: %v", err)
		return nil, nil
	}
	return change.BuildTreeBefore(sb, versionId)
}

func (h *history) buildState(pageId, versionId string) (s *state.State, err error) {
	tree, err := h.buildTree(pageId, versionId)
	if err != nil {
		return
	}
	root := tree.Root()
	if root == nil || root.GetSnapshot() == nil {
		return nil, fmt.Errorf("root missing or not a snapshot")
	}
	s = state.NewDocFromSnapshot(pageId, root.GetSnapshot()).(*state.State)
	s.SetChangeId(root.Id)
	st, err := change.BuildStateSimpleCRDT(s, tree)
	if err != nil {
		return
	}
	if _, _, err = state.ApplyStateFast(st); err != nil {
		return
	}
	return
}
