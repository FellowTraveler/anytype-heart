package core

import (
	"context"

	libCore "github.com/anytypeio/go-anytype-library/core"
	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/pb"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("anytype-mw")

type MiddlewareState struct {
	// client-state: blocks range, text range, focus, screen position, etc
	// history list
	// request list
	// computed state
}

type Middleware struct {
	state               MiddlewareState
	rootPath            string
	pin                 string
	mnemonic            string
	accountSearchCancel context.CancelFunc
	localAccounts       []*pb.ModelAccount
	SendEvent           func(event *pb.Event)
	blockService        block.Service
	*libCore.Anytype
}

func (mw *Middleware) Stop() error {
	if mw != nil && mw.Anytype != nil {
		err := mw.Anytype.Stop()
		if err != nil {
			return err
		}

		mw.Anytype = nil
		mw.accountSearchCancel = nil
	}

	return nil
}
