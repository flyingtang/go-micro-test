// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: prod_service.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ProdService service

type ProdService interface {
	GetProdsList(ctx context.Context, in *ProdListRequest, opts ...client.CallOption) (*ProdListResponse, error)
}

type prodService struct {
	c    client.Client
	name string
}

func NewProdService(name string, c client.Client) ProdService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "services"
	}
	return &prodService{
		c:    c,
		name: name,
	}
}

func (c *prodService) GetProdsList(ctx context.Context, in *ProdListRequest, opts ...client.CallOption) (*ProdListResponse, error) {
	req := c.c.NewRequest(c.name, "ProdService.GetProdsList", in)
	out := new(ProdListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProdService service

type ProdServiceHandler interface {
	GetProdsList	(context.Context, *ProdListRequest, *ProdListResponse) error
}

func RegisterProdServiceHandler(s server.Server, hdlr ProdServiceHandler, opts ...server.HandlerOption) error {
	type prodService interface {
		GetProdsList(ctx context.Context, in *ProdListRequest, out *ProdListResponse) error
	}
	type ProdService struct {
		prodService
	}
	h := &prodServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProdService{h}, opts...))
}

type prodServiceHandler struct {
	ProdServiceHandler
}

func (h *prodServiceHandler) GetProdsList(ctx context.Context, in *ProdListRequest, out *ProdListResponse) error {
	return h.ProdServiceHandler.GetProdsList(ctx, in, out)
}
