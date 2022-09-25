package slice

import (
	"github.com/mb0/diff"
)

type DiffOperation int

const (
	OperationAdd     DiffOperation = iota
	OperationMove    DiffOperation = iota
	OperationRemove  DiffOperation = iota
	OperationReplace DiffOperation = iota
)

type Change struct {
	Op      DiffOperation
	Ids     []string
	AfterId string
}

type MixedInput struct {
	A []string
	B []string
}

func (m *MixedInput) Equal(a, b int) bool {
	return m.A[a] == m.B[b]
}

func Diff(origin, changed []string) []Change {
	m := &MixedInput{
		origin,
		changed,
	}

	var chs []Change

	changes := diff.Diff(len(m.A), len(m.B), m)
	var delIds []string
	for _, c := range changes {
		if c.Del > 0 {
			delIds = append(delIds, m.A[c.A:c.A+c.Del]...)
		}
	}

	if len(delIds) > 0 {
		chs = append(chs, Change{Op: OperationRemove, Ids: delIds})
	}

	for _, c := range changes {
		if c.Ins > 0 {
			inserts := m.B[c.B:c.B+c.Ins]
			afterId := ""
			if c.A >  0 {
				afterId = m.A[c.A-1]
			}
			chs = append(chs, Change{Op: OperationAdd, Ids: inserts, AfterId: afterId})
		}
	}

	return chs
}

func ApplyChanges(origin []string, changes []Change) []string {
	result := make([]string, len(origin))
	copy(result, origin)

	for _, ch := range changes {

		switch ch.Op {
		case OperationAdd:
			pos := -1
			if ch.AfterId != "" {
				pos = FindPos(result, ch.AfterId)
				if pos < 0 {
					continue
				}
			}

			ids := make([]string, len(ch.Ids))
			copy(ids, ch.Ids)
			result = Insert(result, pos+1, ids...)
		case OperationMove:
			// TODO
		case OperationRemove:
			result = Filter(result, func(id string) bool{
				return FindPos(ch.Ids, id) < 0
			})
		case OperationReplace:
			result = ch.Ids
		}
	}

	return result
}
