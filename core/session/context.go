package session

import (
	"context"

	"github.com/anyproto/anytype-heart/core/event"
	"github.com/anyproto/anytype-heart/pb"
)

type Context interface {
	ID() string
	Context() context.Context
	ObjectID() string
	SpaceID() string
	TraceID() string
	SetIsAsync(bool)
	IsActive() bool
	Broadcast(event *pb.Event)
	Send(event *pb.Event)
	SendToOtherSessions(msgs []*pb.EventMessage)
	SetMessages(smartBlockId string, msgs []*pb.EventMessage)
	GetMessages() []*pb.EventMessage
	GetResponseEvent() *pb.ResponseEvent
}

type sessionContext struct {
	ctx           context.Context
	smartBlockId  string
	spaceID       string
	traceId       string
	messages      []*pb.EventMessage
	sessionSender event.Sender
	sessionToken  string
	isAsync       bool
}

func NewContext(cctx context.Context, eventSender event.Sender, spaceID string, opts ...ContextOption) Context {
	ctx := &sessionContext{
		spaceID:       spaceID,
		sessionSender: eventSender,
		ctx:           cctx,
	}
	for _, apply := range opts {
		apply(ctx)
	}
	return ctx
}

func NewChildContext(parent Context) Context {
	child := &sessionContext{
		ctx:          parent.Context(),
		spaceID:      parent.SpaceID(),
		smartBlockId: parent.ObjectID(),
		traceId:      parent.TraceID(),
		sessionToken: parent.ID(),
	}
	v, ok := parent.(*sessionContext)
	if ok {
		child.sessionSender = v.sessionSender
	}
	return child
}

func NewAsyncChildContext(parent *sessionContext) Context {
	ctx := NewChildContext(parent)
	ctx.SetIsAsync(true)
	return ctx
}

type ContextOption func(ctx *sessionContext)

func Async() ContextOption {
	return func(ctx *sessionContext) {
		ctx.isAsync = true
	}
}

func WithSession(token string) ContextOption {
	return func(ctx *sessionContext) {
		ctx.sessionToken = token
	}
}

func WithTraceId(traceId string) ContextOption {
	return func(ctx *sessionContext) {
		ctx.traceId = traceId
	}
}

type Closer interface {
	CloseSession(token string)
}

func (ctx *sessionContext) ID() string {
	return ctx.sessionToken
}

func (ctx *sessionContext) ObjectID() string {
	return ctx.smartBlockId
}

func (ctx *sessionContext) TraceID() string {
	return ctx.traceId
}

func (ctx *sessionContext) SpaceID() string {
	return ctx.spaceID
}

func (ctx *sessionContext) SetIsAsync(isAsync bool) {
	ctx.isAsync = isAsync
}

func (ctx *sessionContext) IsAsync() bool {
	return ctx.isAsync
}

func (ctx *sessionContext) Context() context.Context {
	return ctx.ctx
}

func (ctx *sessionContext) IsActive() bool {
	return ctx.sessionSender.IsActive(ctx.sessionToken)
}

func (ctx *sessionContext) AddMessages(smartBlockId string, msgs []*pb.EventMessage) {
	ctx.smartBlockId = smartBlockId
	ctx.messages = append(ctx.messages, msgs...)
}

func (ctx *sessionContext) SetMessages(smartBlockId string, msgs []*pb.EventMessage) {
	ctx.smartBlockId = smartBlockId
	ctx.messages = msgs
}

func (ctx *sessionContext) GetMessages() []*pb.EventMessage {
	return ctx.messages
}

func (ctx *sessionContext) Send(event *pb.Event) {
	ctx.sessionSender.SendToSession(ctx.sessionToken, event)
}

func (ctx *sessionContext) Broadcast(event *pb.Event) {
	ctx.sessionSender.BroadcastForSpace(ctx.spaceID, event)
}

func (ctx *sessionContext) SendToOtherSessions(msgs []*pb.EventMessage) {
	ctx.sessionSender.BroadcastToOtherSessions(ctx.sessionToken, &pb.Event{
		Messages:  msgs,
		ContextId: ctx.smartBlockId,
		Initiator: nil,
	})
}

func (ctx *sessionContext) GetResponseEvent() *pb.ResponseEvent {
	ctx.SendToOtherSessions(ctx.messages)
	return &pb.ResponseEvent{
		Messages:  ctx.messages,
		ContextId: ctx.smartBlockId,
		TraceId:   ctx.traceId,
	}
}
