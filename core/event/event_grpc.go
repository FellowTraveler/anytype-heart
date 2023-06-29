//go:build !nogrpcserver && !_test
// +build !nogrpcserver,!_test

package event

import (
	"fmt"
	"sync"

	"github.com/anyproto/any-sync/app"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pb/service"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
)

var log = logging.Logger("anytype-grpc")

func NewGrpcSender() *GrpcSender {
	gs := &GrpcSender{
		shutdownCh: make(chan string),
	}

	go func() {
		for id := range gs.shutdownCh {
			gs.CloseSession(id)
		}
	}()

	return gs
}

type GrpcSender struct {
	ServerMutex sync.RWMutex
	Servers     map[string]SessionServer

	shutdownCh chan string
}

func (es *GrpcSender) Init(_ *app.App) (err error) {
	return
}

func (es *GrpcSender) Name() (name string) {
	return CName
}

func (es *GrpcSender) IsActive(token string) bool {
	es.ServerMutex.RLock()
	defer es.ServerMutex.RUnlock()

	_, ok := es.Servers[token]
	return ok
}

func (es *GrpcSender) SendToSession(token string, event *pb.Event) {
	es.ServerMutex.RLock()
	defer es.ServerMutex.RUnlock()

	if s, ok := es.Servers[token]; ok {
		es.sendEvent(s, event)
	}
}

func (es *GrpcSender) sendEvent(server SessionServer, event *pb.Event) {
	go func() {
		err := server.Server.Send(event)
		if err != nil {
			if s, ok := status.FromError(err); ok && s.Code() == codes.Unavailable {
				es.shutdownCh <- server.Token
			}
			log.With("session", server.Token, "spaceID", server.SpaceID).Errorf("failed to send event: %s", err)
		}
	}()
}

func (es *GrpcSender) Broadcast(event *pb.Event) {
	es.broadcast(nil, event)
}

func (es *GrpcSender) BroadcastForSpace(spaceID string, event *pb.Event) {
	// TODO Use spaceID
	es.broadcast(nil, event)
}

// BroadcastToOtherSessions broadcasts the event from current session. Do not broadcast to the current session
func (es *GrpcSender) BroadcastToOtherSessions(token string, event *pb.Event) {
	// TODO Use spaceID
	es.broadcast(&token, event)
}

// broadcast to all servers except server registered by ignoreSession token
func (es *GrpcSender) broadcast(ignoreSession *string, event *pb.Event) {
	es.ServerMutex.RLock()
	defer es.ServerMutex.RUnlock()

	for id, s := range es.Servers {
		if ignoreSession != nil && *ignoreSession == id {
			continue
		}
		go func(s SessionServer, id string) {
			err := s.Server.Send(event)
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Unavailable {
					es.shutdownCh <- id
				}
				log.Errorf("failed to send event: %s", err.Error())
			}
		}(s, id)
	}
}

type SessionServer struct {
	Token   string
	SpaceID string
	Done    chan struct{}
	Server  service.ClientCommands_ListenSessionEventsServer
}

func (es *GrpcSender) SetSessionServer(token string, server service.ClientCommands_ListenSessionEventsServer) SessionServer {
	log.Warnf("listening %s\n", token)
	es.ServerMutex.Lock()
	defer es.ServerMutex.Unlock()
	if es.Servers == nil {
		es.Servers = map[string]SessionServer{}
	}
	srv := SessionServer{
		Token:  token,
		Done:   make(chan struct{}),
		Server: server,
		// TODO SpaceID
	}

	// Old connection with this token will be cancelled automatically
	es.Servers[token] = srv
	return srv
}

func (es *GrpcSender) SetSpaceID(token string, spaceID string) error {
	es.ServerMutex.Lock()
	defer es.ServerMutex.Unlock()

	s, ok := es.Servers[token]
	if !ok {
		return fmt.Errorf("unknown session %s", token)
	}
	s.SpaceID = spaceID
	es.Servers[token] = s
	return nil
}

func (es *GrpcSender) CloseSession(token string) {
	es.ServerMutex.Lock()
	defer es.ServerMutex.Unlock()

	s, ok := es.Servers[token]
	if ok {
		log.Errorf("method close session %s", token)
		close(s.Done)
		delete(es.Servers, token)
	}
}
