package basic

import (
	"fmt"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/globalsign/mgo/bson"
)

type StateTransformer struct {
	*state.State
}

func NewStateTransformer(s *state.State) StateTransformer {
	return StateTransformer{
		State: s,
	}
}

func (s StateTransformer) CreateBlock(groupId string, req pb.RpcBlockCreateRequest) (id string, err error) {
	if req.TargetId != "" {
		if s.IsChild(template.HeaderLayoutId, req.TargetId) {
			req.Position = model.Block_Bottom
			req.TargetId = template.HeaderLayoutId
		}
	}
	if req.Block.GetContent() == nil {
		err = fmt.Errorf("no block content")
		return
	}
	req.Block.Id = ""
	block := simple.New(req.Block)
	block.Model().ChildrenIds = nil
	err = block.Validate()
	if err != nil {
		return
	}
	s.Add(block)
	if err = s.InsertTo(req.TargetId, req.Position, block.Model().Id); err != nil {
		return
	}
	return block.Model().Id, nil
}

func (s StateTransformer) CutBlocks(blockIds []string) (blocks []simple.Block) {
	var uniqMap = make(map[string]struct{})
	for _, bId := range blockIds {
		b := s.Pick(bId)
		if b != nil {
			descendants := s.getAllDescendants(uniqMap, b.Copy(), []simple.Block{})
			blocks = append(blocks, descendants...)
			s.Unlink(b.Model().Id)
		}
	}
	return blocks
}

func (s StateTransformer) PasteBlocks(blocks []simple.Block) error {
	// TODO: use new functions for this
	childIdsRewrite := make(map[string]string)
	for _, b := range blocks {
		for i, cId := range b.Model().ChildrenIds {
			newId := bson.NewObjectId().Hex()
			childIdsRewrite[cId] = newId
			b.Model().ChildrenIds[i] = newId
		}
	}
	for _, b := range blocks {
		var child bool
		if newId, ok := childIdsRewrite[b.Model().Id]; ok {
			b.Model().Id = newId
			child = true
		} else {
			b.Model().Id = bson.NewObjectId().Hex()
		}
		s.Add(b)
		if !child {
			err := s.InsertTo("", model.Block_Inner, b.Model().Id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s StateTransformer) getAllDescendants(uniqMap map[string]struct{}, block simple.Block, blocks []simple.Block) []simple.Block {
	if _, ok := uniqMap[block.Model().Id]; ok {
		return blocks
	}
	blocks = append(blocks, block)
	uniqMap[block.Model().Id] = struct{}{}
	for _, cId := range block.Model().ChildrenIds {
		blocks = s.getAllDescendants(uniqMap, s.Pick(cId).Copy(), blocks)
	}
	return blocks
}
