package history

import (
	"fmt"
	"time"

	"github.com/anytypeio/go-anytype-middleware/app"
	"github.com/anytypeio/go-anytype-middleware/change"
	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/meta"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
)

const CName = "history"

const versionGroupInterval = time.Minute * 5

var log = logging.Logger("anytype-mw-history")

func New() History {
	return new(history)
}

type History interface {
	Show(pageId, versionId string) (bs *pb.EventObjectShow, ver *pb.RpcHistoryVersionsVersion, err error)
	Versions(pageId, lastVersionId string, limit int) (resp []*pb.RpcHistoryVersionsVersion, err error)
	SetVersion(pageId, versionId string) (err error)
	app.Component
}

type BlockService interface {
	ResetToState(pageId string, s *state.State) (err error)
}

type history struct {
	a    core.Service
	bs   BlockService
	meta meta.Service
}

func (h *history) Init(a *app.App) (err error) {
	h.a = a.MustComponent(core.CName).(core.Service)
	h.bs = a.MustComponent(block.CName).(BlockService)
	h.meta = a.MustComponent(meta.CName).(meta.Service)
	return
}

func (h *history) Name() (name string) {
	return CName
}

func (h *history) Show(pageId, versionId string) (bs *pb.EventObjectShow, ver *pb.RpcHistoryVersionsVersion, err error) {
	s, ver, err := h.buildState(pageId, versionId)
	if err != nil {
		return
	}

	metaD := h.meta.FetchMeta(s.DepSmartIds())
	details := make([]*pb.EventObjectDetailsSet, 0, len(metaD))
	var uniqueObjTypes []string
	sbType, err := smartblock.SmartBlockTypeFromID(pageId)
	if err != nil {
		return nil, nil, fmt.Errorf("incorrect sb type: %w", err)
	}
	metaD = append(metaD, meta.Meta{BlockId: pageId, SmartBlockMeta: core.SmartBlockMeta{ObjectTypes: s.ObjectTypes(), Details: s.Details()}})
	for _, m := range metaD {
		details = append(details, &pb.EventObjectDetailsSet{
			Id:      m.BlockId,
			Details: m.Details,
		})

		if len(m.ObjectTypes) > 0 {
			if len(m.ObjectTypes) > 1 {
				log.Error("object has more than 1 object type which is not supported on clients. types are truncated")
			}
			if slice.FindPos(uniqueObjTypes, m.ObjectTypes[0]) == -1 {
				uniqueObjTypes = append(uniqueObjTypes, m.ObjectTypes[0])
			}
		}

	}

	objectTypes := h.meta.FetchObjectTypes(uniqueObjTypes)
	return &pb.EventObjectShow{
		RootId:      pageId,
		Type:        smartblock.SmartBlockTypeToProto(sbType),
		Blocks:      s.Blocks(),
		Details:     details,
		ObjectTypes: objectTypes,
	}, ver, nil
}

func (h *history) Versions(pageId, lastVersionId string, limit int) (resp []*pb.RpcHistoryVersionsVersion, err error) {
	if limit <= 0 {
		limit = 100
	}
	profileId, profileName, err := h.getProfileInfo()
	if err != nil {
		return
	}
	var includeLastId = true

	reverse := func(vers []*pb.RpcHistoryVersionsVersion) []*pb.RpcHistoryVersionsVersion {
		for i, j := 0, len(vers)-1; i < j; i, j = i+1, j-1 {
			vers[i], vers[j] = vers[j], vers[i]
		}
		return vers
	}

	for len(resp) < limit {
		tree, _, e := h.buildTree(pageId, lastVersionId, includeLastId)
		if e != nil {
			return nil, e
		}
		var data []*pb.RpcHistoryVersionsVersion

		tree.Iterate(tree.RootId(), func(c *change.Change) (isContinue bool) {
			data = append(data, &pb.RpcHistoryVersionsVersion{
				Id:          c.Id,
				PreviousIds: c.PreviousIds,
				AuthorId:    profileId,
				AuthorName:  profileName,
				Time:        c.Timestamp,
			})
			return true
		})
		if len(data[0].PreviousIds) == 0 {
			if h.isEmpty(tree.Get(data[0].Id)) {
				data = data[1:]
			}
			resp = append(data, resp...)
			break
		} else {
			resp = append(data, resp...)
			lastVersionId = tree.RootId()
			includeLastId = false
		}

		if len(data) == 0 {
			break
		}

	}

	resp = reverse(resp)

	var groupId int64
	var nextVersionTimestamp int64

	for i := 0; i < len(resp); i++ {
		if nextVersionTimestamp-resp[i].Time > int64(versionGroupInterval.Seconds()) {
			groupId++
		}
		nextVersionTimestamp = resp[i].Time
		resp[i].GroupId = groupId
	}

	return
}

func (h *history) isEmpty(c *change.Change) bool {
	if c.Snapshot != nil && c.Snapshot.Data != nil {
		if c.Snapshot.Data.Details != nil && c.Snapshot.Data.Details.Fields != nil && len(c.Snapshot.Data.Details.Fields) > 0 {
			return false
		}
		for _, b := range c.Snapshot.Data.Blocks {
			if b.GetSmartblock() != nil && b.GetLayout() != nil {
				return false
			}
		}
		return true
	}
	return false
}

func (h *history) SetVersion(pageId, versionId string) (err error) {
	s, _, err := h.buildState(pageId, versionId)
	if err != nil {
		return
	}
	return h.bs.ResetToState(pageId, s)
}

func (h *history) buildTree(pageId, versionId string, includeLastId bool) (tree *change.Tree, blockType smartblock.SmartBlockType, err error) {
	sb, err := h.a.GetBlock(pageId)
	if err != nil {
		err = fmt.Errorf("history: anytype.GetBlock error: %v", err)
		return
	}
	if tree, err = change.BuildTreeBefore(sb, versionId, includeLastId); err != nil {
		return
	}
	return tree, sb.Type(), nil
}

func (h *history) buildState(pageId, versionId string) (s *state.State, ver *pb.RpcHistoryVersionsVersion, err error) {
	tree, sbType, err := h.buildTree(pageId, versionId, true)
	if err != nil {
		return
	}
	root := tree.Root()
	if root == nil || root.GetSnapshot() == nil {
		return nil, nil, fmt.Errorf("root missing or not a snapshot")
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
	switch sbType {
	case smartblock.SmartBlockTypePage, smartblock.SmartBlockTypeProfilePage, smartblock.SmartBlockTypeSet:
		// todo: set case not handled
		template.InitTemplate(s, template.WithTitle)
	}
	s.BlocksInit()
	if ch := tree.Get(versionId); ch != nil {
		profileId, profileName, e := h.getProfileInfo()
		if e != nil {
			err = e
			return
		}
		ver = &pb.RpcHistoryVersionsVersion{
			Id:          ch.Id,
			PreviousIds: ch.PreviousIds,
			AuthorId:    profileId,
			AuthorName:  profileName,
			Time:        ch.Timestamp,
		}
	}
	return
}

func (h *history) getProfileInfo() (profileId, profileName string, err error) {
	profileId = h.a.ProfileID()
	lp, err := h.a.LocalProfile()
	if err != nil {
		return
	}
	profileName = lp.Name
	return
}
