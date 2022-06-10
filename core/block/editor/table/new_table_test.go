package table

import (
	"testing"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRowCreate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		source   *state.State
		newRowId string
		req      pb.RpcBlockTableRowCreateRequest
		want     *state.State
	}{
		{
			name:     "cells are not affected",
			source:   mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col2"}}),
			newRowId: "row3",
			req: pb.RpcBlockTableRowCreateRequest{
				TargetId: "row1",
				Position: model.Block_Bottom,
			},
			want: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row3", "row2"}, [][]string{{"row1-col2"}}),
		},
		{
			name:     "between, bottom position",
			source:   mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, nil),
			newRowId: "row3",
			req: pb.RpcBlockTableRowCreateRequest{
				TargetId: "row1",
				Position: model.Block_Bottom,
			},
			want: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row3", "row2"}, nil),
		},
		{
			name:     "between, top position",
			source:   mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, nil),
			newRowId: "row3",
			req: pb.RpcBlockTableRowCreateRequest{
				TargetId: "row2",
				Position: model.Block_Top,
			},
			want: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row3", "row2"}, nil),
		},
		// TODO: more tests
	} {
		t.Run(tc.name, func(t *testing.T) {
			tb := table{
				generateRowId: idFromSlice([]string{tc.newRowId}),
			}

			err := tb.RowCreate(tc.source, tc.req)

			require.NoError(t, err)

			assert.Equal(t, tc.want, tc.source)
		})
	}
}

func TestExpand(t *testing.T) {
	for _, tc := range []struct {
		name      string
		source    *state.State
		newColIds []string
		newRowIds []string
		req       pb.RpcBlockTableExpandRequest
		want      *state.State
	}{
		{
			name:      "only rows",
			source:    mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row2-col2"}}),
			newRowIds: []string{"row3", "row4"},
			req: pb.RpcBlockTableExpandRequest{
				TargetId: "table",
				Rows:     2,
			},
			want: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2", "row3", "row4"}, [][]string{{"row2-col2"}}),
		},
		{
			name:      "only columns",
			source:    mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row2-col2"}}),
			newColIds: []string{"col3", "col4"},
			req: pb.RpcBlockTableExpandRequest{
				TargetId: "table",
				Columns:  2,
			},
			want: mkTestTable([]string{"col1", "col2", "col3", "col4"}, []string{"row1", "row2"}, [][]string{{"row2-col2"}}),
		},
		{
			name:      "both columns and rows",
			source:    mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row2-col2"}}),
			newRowIds: []string{"row3", "row4"},
			newColIds: []string{"col3", "col4"},
			req: pb.RpcBlockTableExpandRequest{
				TargetId: "table",
				Rows:     2,
				Columns:  2,
			},
			want: mkTestTable([]string{"col1", "col2", "col3", "col4"}, []string{"row1", "row2", "row3", "row4"}, [][]string{{"row2-col2"}}),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tb := table{
				generateColId: idFromSlice(tc.newColIds),
				generateRowId: idFromSlice(tc.newRowIds),
			}
			err := tb.Expand(tc.source, tc.req)
			require.NoError(t, err)
			assert.Equal(t, tc.want, tc.source)
		})
	}
}

func TestRowListFill(t *testing.T) {
	for _, tc := range []struct {
		name   string
		source *state.State
		req    pb.RpcBlockTableRowListFillRequest
		want   *state.State
	}{
		{
			name:   "empty",
			source: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{}),
			req: pb.RpcBlockTableRowListFillRequest{
				BlockIds: []string{"row1", "row2"},
			},
			want: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col2"},
					{"row2-col1", "row2-col2"},
				}),
		},
		{
			name: "fully filled",
			source: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col2"},
					{"row2-col1", "row2-col2"},
				}),
			req: pb.RpcBlockTableRowListFillRequest{
				BlockIds: []string{"row1", "row2"},
			},
			want: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col2"},
					{"row2-col1", "row2-col2"},
				}),
		},
		{
			name: "partially filled",
			source: mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2", "row3", "row4", "row5"},
				[][]string{
					{"row1-col1"},
					{"row2-col2"},
					{"row3-col3"},
					{"row4-col1", "row4-col3"},
					{"row5-col2", "row4-col3"},
				}),
			req: pb.RpcBlockTableRowListFillRequest{
				BlockIds: []string{"row1", "row2", "row3", "row4", "row5"},
			},
			want: mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2", "row3", "row4", "row5"},
				[][]string{
					{"row1-col1", "row1-col2", "row1-col3"},
					{"row2-col1", "row2-col2", "row2-col3"},
					{"row3-col1", "row3-col2", "row3-col3"},
					{"row4-col1", "row4-col2", "row4-col3"},
					{"row5-col1", "row5-col2", "row5-col3"},
				}),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tb := table{}
			err := tb.RowListFill(tc.source, tc.req)
			require.NoError(t, err)
			assert.Equal(t, tc.want, tc.source)
		})
	}
}

