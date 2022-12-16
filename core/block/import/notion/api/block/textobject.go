package block

import (
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gogo/protobuf/types"

	"github.com/anytypeio/go-anytype-middleware/core/block/import/converter"
	"github.com/anytypeio/go-anytype-middleware/core/block/import/notion/api"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	textUtil "github.com/anytypeio/go-anytype-middleware/util/text"
)

const DateMentionTimeFormat = "2006-01-02"

type ChildrenMapper interface {
	MapChildren(req *MapRequest) *MapResponse
}

type TextObject struct {
	RichText []api.RichText `json:"rich_text"`
	Color    string         `json:"color"`
}

func (t *TextObject) GetTextBlocks(style model.BlockContentTextStyle, childIds []string, req *MapRequest) *MapResponse {
	marks := []*model.BlockContentTextMark{}
	id := bson.NewObjectId().Hex()
	allBlocks := make([]*model.Block, 0)
	allIds := make([]string, 0)
	var (
		text      strings.Builder
		relations []*converter.Relation
		details   map[string]*types.Value
	)
	for _, rt := range t.RichText {
		if rt.Type == api.Text {
			marks = append(marks, t.handleTextType(rt, &text, req.NotionPageIdsToAnytype, req.NotionDatabaseIdsToAnytype)...)
		}
		if rt.Type == api.Mention {
			// Return Relation block for Date mention
			if rt.Mention.Type == api.Date {
				var (
					relationBlock   *model.Block
					relationBlockID string
					relation        *converter.Relation
				)
				relationBlock, relation, details, relationBlockID = t.handleDateMention(rt, &text)
				if relationBlock != nil {
					allBlocks = append(allBlocks, relationBlock)
					allIds = append(allIds, relationBlockID)
					relations = append(relations, relation)
				}
				continue
			}
			marks = append(marks, t.handleMentionType(rt, &text, req)...)
		}
		if rt.Type == api.Equation {
			eqBlock := rt.Equation.HandleEquation()
			allBlocks = append(allBlocks, eqBlock)
			allIds = append(allIds, eqBlock.Id)
		}
	}
	var backgroundColor, textColor string
	if strings.Contains(t.Color, api.NotionBackgroundColorSuffix) {
		backgroundColor = api.NotionColorToAnytype[t.Color]
	} else {
		textColor = api.NotionColorToAnytype[t.Color]
	}

	if t.isNotTextBlocks() {
		return &MapResponse{
			Blocks:    allBlocks,
			Relations: relations,
			Details:   details,
			BlockIDs:  allIds,
		}
	}
	allBlocks = append(allBlocks, &model.Block{
		Id:              id,
		ChildrenIds:     childIds,
		BackgroundColor: backgroundColor,
		Content: &model.BlockContentOfText{
			Text: &model.BlockContentText{
				Text:    text.String(),
				Style:   style,
				Marks:   &model.BlockContentTextMarks{Marks: marks},
				Checked: false,
				Color:   textColor,
			},
		},
	})
	for _, b := range allBlocks {
		allIds = append(allIds, b.Id)
	}
	return &MapResponse{
		Blocks:    allBlocks,
		Relations: relations,
		Details:   details,
		BlockIDs:  allIds,
	}
}

func (t *TextObject) handleTextType(rt api.RichText,
	text *strings.Builder,
	notionPageIdsToAnytype,
	notionDatabaseIdsToAnytype map[string]string) []*model.BlockContentTextMark {
	marks := []*model.BlockContentTextMark{}
	from := textUtil.UTF16RuneCountString(text.String())
	if rt.Text != nil && rt.Text.Link != nil && rt.Text.Link.Url != "" {
		text.WriteString(rt.Text.Content)
	} else {
		text.WriteString(rt.PlainText)
	}
	to := textUtil.UTF16RuneCountString(text.String())
	if rt.Text != nil && rt.Text.Link != nil && rt.Text.Link.Url != "" {
		url := strings.Trim(rt.Text.Link.Url, "/")
		if databaseID, ok := notionDatabaseIdsToAnytype[url]; ok {
			url = databaseID
		}
		if pageID, ok := notionPageIdsToAnytype[url]; ok {
			url = pageID
		}
		marks = append(marks, &model.BlockContentTextMark{
			Range: &model.Range{
				From: int32(from),
				To:   int32(to),
			},
			Type:  model.BlockContentTextMark_Link,
			Param: url,
		})
	}
	marks = append(marks, rt.BuildMarkdownFromAnnotations(int32(from), int32(to))...)
	return marks
}

