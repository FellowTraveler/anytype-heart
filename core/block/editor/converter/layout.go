package converter

import (
	"fmt"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
)

type LayoutConverter struct {
	objectStore objectstore.ObjectStore
}

func NewLayoutConverter(objectStore objectstore.ObjectStore) *LayoutConverter {
	return &LayoutConverter{
		objectStore: objectStore,
	}
}

func (c *LayoutConverter) Convert(st *state.State, fromLayout, toLayout model.ObjectTypeLayout) error {
	if fromLayout == model.ObjectType_note && toLayout == model.ObjectType_collection {
		return c.fromNoteToCollection(st)
	}
	if fromLayout == model.ObjectType_set && toLayout == model.ObjectType_collection {
		return c.fromSetToCollection(st)
	}
	if toLayout == model.ObjectType_collection {
		return c.fromAnyToCollection(st)
	}

	if toLayout == model.ObjectType_note {
		return c.fromAnyToNote(st)
	}
	if fromLayout == model.ObjectType_note {
		return c.fromNoteToAny(st)
	}
	return nil
}

func (c *LayoutConverter) fromSetToCollection(st *state.State) error {
	dvBlock := st.Get(template.DataviewBlockId)
	if dvBlock == nil {
		return fmt.Errorf("dataview block is not found")
	}
	details := st.Details()
	typesFromSet := pbtypes.GetStringList(details, bundle.RelationKeySetOf.String())

	c.removeRelationSetOf(st)

	dvBlock.Model().GetDataview().IsCollection = true

	recs, _, qErr := c.objectStore.Query(nil, database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyType.String(),
				Condition:   model.BlockContentDataviewFilter_In,
				Value:       pbtypes.StringList(typesFromSet),
			},
		},
	})
	if qErr != nil {
		return fmt.Errorf("can't get records for collection: %w", qErr)
	}
	ids := make([]string, 0, len(recs))
	for _, r := range recs {
		ids = append(ids, pbtypes.GetString(r.Details, bundle.RelationKeyId.String()))
	}
	st.StoreSlice(template.CollectionStoreKey, ids)
	return nil
}

func (c *LayoutConverter) fromNoteToCollection(st *state.State) error {
	if err := c.fromAnyToNote(st); err != nil {
		return err
	}

	return c.fromAnyToCollection(st)
}

func (c *LayoutConverter) fromAnyToCollection(st *state.State) error {
	blockContent := template.MakeCollectionDataviewContent()
	template.InitTemplate(st, template.WithDataview(*blockContent, false))
	return nil
}

func (c *LayoutConverter) fromNoteToAny(st *state.State) error {
	if name, ok := st.Details().Fields[bundle.RelationKeyName.String()]; !ok || name.GetStringValue() == "" {
		textBlock, err := st.GetFirstTextBlock()
		if err != nil {
			return err
		}
		if textBlock == nil {
			return nil
		}
		st.SetDetail(bundle.RelationKeyName.String(), pbtypes.String(textBlock.Model().GetText().GetText()))

		for _, id := range textBlock.Model().ChildrenIds {
			st.Unlink(id)
		}
		err = st.InsertTo(textBlock.Model().Id, model.Block_Bottom, textBlock.Model().ChildrenIds...)
		if err != nil {
			return fmt.Errorf("insert children: %w", err)
		}
		st.Unlink(textBlock.Model().Id)
	}
	return nil
}

func (c *LayoutConverter) fromAnyToNote(st *state.State) error {
	if name, ok := st.Details().Fields[bundle.RelationKeyName.String()]; ok && name.GetStringValue() != "" {
		newBlock := simple.New(&model.Block{
			Content: &model.BlockContentOfText{
				Text: &model.BlockContentText{Text: name.GetStringValue()},
			},
		})
		st.Add(newBlock)

		if err := st.InsertTo(template.HeaderLayoutId, model.Block_Bottom, newBlock.Model().Id); err != nil {
			return err
		}

		st.RemoveDetail(bundle.RelationKeyName.String())
	}
	return nil
}

func (c *LayoutConverter) removeRelationSetOf(st *state.State) {
	st.RemoveDetail(bundle.RelationKeySetOf.String())

	fr := pbtypes.GetStringList(st.Details(), bundle.RelationKeyFeaturedRelations.String())
	fr = slice.Remove(fr, bundle.RelationKeySetOf.String())
	st.SetDetail(bundle.RelationKeyFeaturedRelations.String(), pbtypes.StringList(fr))
}
