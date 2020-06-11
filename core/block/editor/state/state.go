package state

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/anytypeio/go-anytype-library/logging"
	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-middleware/core/block/history"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
	"github.com/gogo/protobuf/types"
)

var log = logging.Logger("anytype-mw-state")

type Doc interface {
	RootId() string
	NewState() *State
	NewStateCtx(ctx *Context) *State
	Blocks() []*model.Block
	Pick(id string) (b simple.Block)
	Append(targetId string, id string) (ok bool)
}

func NewDoc(rootId string, blocks map[string]simple.Block) Doc {
	if blocks == nil {
		blocks = make(map[string]simple.Block)
	}
	return &State{
		rootId: rootId,
		blocks: blocks,
	}
}

type State struct {
	ctx      *Context
	parent   *State
	blocks   map[string]simple.Block
	rootId   string
	newIds   []string
	changeId string
	changes  []*pb.ChangeContent
	details  *types.Struct

	changesStructureIgnoreIds []string

	bufIterateParentIds []string
}

func (s *State) RootId() string {
	return s.rootId
}

func (s *State) NewState() *State {
	return &State{parent: s, blocks: make(map[string]simple.Block), rootId: s.rootId}
}

func (s *State) NewStateCtx(ctx *Context) *State {
	return &State{parent: s, blocks: make(map[string]simple.Block), rootId: s.rootId, ctx: ctx}
}

func (s *State) Context() *Context {
	return s.ctx
}

func (s *State) Add(b simple.Block) (ok bool) {
	id := b.Model().Id
	if s.Pick(id) == nil {
		s.blocks[id] = b
		if s.parent != nil {
			s.newIds = append(s.newIds, id)
		}
		return true
	}
	return false
}

func (s *State) Set(b simple.Block) {
	if !s.Exists(b.Model().Id) {
		s.Add(b)
	} else {
		s.blocks[b.Model().Id] = b
	}
}

func (s *State) Get(id string) (b simple.Block) {
	if b = s.blocks[id]; b != nil {
		return
	}
	if s.parent != nil {
		if b = s.Pick(id); b != nil {
			b = b.Copy()
			s.blocks[id] = b
			return
		}
	}
	return
}

func (s *State) Pick(id string) (b simple.Block) {
	var (
		t  = s
		ok bool
	)
	for t != nil {
		if b, ok = t.blocks[id]; ok {
			return
		}
		t = t.parent
	}
	return
}

func (s *State) PickOrigin(id string) (b simple.Block) {
	if s.parent != nil {
		return s.parent.Pick(id)
	}
	return
}

func (s *State) Unlink(id string) (ok bool) {
	if parent := s.GetParentOf(id); parent != nil {
		parentM := parent.Model()
		parentM.ChildrenIds = slice.Remove(parentM.ChildrenIds, id)
		return true
	}
	return
}

func (s *State) Append(targetId string, id string) (ok bool) {
	parent := s.Get(targetId).Model()
	parent.ChildrenIds = append(parent.ChildrenIds, id)
	return true
}

func (s *State) GetParentOf(id string) (res simple.Block) {
	if parent := s.PickParentOf(id); parent != nil {
		return s.Get(parent.Model().Id)
	}
	return
}

func (s *State) PickParentOf(id string) (res simple.Block) {
	s.Iterate(func(b simple.Block) bool {
		if slice.FindPos(b.Model().ChildrenIds, id) != -1 {
			res = b
			return false
		}
		return true
	})
	return
}

func (s *State) PickOriginParentOf(id string) (res simple.Block) {
	if s.parent != nil {
		return s.parent.PickParentOf(id)
	}
	return
}

func (s *State) Diff(id string) ([]*pb.EventMessage, error) {
	if new := s.blocks[id]; new != nil && s.parent != nil {
		if old := s.PickOrigin(id); old != nil {
			return old.Diff(new)
		}
	}
	return nil, nil
}

func (s *State) Iterate(f func(b simple.Block) (isContinue bool)) (err error) {
	var iter func(id string) (isContinue bool, err error)
	var parentIds = s.bufIterateParentIds[:0]
	iter = func(id string) (isContinue bool, err error) {
		if slice.FindPos(parentIds, id) != -1 {
			return false, fmt.Errorf("cycle reference: %v %s", parentIds, id)
		}
		parentIds = append(parentIds, id)
		parentSize := len(parentIds)
		b := s.Pick(id)
		if b != nil {
			if isContinue = f(b); !isContinue {
				return
			}
			for _, cid := range b.Model().ChildrenIds {
				if isContinue, err = iter(cid); !isContinue || err != nil {
					return
				}
				parentIds = parentIds[:parentSize]
			}
		}
		return true, nil
	}
	_, err = iter(s.RootId())
	return
}

func (s *State) Exists(id string) (ok bool) {
	return s.Pick(id) != nil
}

func ApplyState(s *State) (msgs []*pb.EventMessage, action history.Action, err error) {
	return s.apply(false, false)
}

func ApplyStateFast(s *State) (msgs []*pb.EventMessage, action history.Action, err error) {
	return s.apply(true, false)
}

func ApplyStateFastOne(s *State) (msgs []*pb.EventMessage, action history.Action, err error) {
	return s.apply(true, true)
}

