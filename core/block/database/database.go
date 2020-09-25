package database

import (
	"fmt"

	"github.com/anytypeio/go-anytype-middleware/core/anytype"
	"github.com/anytypeio/go-anytype-middleware/core/block/database/pages"
	"github.com/anytypeio/go-anytype-middleware/pb"
	coresb "github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/gogo/protobuf/types"
)

type Ctrl interface {
	Anytype() anytype.Service
	SetDetails(req pb.RpcBlockSetDetailsRequest) error
	CreateSmartBlock(sbType coresb.SmartBlockType, details *types.Struct) (pageId string, err error)
}

type Router interface {
	Get(id string) (database.Database, error)
}

func New(s Ctrl) Router {
	return &router{s: s}
}

type router struct{ s Ctrl }

func (r router) Get(id string) (database.Database, error) {
	switch id {
	case "pages":
		return pages.New(r.s.Anytype().PageStore(), r.s.SetDetails, r.s.CreateSmartBlock), nil
	}
	return nil, fmt.Errorf("db not found")
}
