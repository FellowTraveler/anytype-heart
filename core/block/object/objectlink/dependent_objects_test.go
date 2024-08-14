package objectlink

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/anyproto/anytype-heart/core/block/editor/state"
	"github.com/anyproto/anytype-heart/core/block/simple"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

type fakeConverter struct {
}

func (f *fakeConverter) GetRelationIdByKey(ctx context.Context, key domain.RelationKey) (id string, err error) {
	return fakeDerivedID(key.String()), nil
}

func (f *fakeConverter) GetTypeIdByKey(ctx context.Context, key domain.TypeKey) (id string, err error) {
	return fakeDerivedID(key.String()), nil
}

func fakeDerivedID(key string) string {
	return fmt.Sprintf("derivedFrom(%s)", key)
}

func TestState_DepSmartIdsLinks(t *testing.T) {
	// given
	stateWithLinks := state.NewDoc("root", map[string]simple.Block{
		"root": simple.New(&model.Block{
			Id:          "root",
			ChildrenIds: []string{"childBlock", "childBlock2", "childBlock3"},
		}),
		"childBlock": simple.New(&model.Block{Id: "childBlock",
			Content: &model.BlockContentOfText{
				Text: &model.BlockContentText{Marks: &model.BlockContentTextMarks{
					Marks: []*model.BlockContentTextMark{
						{
							Range: &model.Range{
								From: 0,
								To:   8,
							},
							Type:  model.BlockContentTextMark_Object,
							Param: "objectID",
						},
						{
							Range: &model.Range{
								From: 9,
								To:   19,
							},
							Type:  model.BlockContentTextMark_Mention,
							Param: "objectID2",
						},
					},
				}},
			}}),
		"childBlock2": simple.New(&model.Block{Id: "childBlock2",
			Content: &model.BlockContentOfBookmark{
				Bookmark: &model.BlockContentBookmark{
					TargetObjectId: "objectID3",
				},
			}}),
		"childBlock3": simple.New(&model.Block{Id: "childBlock3",
			Content: &model.BlockContentOfLink{
				Link: &model.BlockContentLink{
					TargetBlockId: "objectID4",
				},
			}}),
	}).(*state.State)
	converter := &fakeConverter{}

	t.Run("block option is turned on: get ids from blocks", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{Blocks: true})
		assert.Len(t, objectIDs, 4)
	})

	t.Run("all options are turned off", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{})
		assert.Len(t, objectIDs, 0)
	})
}

func TestState_DepSmartIdsLinksAndRelations(t *testing.T) {
	// given
	stateWithLinks := state.NewDoc("root", map[string]simple.Block{
		"root": simple.New(&model.Block{
			Id:          "root",
			ChildrenIds: []string{"childBlock", "childBlock2", "childBlock3", "dataview", "image"},
		}),
		"childBlock": simple.New(&model.Block{Id: "childBlock",
			Content: &model.BlockContentOfText{
				Text: &model.BlockContentText{Marks: &model.BlockContentTextMarks{
					Marks: []*model.BlockContentTextMark{
						{
							Range: &model.Range{
								From: 0,
								To:   8,
							},
							Type:  model.BlockContentTextMark_Object,
							Param: "objectID",
						},
						{
							Range: &model.Range{
								From: 9,
								To:   19,
							},
							Type:  model.BlockContentTextMark_Mention,
							Param: "objectID2",
						},
					},
				}},
			}}),
		"childBlock2": simple.New(&model.Block{Id: "childBlock2",
			Content: &model.BlockContentOfBookmark{
				Bookmark: &model.BlockContentBookmark{
					TargetObjectId: "objectID3",
				},
			}}),
		"childBlock3": simple.New(&model.Block{Id: "childBlock3",
			Content: &model.BlockContentOfLink{
				Link: &model.BlockContentLink{
					TargetBlockId: "objectID4",
				},
			}}),
		"dataview": simple.New(&model.Block{Id: "dataview",
			Content: &model.BlockContentOfDataview{
				Dataview: &model.BlockContentDataview{
					Views: []*model.BlockContentDataviewView{{
						Id:                  "Today's tasks",
						DefaultObjectTypeId: "task",
						DefaultTemplateId:   "Task with a picture",
					}},
					TargetObjectId: "taskTracker",
				},
			}}),
		"image": simple.New(&model.Block{Id: "image",
			Content: &model.BlockContentOfFile{
				File: &model.BlockContentFile{
					TargetObjectId: "image with cute kitten",
				},
			}}),
	}).(*state.State)
	converter := &fakeConverter{}

	relations := []*model.RelationLink{
		{
			Key:    "relation1",
			Format: model.RelationFormat_file,
		},
		{
			Key:    "relation2",
			Format: model.RelationFormat_tag,
		},
		{
			Key:    "relation3",
			Format: model.RelationFormat_status,
		},
		{
			Key:    "relation4",
			Format: model.RelationFormat_object,
		},
	}
	stateWithLinks.AddRelationLinks(relations...)

	t.Run("blocks option is turned on: get ids from blocks", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{Blocks: true})
		assert.Len(t, objectIDs, 8)
	})

	t.Run("dataview only target option is turned on: get only target from blocks", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{Blocks: true, DataviewBlockOnlyTarget: true})
		assert.Len(t, objectIDs, 6)
	})

	t.Run("blocks option and relations options are turned on: get ids from blocks and relations", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{Blocks: true, Relations: true})
		assert.Len(t, objectIDs, 12) // 8 links + 4 relations
	})
}