func (s *State) apply(fast, one bool) (msgs []*pb.EventMessage, action history.Action, err error) {
	if s.parent != nil && (s.parent.parent != nil || fast) {
		s.intermediateApply()
		if one {
			return
		}
		return s.parent.apply(fast, one)
	}
	if fast {
		return
	}
	st := time.Now()
	if s.parent != nil && s.changeId != "" {
		s.parent.changeId = s.changeId
	}
	if s.parent != nil && s.details != nil {
		s.parent.details = s.details
	}
	if !fast {
		if err = s.normalize(); err != nil {
			return
		}
	}
	var (
		inUse       = make(map[string]struct{})
		affectedIds = make([]string, 0, len(s.blocks))
		newBlocks   []*model.Block
	)

	s.Iterate(func(b simple.Block) (isContinue bool) {
		id := b.Model().Id
		inUse[id] = struct{}{}
		if _, ok := s.blocks[id]; ok {
			affectedIds = append(affectedIds, id)
		}
		return true
	})

	flushNewBlocks := func() {
		if len(newBlocks) > 0 {
			msgs = append(msgs, &pb.EventMessage{
				Value: &pb.EventMessageValueOfBlockAdd{
					BlockAdd: &pb.EventBlockAdd{
						Blocks: newBlocks,
					},
				},
			})
		}
		newBlocks = nil
	}

	// new and changed blocks
	// we need to create events with affectedIds order for correct changes generation
	for _, id := range affectedIds {
		orig := s.PickOrigin(id)
		if orig == nil {
			bc := s.blocks[id].Copy()
			newBlocks = append(newBlocks, bc.Model())
			action.Add = append(action.Add, bc)
		} else {
			flushNewBlocks()
			b := s.blocks[id]
			diff, err := orig.Diff(b)
			if err != nil {
				return nil, history.Action{}, err
			}
			if len(diff) > 0 {
				msgs = append(msgs, diff...)
				if file := orig.Model().GetFile(); file != nil {
					if file.State == model.BlockContentFile_Uploading {
						file.State = model.BlockContentFile_Empty
					}
				}
				action.Change = append(action.Change, history.Change{
					Before: orig.Copy(),
					After:  b.Copy(),
				})
			}
		}
	}
	flushNewBlocks()

	// removed blocks
	var (
		toRemove []string
		bm       map[string]simple.Block
	)
	if s.parent != nil {
		bm = s.parent.blocks
	} else {
		bm = s.blocks
	}
	for id := range bm {
		if _, ok := inUse[id]; !ok {
			toRemove = append(toRemove, id)
		}
	}
	if len(toRemove) > 0 {
		msgs = append(msgs, &pb.EventMessage{
			Value: &pb.EventMessageValueOfBlockDelete{
				BlockDelete: &pb.EventBlockDelete{BlockIds: toRemove},
			},
		})
	}
	// generate changes
	s.fillChanges(msgs)

	// apply to parent
	for _, b := range s.blocks {
		if s.parent != nil {
			s.parent.blocks[b.Model().Id] = b
		}
	}
	for _, id := range toRemove {
		action.Remove = append(action.Remove, s.Pick(id))
		if s.parent != nil {
			delete(s.parent.blocks, id)
		}
	}
	if s.parent != nil {
		s.parent.changes = s.changes
	}
	log.Infof("middle: state apply: %d affected; %d for remove; %d copied; %d changes; for a %v", len(affectedIds), len(toRemove), len(s.blocks), len(s.changes), time.Since(st))
	return
}

func (s *State) intermediateApply() {
	st := time.Now()
	if s.changeId != "" {
		s.parent.changeId = s.changeId
	}
	for _, b := range s.blocks {
		s.parent.Set(b)
	}
	if s.details != nil {
		s.parent.details = s.details
	}
	s.parent.changes = append(s.parent.changes, s.changes...)
	log.Infof("middle: state intermediate apply: %d to update; for a %v", len(s.blocks), time.Since(st))
	return
}

func (s *State) Blocks() []*model.Block {
	if s.Pick(s.RootId()) == nil {
		return nil
	}
	return s.fillSlice(s.RootId(), make([]*model.Block, 0, len(s.blocks)))
}

func (s *State) fillSlice(id string, blocks []*model.Block) []*model.Block {
	b := s.Pick(id)
	blocks = append(blocks, b.Copy().Model())
	for _, chId := range b.Model().ChildrenIds {
		blocks = s.fillSlice(chId, blocks)
	}
	return blocks
}

func (s *State) String() (res string) {
	buf := bytes.NewBuffer(nil)
	s.writeString(buf, 0, s.RootId())
	return buf.String()
}

func (s *State) writeString(buf *bytes.Buffer, l int, id string) {
	b := s.Pick(id)
	buf.WriteString(strings.Repeat("\t", l))
	if b == nil {
		buf.WriteString(id)
		buf.WriteString(" MISSING")
	} else {
		buf.WriteString(b.String())
	}
	buf.WriteString("\n")
	if b != nil {
		for _, cid := range b.Model().ChildrenIds {
			s.writeString(buf, l+1, cid)
		}
	}
}

func (s *State) SetDetails(d *types.Struct) *State {
	s.details = d
	return s
}

func (s *State) Details() *types.Struct {
	if s.details == nil && s.parent != nil {
		return s.parent.Details()
	}
	return s.details
}
