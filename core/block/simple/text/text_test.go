package text

import (
	"testing"

	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/base"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestText_Diff(t *testing.T) {
	testBlock := func() *Text {
		return NewText(&model.Block{
			Restrictions: &model.BlockRestrictions{},
			Content:      &model.BlockContentOfText{Text: &model.BlockContentText{}},
		}).(*Text)
	}
	t.Run("type error", func(t *testing.T) {
		b1 := testBlock()
		b2 := base.NewBase(&model.Block{})
		_, err := b1.Diff(b2)
		assert.Error(t, err)
	})
	t.Run("no diff", func(t *testing.T) {
		b1 := testBlock()
		b2 := testBlock()
		b1.SetText("same text", &model.BlockContentTextMarks{})
		b2.SetText("same text", &model.BlockContentTextMarks{})
		d, err := b1.Diff(b2)
		require.NoError(t, err)
		assert.Len(t, d, 0)
	})
	t.Run("base diff", func(t *testing.T) {
		b1 := testBlock()
		b2 := testBlock()
		b2.Restrictions.Read = true
		d, err := b1.Diff(b2)
		require.NoError(t, err)
		assert.Len(t, d, 1)
	})
	t.Run("content diff", func(t *testing.T) {
		b1 := testBlock()
		b2 := testBlock()
		b2.SetText("text", &model.BlockContentTextMarks{
			Marks: []*model.BlockContentTextMark{
				{
					Range: &model.Range{1, 2},
					Type:  model.BlockContentTextMark_Italic,
				},
			},
		})
		b2.SetStyle(model.BlockContentText_Header2)
		b2.SetChecked(true)
		diff, err := b1.Diff(b2)
		require.NoError(t, err)
		require.Len(t, diff, 1)
		textChange := diff[0].Value.(*pb.EventMessageValueOfBlockSetText).BlockSetText
		assert.NotNil(t, textChange.Style)
		assert.NotNil(t, textChange.Checked)
		assert.NotNil(t, textChange.Text)
		assert.NotNil(t, textChange.Marks)
	})
}

func TestText_Split(t *testing.T) {
	testBlock := func() *Text {
		return NewText(&model.Block{
			Restrictions: &model.BlockRestrictions{},
			Content: &model.BlockContentOfText{Text: &model.BlockContentText{
				Text: "1234567890",
				Marks: &model.BlockContentTextMarks{
					Marks: []*model.BlockContentTextMark{
						{
							Type: model.BlockContentTextMark_Bold,
							Range: &model.Range{
								From: 0,
								To:   10,
							},
						},
						{
							Type: model.BlockContentTextMark_Italic,
							Range: &model.Range{
								From: 6,
								To:   10,
							},
						},
						{
							Type: model.BlockContentTextMark_BackgroundColor,
							Range: &model.Range{
								From: 3,
								To:   4,
							},
						},
					},
				},
			}},
		}).(*Text)
	}
	t.Run("should split block", func(t *testing.T) {
		b := testBlock()
		newBlock, err := b.Split(5)
		require.NoError(t, err)
		nb := newBlock.(*Text)
		assert.Equal(t, "12345", b.content.Text)
		assert.Equal(t, "67890", nb.content.Text)
		require.Len(t, b.content.Marks.Marks, 2)
		require.Len(t, nb.content.Marks.Marks, 2)
		assert.Equal(t, model.Range{0, 5}, *b.content.Marks.Marks[0].Range)
		assert.Equal(t, model.Range{3, 4}, *b.content.Marks.Marks[1].Range)
		assert.Equal(t, model.Range{0, 5}, *nb.content.Marks.Marks[0].Range)
		assert.Equal(t, model.Range{1, 5}, *nb.content.Marks.Marks[1].Range)
	})
	t.Run("out of range", func(t *testing.T) {
		b := testBlock()
		_, err := b.Split(15)
		require.Equal(t, ErrOutOfRange, err)
	})
}