func TestColumnCreate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		source   *state.State
		newColId string
		req      pb.RpcBlockTableColumnCreateRequest
		want     *state.State
	}{
		{
			name:     "to the right",
			source:   mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col2"}}),
			newColId: "col3",
			req: pb.RpcBlockTableColumnCreateRequest{
				TargetId: "col1",
				Position: model.Block_Right,
			},
			want: mkTestTable([]string{"col1", "col3", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col2"}}),
		},
		{
			name:     "to the left",
			source:   mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col2"}}),
			newColId: "col3",
			req: pb.RpcBlockTableColumnCreateRequest{
				TargetId: "col1",
				Position: model.Block_Left,
			},
			want: mkTestTable([]string{"col3", "col1", "col2"}, []string{"row1", "row2"}, [][]string{{"row1-col2"}}),
		},
		// TODO: more tests
	} {
		t.Run(tc.name, func(t *testing.T) {
			tb := table{
				generateColId: idFromSlice([]string{tc.newColId}),
			}

			err := tb.ColumnCreate(tc.source, tc.req)

			require.NoError(t, err)

			assert.Equal(t, tc.want, tc.source)
		})
	}
}

func TestColumnDuplicate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		source   *state.State
		newColId string
		req      pb.RpcBlockTableColumnDuplicateRequest
		want     *state.State
	}{
		{
			name: "fully filled",
			source: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col2"},
					{"row2-col1", "row2-col2"},
				}),
			newColId: "col3",
			req: pb.RpcBlockTableColumnDuplicateRequest{
				BlockId:  "col1",
				TargetId: "col2",
				Position: model.Block_Right,
			},
			want: mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col2", "row1-col3"},
					{"row2-col1", "row2-col2", "row2-col3"},
				}),
		},
		{
			name: "partially filled",
			source: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2", "row3"},
				[][]string{
					{"row1-col1"},
					{"row2-col2"},
					{},
				}),
			newColId: "col3",
			req: pb.RpcBlockTableColumnDuplicateRequest{
				BlockId:  "col2",
				TargetId: "col1",
				Position: model.Block_Left,
			},
			want: mkTestTable([]string{"col3", "col1", "col2"}, []string{"row1", "row2", "row3"},
				[][]string{
					{"row1-col1"},
					{"row2-col3", "row2-col2"},
					{},
				}),
		},
		// TODO: more tests
	} {
		t.Run(tc.name, func(t *testing.T) {
			tb := table{
				generateColId: idFromSlice([]string{tc.newColId}),
			}
			id, err := tb.ColumnDuplicate(tc.source, tc.req)
			require.NoError(t, err)
			assert.Equal(t, tc.want, tc.source)
			assert.Equal(t, tc.newColId, id)
		})
	}
}

