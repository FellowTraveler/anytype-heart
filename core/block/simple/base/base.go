package base

import (
	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/gogo/protobuf/types"
	"github.com/mohae/deepcopy"
)

func init() {
	simple.RegisterCreator(func(m *model.Block) simple.Block {
		if m.GetIcon() != nil {
			return NewIcon(m)
		}
		return nil
	})
	simple.RegisterFallback(func(m *model.Block) simple.Block {
		return NewBase(m)
	})
}

func NewBase(block *model.Block) simple.Block {
	return &Base{Block: block}
}

type Base struct {
	*model.Block
}

func (s *Base) Virtual() bool {
	return false
}

func (s *Base) Model() *model.Block {
	return s.Block
}

func (s *Base) Diff(block simple.Block) (msgs []*pb.EventMessage, err error) {
	m := block.Model()
	if m.IsArchived != s.IsArchived {
		m := &pb.EventMessage{Value: &pb.EventMessageValueOfBlockSetIsArchived{BlockSetIsArchived: &pb.EventBlockSetIsArchived{
			Id:         s.Id,
			IsArchived: m.IsArchived,
		}}}
		msgs = append(msgs, m)
	}

	if !stringSlicesEq(m.ChildrenIds, s.ChildrenIds) {
		m := &pb.EventMessage{Value: &pb.EventMessageValueOfBlockSetChildrenIds{BlockSetChildrenIds: &pb.EventBlockSetChildrenIds{
			Id:          s.Id,
			ChildrenIds: m.ChildrenIds,
		}}}
		msgs = append(msgs, m)
	}

	if s.Restrictions == nil {
		s.Restrictions = &model.BlockRestrictions{}
	}
	if m.Restrictions == nil {
		m.Restrictions = &model.BlockRestrictions{}
	}
	if *s.Restrictions != *m.Restrictions {
		m := &pb.EventMessage{Value: &pb.EventMessageValueOfBlockSetRestrictions{BlockSetRestrictions: &pb.EventBlockSetRestrictions{
			Id:           s.Id,
			Restrictions: m.Restrictions,
		}}}
		msgs = append(msgs, m)
	}
	if !fieldsEq(s.Fields, m.Fields) {
		m := &pb.EventMessage{Value: &pb.EventMessageValueOfBlockSetFields{BlockSetFields: &pb.EventBlockSetFields{
			Id:     s.Id,
			Fields: m.Fields,
		}}}
		msgs = append(msgs, m)
	}
	return
}

func (b *Base) Copy() simple.Block {
	return NewBase(deepcopy.Copy(b.Model()).(*model.Block))
}

func stringSlicesEq(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func fieldsEq(f1, f2 *types.Struct) bool {
	if f1 == nil {
		f1 = &types.Struct{}
	}
	if f2 == nil {
		f2 = &types.Struct{}
	}
	return f1.Compare(f2) == 0
}
