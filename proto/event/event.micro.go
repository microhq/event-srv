// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/microhq/event-srv/proto/event/event.proto

/*
Package go_micro_srv_event is a generated protocol buffer package.

It is generated from these files:
	github.com/microhq/event-srv/proto/event/event.proto

It has these top-level messages:
	Record
	ReadRequest
	ReadResponse
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	SearchRequest
	SearchResponse
	StreamRequest
	StreamResponse
*/
package go_micro_srv_event

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	api "github.com/micro/go-api"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Event service

type EventService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (Event_StreamService, error)
	// should not really be used since event sourcing is append only
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type eventService struct {
	c    client.Client
	name string
}

func NewEventService(name string, c client.Client) EventService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.event"
	}
	return &eventService{
		c:    c,
		name: name,
	}
}

func (c *eventService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Event.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Event.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "Event.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (Event_StreamService, error) {
	req := c.c.NewRequest(c.name, "Event.Stream", &StreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &eventServiceStream{stream}, nil
}

type Event_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamResponse, error)
}

type eventServiceStream struct {
	stream client.Stream
}

func (x *eventServiceStream) Close() error {
	return x.stream.Close()
}

func (x *eventServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *eventServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *eventServiceStream) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Event.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Event.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Event service

type EventHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	Stream(context.Context, *StreamRequest, Event_StreamStream) error
	// should not really be used since event sourcing is append only
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterEventHandler(s server.Server, hdlr EventHandler, opts ...server.HandlerOption) error {
	type event interface {
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error
		Stream(ctx context.Context, stream server.Stream) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type Event struct {
		event
	}
	h := &eventHandler{hdlr}
	return s.Handle(s.NewHandler(&Event{h}, opts...))
	return s.Handle(s.NewHandler(&Event{h}, opts...))
	return s.Handle(s.NewHandler(&Event{h}, opts...))
	return s.Handle(s.NewHandler(&Event{h}, opts...))
	return s.Handle(s.NewHandler(&Event{h}, opts...))
	return s.Handle(s.NewHandler(&Event{h}, opts...))
}

type eventHandler struct {
	EventHandler
}

func (h *eventHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.EventHandler.Read(ctx, in, out)
}

func (h *eventHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.EventHandler.Create(ctx, in, out)
}

func (h *eventHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.EventHandler.Search(ctx, in, out)
}

func (h *eventHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.EventHandler.Stream(ctx, m, &eventStreamStream{stream})
}

type Event_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamResponse) error
}

type eventStreamStream struct {
	stream server.Stream
}

func (x *eventStreamStream) Close() error {
	return x.stream.Close()
}

func (x *eventStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *eventStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *eventStreamStream) Send(m *StreamResponse) error {
	return x.stream.Send(m)
}

func (h *eventHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.EventHandler.Update(ctx, in, out)
}

func (h *eventHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.EventHandler.Delete(ctx, in, out)
}