func TestColumnMove(t *testing.T) {
	for _, tc := range []struct {
		name   string
		source *state.State
		req    pb.RpcBlockTableColumnMoveRequest
		want   *state.State
	}{
		{
			name: "partial table",
			source: mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col2", "row1-col3"},
					{"row2-col1", "row2-col3"},
				}),
			req: pb.RpcBlockTableColumnMoveRequest{
				TargetId:     "col1",
				DropTargetId: "col3",
				Position:     model.Block_Left,
			},
			want: mkTestTable([]string{"col2", "col1", "col3"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col2", "row1-col1", "row1-col3"},
					{"row2-col1", "row2-col3"},
				}),
		},
		{
			name: "filled table",
			source: mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col2", "row1-col3"},
					{"row2-col1", "row2-col2", "row2-col3"},
				}),
			req: pb.RpcBlockTableColumnMoveRequest{
				TargetId:     "col3",
				DropTargetId: "col1",
				Position:     model.Block_Right,
			},
			want: mkTestTable([]string{"col1", "col3", "col2"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col3", "row1-col2"},
					{"row2-col1", "row2-col3", "row2-col2"},
				}),
		},
		// TODO: more tests
	} {
		t.Run(tc.name, func(t *testing.T) {
			tb := table{}

			err := tb.ColumnMove(tc.source, tc.req)

			require.NoError(t, err)

			assert.Equal(t, tc.want, tc.source)
		})
	}
}

func TestColumnDelete(t *testing.T) {
	for _, tc := range []struct {
		name   string
		source *state.State
		req    pb.RpcBlockTableColumnDeleteRequest
		want   *state.State
	}{
		{
			name: "partial table",
			source: mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col2", "row1-col3"},
					{"row2-col1", "row2-col3"},
				}),
			req: pb.RpcBlockTableColumnDeleteRequest{
				TargetId: "col2",
			},
			want: mkTestTable([]string{"col1", "col3"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col3"},
					{"row2-col1", "row2-col3"},
				}),
		},
		{
			name: "filled table",
			source: mkTestTable([]string{"col1", "col2", "col3"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col2", "row1-col3"},
					{"row2-col1", "row2-col2", "row2-col3"},
				}),
			req: pb.RpcBlockTableColumnDeleteRequest{
				TargetId: "col3",
			},
			want: mkTestTable([]string{"col1", "col2"}, []string{"row1", "row2"},
				[][]string{
					{"row1-col1", "row1-col2"},
					{"row2-col1", "row2-col2"},
				}),
		},
		// TODO: more tests
	} {
		t.Run(tc.name, func(t *testing.T) {
			tb := table{}

			err := tb.ColumnDelete(tc.source, tc.req)

			require.NoError(t, err)

			assert.Equal(t, tc.want.Blocks(), tc.source.Blocks())
		})
	}
}

func mkTestTable(columns []string, rows []string, cells [][]string) *state.State {
	s := state.NewDoc("root", nil).NewState()
	blocks := []*model.Block{
		{
			Id:          "root",
			ChildrenIds: []string{"table"},
		},
		{
			Id:          "table",
			ChildrenIds: []string{"columns", "rows"},
			Content:     &model.BlockContentOfTable{Table: &model.BlockContentTable{}},
		},
		{
			Id:          "columns",
			ChildrenIds: columns,
			Content: &model.BlockContentOfLayout{
				Layout: &model.BlockContentLayout{
					Style: model.BlockContentLayout_TableColumns,
				},
			},
		},
		{
			Id:          "rows",
			ChildrenIds: rows,
			Content: &model.BlockContentOfLayout{
				Layout: &model.BlockContentLayout{
					Style: model.BlockContentLayout_TableRows,
				},
			},
		},
	}

	for _, c := range columns {
		blocks = append(blocks, &model.Block{
			Id:      c,
			Content: &model.BlockContentOfTableColumn{TableColumn: &model.BlockContentTableColumn{}},
		})
	}

	cellsByRow := map[string][]string{}
	for _, cc := range cells {
		if len(cc) == 0 {
			continue
		}
		rowId, _, err := parseCellId(cc[0])
		if err != nil {
			panic(err)
		}
		cellsByRow[rowId] = cc

		for _, c := range cc {
			blocks = append(blocks, &model.Block{
				Id:      c,
				Content: &model.BlockContentOfText{Text: &model.BlockContentText{}},
			})
		}
	}

	for _, r := range rows {
		blocks = append(blocks, &model.Block{
			Id:          r,
			ChildrenIds: cellsByRow[r],
			Content:     &model.BlockContentOfTableRow{TableRow: &model.BlockContentTableRow{}},
		})
	}

	for _, b := range blocks {
		s.Add(simple.New(b))
	}
	return s
}

func idFromSlice(ids []string) func() string {
	var i int
	return func() string {
		id := ids[i]
		i++
		return id
	}
}
