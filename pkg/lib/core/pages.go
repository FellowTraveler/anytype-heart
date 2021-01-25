package core

import (
	"sort"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
)

func (a *Anytype) ObjectStore() localstore.ObjectStore {
	return a.localStore.Objects
}

// deprecated, to be removed
func (a *Anytype) ObjectInfoWithLinks(id string) (*model.ObjectInfoWithLinks, error) {
	return a.localStore.Objects.GetWithLinksInfoByID(id)
}

// deprecated, to be removed
func (a *Anytype) ObjectList() ([]*model.ObjectInfo, error) {
	ids, err := a.t.Logstore().Threads()
	if err != nil {
		return nil, err
	}

	var idsS = make([]string, 0, len(ids))
	for _, id := range ids {
		t, _ := smartblock.SmartBlockTypeFromThreadID(id)
		if t != smartblock.SmartBlockTypePage &&
			t != smartblock.SmartBlockTypeProfilePage &&
			t != smartblock.SmartBlockTypeHome {
			continue
		}

		idsS = append(idsS, id.String())
	}

	pages, err := a.localStore.Objects.GetByIDs(idsS...)
	if err != nil {
		return nil, err
	}

	sort.Slice(pages, func(i, j int) bool {
		// show pages with inbound links first
		if pages[i].HasInboundLinks && !pages[j].HasInboundLinks {
			return true
		}

		if !pages[i].HasInboundLinks && pages[j].HasInboundLinks {
			return false
		}

		// then sort by Last Opened date
		var lastOpenedI, lastOpenedJ int64

		if pages[i].Details != nil {
			if pages[i].Details.Fields[bundle.RelationKeyLastOpenedDate.String()] != nil {
				lastOpenedI = int64(pages[i].Details.Fields[bundle.RelationKeyLastOpenedDate.String()].GetNumberValue())
			}
		}

		if pages[j].Details != nil {
			if pages[j].Details.Fields[bundle.RelationKeyLastOpenedDate.String()] != nil {
				lastOpenedJ = int64(pages[j].Details.Fields[bundle.RelationKeyLastOpenedDate.String()].GetNumberValue())
			}
		}

		if lastOpenedI > lastOpenedJ {
			return true
		}

		if lastOpenedI < lastOpenedJ {
			return false
		}

		return pages[i].Id < pages[j].Id
	})

	return pages, nil
}