func TestState_DepSmartIdsLinksDetailsAndRelations(t *testing.T) {
	// given
	stateWithLinks := state.NewDoc("root", map[string]simple.Block{
		"root": simple.New(&model.Block{
			Id:          "root",
			ChildrenIds: []string{"childBlock", "childBlock2", "childBlock3"},
		}),
		"childBlock": simple.New(&model.Block{Id: "childBlock",
			Content: &model.BlockContentOfText{
				Text: &model.BlockContentText{Marks: &model.BlockContentTextMarks{
					Marks: []*model.BlockContentTextMark{
						{
							Range: &model.Range{
								From: 0,
								To:   8,
							},
							Type:  model.BlockContentTextMark_Object,
							Param: "objectID",
						},
						{
							Range: &model.Range{
								From: 9,
								To:   19,
							},
							Type:  model.BlockContentTextMark_Mention,
							Param: "objectID2",
						},
					},
				}},
			}}),
		"childBlock2": simple.New(&model.Block{Id: "childBlock2",
			Content: &model.BlockContentOfBookmark{
				Bookmark: &model.BlockContentBookmark{
					TargetObjectId: "objectID3",
				},
			}}),
		"childBlock3": simple.New(&model.Block{Id: "childBlock3",
			Content: &model.BlockContentOfLink{
				Link: &model.BlockContentLink{
					TargetBlockId: "objectID4",
				},
			}}),
	}).(*state.State)
	converter := &fakeConverter{}

	relations := []*model.RelationLink{
		{
			Key:    "relation1",
			Format: model.RelationFormat_file,
		},
		{
			Key:    "relation2",
			Format: model.RelationFormat_tag,
		},
		{
			Key:    "relation3",
			Format: model.RelationFormat_status,
		},
		{
			Key:    "relation4",
			Format: model.RelationFormat_object,
		},
		{
			Key:    "relation5",
			Format: model.RelationFormat_date,
		},
	}
	stateWithLinks.AddRelationLinks(relations...)
	stateWithLinks.SetDetail("relation1", pbtypes.String("file"))
	stateWithLinks.SetDetail("relation2", pbtypes.String("option1"))
	stateWithLinks.SetDetail("relation3", pbtypes.String("option2"))
	stateWithLinks.SetDetail("relation4", pbtypes.String("option3"))
	stateWithLinks.SetDetail("relation5", pbtypes.Int64(time.Now().Unix()))

	t.Run("blocks option is turned on: get ids from blocks", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{Blocks: true})
		assert.Len(t, objectIDs, 4) // links
	})
	t.Run("blocks option and relations option are turned on: get ids from blocks and relations", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{Blocks: true, Relations: true})
		assert.Len(t, objectIDs, 9) // 4 links + 5 relations
	})
	t.Run("blocks, relations and details option are turned on: get ids from blocks, relations and details", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{Blocks: true, Relations: true, Details: true})
		assert.Len(t, objectIDs, 14) // 4 links + 5 relations + 3 options + 1 fileID + 1 date
	})
}

func TestState_DepSmartIdsLinksCreatorModifierWorkspace(t *testing.T) {
	// given
	stateWithLinks := state.NewDoc("root", nil).(*state.State)
	relations := []*model.RelationLink{
		{
			Key:    "relation1",
			Format: model.RelationFormat_date,
		},
		{
			Key:    bundle.RelationKeyCreatedDate.String(),
			Format: model.RelationFormat_date,
		},
		{
			Key:    bundle.RelationKeyCreator.String(),
			Format: model.RelationFormat_object,
		},
		{
			Key:    bundle.RelationKeyLastModifiedBy.String(),
			Format: model.RelationFormat_object,
		},
	}
	stateWithLinks.AddRelationLinks(relations...)
	stateWithLinks.SetDetail("relation1", pbtypes.Int64(time.Now().Unix()))
	stateWithLinks.SetDetail(bundle.RelationKeyCreatedDate.String(), pbtypes.Int64(time.Now().Unix()))
	stateWithLinks.SetDetail(bundle.RelationKeyCreator.String(), pbtypes.String("creator"))
	stateWithLinks.SetDetail(bundle.RelationKeyLastModifiedBy.String(), pbtypes.String("lastModifiedBy"))
	converter := &fakeConverter{}

	t.Run("details option is turned on: get ids only from details", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{Details: true, CreatorModifierWorkspace: true})
		assert.Len(t, objectIDs, 3) // creator + lastModifiedBy + 1 date
	})

	t.Run("details and relations options are turned on: get ids from details and relations", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{Details: true, Relations: true, CreatorModifierWorkspace: true})
		assert.Len(t, objectIDs, 7) // 4 relations + creator + lastModifiedBy + 1 date
	})
}

func TestState_DepSmartIdsObjectTypes(t *testing.T) {
	// given
	stateWithLinks := state.NewDoc("root", nil).(*state.State)
	stateWithLinks.SetObjectTypeKey(bundle.TypeKeyPage)
	converter := &fakeConverter{}

	t.Run("all options are turned off", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{})
		assert.Len(t, objectIDs, 0)
	})
	t.Run("objTypes option is turned on, get only object types id", func(t *testing.T) {
		objectIDs := DependentObjectIDs(stateWithLinks, converter, Flags{Types: true})
		assert.Equal(t, []string{
			fakeDerivedID(bundle.TypeKeyPage.String()),
		}, objectIDs)
	})
}