func (t *TextObject) handleMentionType(rt api.RichText, text *strings.Builder, req *MapRequest) []*model.BlockContentTextMark {
	if rt.Mention.Type == api.UserMention {
		return t.handleUserMention(rt, text)
	}
	if rt.Mention.Type == api.Database {
		return t.handleDatabaseMention(rt, text, req.NotionDatabaseIdsToAnytype, req.DatabaseNameToID)
	}
	if rt.Mention.Type == api.Page {
		return t.handlePageMention(rt, text, req.NotionPageIdsToAnytype, req.PageNameToID)
	}
	if rt.Mention.Type == api.LinkPreview {
		return t.handleLinkPreviewMention(rt, text)
	}
	return nil
}

func (t *TextObject) handleUserMention(rt api.RichText, text *strings.Builder) []*model.BlockContentTextMark {
	from := textUtil.UTF16RuneCountString(text.String())
	text.WriteString(rt.PlainText)
	to := textUtil.UTF16RuneCountString(text.String())
	return rt.BuildMarkdownFromAnnotations(int32(from), int32(to))
}

func (t *TextObject) handleDatabaseMention(rt api.RichText,
	text *strings.Builder,
	notionDatabaseIdsToAnytype, databaseNameToID map[string]string) []*model.BlockContentTextMark {
	if notionDatabaseIdsToAnytype == nil {
		return nil
	}
	from := textUtil.UTF16RuneCountString(text.String())
	text.WriteString(databaseNameToID[rt.Mention.Database.ID])
	to := textUtil.UTF16RuneCountString(text.String())
	marks := rt.BuildMarkdownFromAnnotations(int32(from), int32(to))
	marks = append(marks, &model.BlockContentTextMark{
		Range: &model.Range{
			From: int32(from),
			To:   int32(to),
		},
		Type:  model.BlockContentTextMark_Mention,
		Param: notionDatabaseIdsToAnytype[rt.Mention.Database.ID],
	})
	return marks
}

func (t *TextObject) handlePageMention(rt api.RichText,
	text *strings.Builder,
	notionPageIdsToAnytype, pageNameToID map[string]string) []*model.BlockContentTextMark {
	if notionPageIdsToAnytype == nil {
		return nil
	}
	from := textUtil.UTF16RuneCountString(text.String())
	text.WriteString(pageNameToID[rt.Mention.Page.ID])
	to := textUtil.UTF16RuneCountString(text.String())
	marks := rt.BuildMarkdownFromAnnotations(int32(from), int32(to))
	marks = append(marks, &model.BlockContentTextMark{
		Range: &model.Range{
			From: int32(from),
			To:   int32(to),
		},
		Type:  model.BlockContentTextMark_Mention,
		Param: notionPageIdsToAnytype[rt.Mention.Page.ID],
	})
	return marks
}

func (t *TextObject) handleDateMention(rt api.RichText, text *strings.Builder) (*model.Block, *converter.Relation, map[string]*types.Value, string) {
	var textDate string
	if rt.Mention.Date.Start != "" {
		textDate = rt.Mention.Date.Start
	}
	if rt.Mention.Date.End != "" {
		textDate += " " + rt.Mention.Date.End
	}
	date, err := time.Parse(DateMentionTimeFormat, textDate)
	if err != nil {
		return nil, nil, nil, ""
	}
	rel := converter.Relation{
		BlockID: bson.NewObjectId().Hex(),
		Name:    "Date",
		Format:  model.RelationFormat_date,
	}
	id := bson.NewObjectId().Hex()
	details := map[string]*types.Value{rel.Name: pbtypes.Int64(date.Unix())}
	return &model.Block{
		Id: id,
		Content: &model.BlockContentOfRelation{
			Relation: &model.BlockContentRelation{
				Key: rel.BlockID,
			},
		},
	}, &rel, details, id
}

func (t *TextObject) handleLinkPreviewMention(rt api.RichText, text *strings.Builder) []*model.BlockContentTextMark {
	from := textUtil.UTF16RuneCountString(text.String())
	text.WriteString(rt.Mention.LinkPreview.Url)
	to := textUtil.UTF16RuneCountString(text.String())
	marks := rt.BuildMarkdownFromAnnotations(int32(from), int32(to))
	marks = append(marks, &model.BlockContentTextMark{
		Range: &model.Range{
			From: int32(from),
			To:   int32(to),
		},
		Type: model.BlockContentTextMark_Link,
	})
	return marks
}

func (t *TextObject) isNotTextBlocks() bool {
	return (len(t.RichText) == 1 && t.RichText[0].Type == api.Mention && t.RichText[0].Mention.Type == api.Date) ||
		len(t.RichText) == 1 && t.RichText[0].Type == api.Equation
}

type TextObjectWithChildren struct {
	TextObject
	Children []interface{} `json:"children"`
}

func (t *TextObjectWithChildren) MapChildren(req *MapRequest) *MapResponse {
	childReq := *req
	childReq.Blocks = t.Children
	resp := MapBlocks(&childReq)
	return resp
}
