package block

import (
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/dataview"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/session"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
)

func (s *Service) AddDataviewFilter(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	filter *model.BlockContentDataviewFilter,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.AddFilter(viewId, filter)
	})
}

func (s *Service) RemoveDataviewFilters(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	filterIDs []string,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.RemoveFilters(viewId, filterIDs)
	})
}

func (s *Service) ReplaceDataviewFilter(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	filterID string,
	filter *model.BlockContentDataviewFilter,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.ReplaceFilter(viewId, filterID, filter)
	})
}

func (s *Service) ReorderDataviewFilters(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	filterIDs []string,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.ReorderFilters(viewId, filterIDs)
	})
}

func (s *Service) AddDataviewSort(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	sort *model.BlockContentDataviewSort,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.AddSort(viewId, sort)
	})
}

func (s *Service) RemoveDataviewSorts(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	relationKeys []string,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.RemoveSorts(viewId, relationKeys)
	})
}

func (s *Service) ReplaceDataviewSort(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	relationKey string,
	sort *model.BlockContentDataviewSort,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.ReplaceSort(viewId, relationKey, sort)
	})
}

func (s *Service) ReorderDataviewSorts(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	relationKeys []string,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.ReorderSorts(viewId, relationKeys)
	})
}

func (s *Service) AddDataviewViewRelation(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	relation *model.BlockContentDataviewRelation,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.AddViewRelation(viewId, relation)
	})
}

func (s *Service) RemoveDataviewViewRelations(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	relationKeys []string,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.RemoveViewRelations(viewId, relationKeys)
	})
}

func (s *Service) ReplaceDataviewViewRelation(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	relationKey string,
	relation *model.BlockContentDataviewRelation,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.ReplaceViewRelation(viewId, relationKey, relation)
	})
}

func (s *Service) ReorderDataviewViewRelations(
	ctx *session.Context,
	contextId string,
	blockId string,
	viewId string,
	relationKeys []string,
) (err error) {
	return DoStateCtx(s, ctx, contextId, func(s *state.State, d dataview.Dataview) error {
		dv, err := d.GetDataviewBlock(s, blockId)
		if err != nil {
			return err
		}

		return dv.ReorderViewRelations(viewId, relationKeys)
	})
}
