package subscription

import (
	"github.com/mb0/diff"
)

func newListDiff(ids []string) *listDiff {
	ld := &listDiff{
		beforeIds:  ids,
		beforeIdsM: map[string]struct{}{},
		afterIdsM:  map[string]struct{}{},
	}
	for _, id := range ids {
		ld.beforeIdsM[id] = struct{}{}
	}
	return ld
}

type listDiff struct {
	beforeIds, afterIds   []string
	beforeIdsM, afterIdsM map[string]struct{}
}

func (ld *listDiff) Equal(i, j int) bool { return ld.beforeIds[i] == ld.afterIds[j] }

func (ld *listDiff) fillAfter(id string) {
	ld.afterIds = append(ld.afterIds, id)
}

func (ld *listDiff) reverse() {
	for i, j := 0, len(ld.afterIds)-1; i < j; i, j = i+1, j-1 {
		ld.afterIds[i], ld.afterIds[j] = ld.afterIds[j], ld.afterIds[i]
	}
}

func (ld *listDiff) reset() {
	ld.beforeIds, ld.afterIds = ld.afterIds, ld.beforeIds
	ld.afterIds = ld.afterIds[:0]
	ld.beforeIdsM = ld.afterIdsM
	ld.afterIdsM = make(map[string]struct{})
}

func (ld *listDiff) diff(ctx *opCtx, subId string, keys []string) (wasAddOrRemove bool, ids []string) {
	for _, id := range ld.afterIds {
		ld.afterIdsM[id] = struct{}{}
	}

	hasBefore := func(id string) bool {
		if _, ok := ld.beforeIdsM[id]; ok {
			return true
		}
		return false
	}
	hasAfter := func(id string) bool {
		if _, ok := ld.afterIdsM[id]; ok {
			return true
		}
		return false
	}
	getPrevId := func(s []string, i int) string {
		if i == 0 {
			return ""
		}
		return s[i-1]
	}
	diffData := diff.Diff(len(ld.beforeIds), len(ld.afterIds), ld)
	for _, ch := range diffData {
		for i := 0; i < ch.Ins; i++ {
			idx := ch.B + i
			isAdd := !hasBefore(ld.afterIds[idx])
			ctx.position = append(ctx.position, opPosition{
				id:      ld.afterIds[idx],
				subId:   subId,
				keys:    keys,
				afterId: getPrevId(ld.afterIds, idx),
				isAdd:   isAdd,
			})
			if isAdd {
				ids = append(ids, ld.afterIds[idx])
				wasAddOrRemove = true
			}
		}
		for i := 0; i < ch.Del; i++ {
			idx := ch.A + i
			if !hasAfter(ld.beforeIds[idx]) {
				ctx.remove = append(ctx.remove, opRemove{
					id:    ld.beforeIds[idx],
					subId: subId,
				})
				wasAddOrRemove = true
			}
		}
	}
	return
}
