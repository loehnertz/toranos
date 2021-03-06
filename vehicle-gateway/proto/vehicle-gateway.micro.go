// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: vehicle-gateway.proto

/*
Package vehicle_gateway is a generated protocol buffer package.

It is generated from these files:
	vehicle-gateway.proto

It has these top-level messages:
	ReachVehicleRequest
	ReachVehicleResponse
*/
package vehicle_gateway

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VehicleGateway service

type VehicleGatewayService interface {
	ReachVehicle(ctx context.Context, in *ReachVehicleRequest, opts ...client.CallOption) (*ReachVehicleResponse, error)
}

type vehicleGatewayService struct {
	c    client.Client
	name string
}

func NewVehicleGatewayService(name string, c client.Client) VehicleGatewayService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "vehiclegateway"
	}
	return &vehicleGatewayService{
		c:    c,
		name: name,
	}
}

func (c *vehicleGatewayService) ReachVehicle(ctx context.Context, in *ReachVehicleRequest, opts ...client.CallOption) (*ReachVehicleResponse, error) {
	req := c.c.NewRequest(c.name, "VehicleGateway.ReachVehicle", in)
	out := new(ReachVehicleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VehicleGateway service

type VehicleGatewayHandler interface {
	ReachVehicle(context.Context, *ReachVehicleRequest, *ReachVehicleResponse) error
}

func RegisterVehicleGatewayHandler(s server.Server, hdlr VehicleGatewayHandler, opts ...server.HandlerOption) {
	type vehicleGateway interface {
		ReachVehicle(ctx context.Context, in *ReachVehicleRequest, out *ReachVehicleResponse) error
	}
	type VehicleGateway struct {
		vehicleGateway
	}
	h := &vehicleGatewayHandler{hdlr}
	s.Handle(s.NewHandler(&VehicleGateway{h}, opts...))
}

type vehicleGatewayHandler struct {
	VehicleGatewayHandler
}

func (h *vehicleGatewayHandler) ReachVehicle(ctx context.Context, in *ReachVehicleRequest, out *ReachVehicleResponse) error {
	return h.VehicleGatewayHandler.ReachVehicle(ctx, in, out)
}
