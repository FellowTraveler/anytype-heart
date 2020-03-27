package state

import (
	"time"

	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-middleware/core/block/history"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("anytype-state")

type Doc interface {
	RootId() string
	NewState() *State
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
	parent   *State
	blocks   map[string]simple.Block
	rootId   string
	toRemove []string
	newIds   []string
}

func (s *State) RootId() string {
	return s.rootId
}

func (s *State) NewState() *State {
	return &State{parent: s, blocks: make(map[string]simple.Block), rootId: s.rootId}
}

func (s *State) Add(b simple.Block) (ok bool) {
	id := b.Model().Id
	if s.Pick(id) == nil {
		s.blocks[id] = b
		s.newIds = append(s.newIds, id)
		return true
	}
	return false
}

func (s *State) Set(b simple.Block) {
	s.blocks[b.Model().Id] = b
}

func (s *State) Get(id string) (b simple.Block) {
	if slice.FindPos(s.toRemove, id) != -1 {
		return nil
	}
	if b = s.blocks[id]; b != nil {
		return
	}
	if s.parent != nil {
		if b = s.parent.Get(id); b != nil {
			b = b.Copy()
			s.blocks[id] = b
			return
		}
	}
	return
}

func (s *State) Pick(id string) (b simple.Block) {
	if slice.FindPos(s.toRemove, id) != -1 {
		return nil
	}
	if b = s.blocks[id]; b != nil {
		return
	}
	if s.parent != nil {
		return s.parent.Pick(id)
	}
	return
}

func (s *State) PickOrigin(id string) (b simple.Block) {
	if s.parent != nil {
		return s.parent.Pick(id)
	}
	return
}

func (s *State) Remove(id string) (ok bool) {
	if slice.FindPos(s.toRemove, id) != -1 {
		return false
	}
	if s.Pick(id) != nil {
		if _, ok = s.blocks[id]; ok {
			delete(s.blocks, id)
		}
		s.toRemove = append(s.toRemove, id)
		s.Unlink(id)
		return true
	}
	return
}

func (s *State) Unlink(id string) (ok bool) {
	s.Iterate(func(b simple.Block) bool {
		if slice.FindPos(b.Model().ChildrenIds, id) != -1 {
			parent := s.Get(b.Model().Id).Model()
			parent.ChildrenIds = slice.Remove(parent.ChildrenIds, id)
			ok = true
			return false
		}
		return true
	})
	return
}

func (s *State) Append(targetId string, id string) (ok bool) {
	//s.Iterate(func(b simple.Block) bool {
	//	if b.Model().Id == targetId {
			parent := s.Get(targetId).Model()
			parent.ChildrenIds = append(parent.ChildrenIds, id)//slice.Remove(parent.ChildrenIds, id)
			ok = true
			return ok
	//	}
	//	return true
	//})
	return
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

func (s *State) Iterate(f func(b simple.Block) (isContinue bool)) {
	for _, newId := range s.newIds {
		if ! f(s.blocks[newId]) {
			return
		}
	}
	if s.parent == nil {
		for _, b := range s.blocks {
			if ! f(b) {
				return
			}
		}
	} else {
		s.parent.Iterate(func(b simple.Block) (isContinue bool) {
			if s.Exists(b.Model().Id) {
				return f(s.Pick(b.Model().Id))
			}
			return true
		})
		return
	}
}

func (s *State) Exists(id string) (ok bool) {
	return s.Pick(id) != nil
}

func ApplyState(s *State) (msgs []*pb.EventMessage, action history.Action, err error) {
	return s.apply()
}

func (s *State) apply() (msgs []*pb.EventMessage, action history.Action, err error) {
	st := time.Now()
	s.normalize()
	var toSave []*model.Block
	var newBlocks []*model.Block
	for id, b := range s.blocks {
		if slice.FindPos(s.toRemove, id) != -1 {
			continue
		}
		if err := s.validateBlock(b); err != nil {
			return nil, history.Action{}, err
		}
		orig := s.PickOrigin(id)
		if orig == nil {
			newBlocks = append(newBlocks, b.Model())
			toSave = append(toSave, b.Model())
			action.Add = append(action.Add, b)
			continue
		}

		diff, err := orig.Diff(b)
		if err != nil {
			return nil, history.Action{}, err
		}
		if len(diff) > 0 {
			toSave = append(toSave, b.Model())
			msgs = append(msgs, diff...)
			if file := orig.Model().GetFile(); file != nil {
				if file.State == model.BlockContentFile_Uploading {
					file.State = model.BlockContentFile_Empty
				}
			}
			action.Change = append(action.Change, history.Change{
				Before: orig,
				After:  b,
			})
		}
	}
	if len(s.toRemove) > 0 {
		msgs = append(msgs, &pb.EventMessage{
			Value: &pb.EventMessageValueOfBlockDelete{
				BlockDelete: &pb.EventBlockDelete{BlockIds: s.toRemove},
			},
		})
	}
	if len(newBlocks) > 0 {
		msgs = append(msgs, &pb.EventMessage{
			Value: &pb.EventMessageValueOfBlockAdd{
				BlockAdd: &pb.EventBlockAdd{
					Blocks: newBlocks,
				},
			},
		})
	}
	for _, b := range s.blocks {
		if s.parent != nil {
			s.parent.blocks[b.Model().Id] = b
		}
	}
	for _, id := range s.toRemove {
		if old := s.PickOrigin(id); old != nil {
			action.Remove = append(action.Remove, old)
		}
		if s.parent != nil {
			delete(s.parent.blocks, id)
		}
	}
	log.Infof("middle: state apply: %d for save; %d for remove; %d copied; for a %v", len(toSave), len(s.toRemove), len(s.blocks), time.Since(st))
	return
}

func (s *State) Blocks() []*model.Block {
	res := make([]*model.Block, 0, len(s.blocks))
	for _, b := range s.blocks {
		res = append(res, b.Copy().Model())
	}
	return res
}
