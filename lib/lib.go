package lib

import (
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/anytypeio/go-anytype-middleware/core"
	"github.com/anytypeio/go-anytype-middleware/core/event"
	"github.com/anytypeio/go-anytype-middleware/pb"

	"github.com/anytypeio/go-anytype-library/logging"
	"github.com/gogo/protobuf/proto"
)

var log = logging.Logger("anytype-mw")

var mw = core.New()

func init() {
	registerClientCommandsHandler(mw)
	if debug, ok := os.LookupEnv("ANYPROF"); ok && debug != "" {
		go func() {
			http.ListenAndServe(debug, nil)
		}()
	}
}

func SetEventHandler(eh func(event *pb.Event)) {
	mw.EventSender = event.NewCallbackSender(eh)
}

func SetEventHandlerMobile(eh MessageHandler) {
	SetEventHandler(func(event *pb.Event) {
		b, err := proto.Marshal(event)
		if err != nil {
			log.Errorf("eventHandler failed to marshal error: %s", err.Error())
		}
		eh.Handle(b)
	})
}
