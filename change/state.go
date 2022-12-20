package change

import (
	"time"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
)

func NewStateCache() *stateCache {
	return &stateCache{
		states: make(map[string]struct {
			refs  int
			state *state.State
		}),
	}
}

type stateCache struct {
	states map[string]struct {
		refs  int
		state *state.State
	}
}

func (sc *stateCache) Set(id string, s *state.State, refs int) {
	sc.states[id] = struct {
		refs  int
		state *state.State
	}{refs: refs, state: s}
}

func (sc *stateCache) Get(id string) *state.State {
	item := sc.states[id]
	item.refs--
	if item.refs == 0 {
		delete(sc.states, id)
	} else {
		sc.states[id] = item
	}
	return item.state
}

// Simple implementation hopes for CRDT and ignores errors. No merge
func BuildStateSimpleCRDT(root *state.State, t *Tree) (s *state.State, err error) {
	var (
		startId    string
		applyRoot  bool
		st         = time.Now()
		lastChange *Change
		count      int
	)
	if startId = root.ChangeId(); startId == "" {
		startId = t.RootId()
		applyRoot = true
	}

	t.Iterate(startId, func(c *Change) (isContinue bool) {
		count++
		lastChange = c
		if startId == c.Id {
			s = root.NewState()
			if applyRoot {
				s.ApplyChangeIgnoreErr(c.Change.Content...)
				s.SetChangeId(c.Id)
				s.AddFileKeys(c.FileKeys...)
			}
			return true
		}
		ns := s.NewState()
		ns.ApplyChangeIgnoreErr(c.Change.Content...)
		ns.SetChangeId(c.Id)
		ns.AddFileKeys(c.FileKeys...)
		_, _, err = state.ApplyStateFastOne(ns)
		if err != nil {
			return false
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	if lastChange != nil {
		s.SetLastModified(lastChange.Timestamp, lastChange.Account)
	}

	log.Infof("build state (crdt): changes: %d; dur: %v;", count, time.Since(st))
	return s, err
}

func merge(t *Tree, states ...*state.State) (s *state.State) {
	for _, st := range states {
		if s == nil {
			s = st
		} else {
			s = s.Merge(st)
		}
	}
	return s
}
