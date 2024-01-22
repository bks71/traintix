// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: pb/ticketing.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Ticketing_Purchase_FullMethodName                 = "/pb.Ticketing/Purchase"
	Ticketing_GetReservation_FullMethodName           = "/pb.Ticketing/GetReservation"
	Ticketing_GetReservationsBySection_FullMethodName = "/pb.Ticketing/GetReservationsBySection"
	Ticketing_CancelReservation_FullMethodName        = "/pb.Ticketing/CancelReservation"
	Ticketing_ChangeSeat_FullMethodName               = "/pb.Ticketing/ChangeSeat"
)

// TicketingClient is the client API for Ticketing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketingClient interface {
	Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*TicketingResponse, error)
	GetReservation(ctx context.Context, in *GetReservationRequest, opts ...grpc.CallOption) (*TicketingResponse, error)
	GetReservationsBySection(ctx context.Context, in *GetReservationsBySectionRequest, opts ...grpc.CallOption) (*TicketingResponse, error)
	CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*TicketingResponse, error)
	ChangeSeat(ctx context.Context, in *ChangeSeatRequest, opts ...grpc.CallOption) (*TicketingResponse, error)
}

type ticketingClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketingClient(cc grpc.ClientConnInterface) TicketingClient {
	return &ticketingClient{cc}
}

func (c *ticketingClient) Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*TicketingResponse, error) {
	out := new(TicketingResponse)
	err := c.cc.Invoke(ctx, Ticketing_Purchase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingClient) GetReservation(ctx context.Context, in *GetReservationRequest, opts ...grpc.CallOption) (*TicketingResponse, error) {
	out := new(TicketingResponse)
	err := c.cc.Invoke(ctx, Ticketing_GetReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingClient) GetReservationsBySection(ctx context.Context, in *GetReservationsBySectionRequest, opts ...grpc.CallOption) (*TicketingResponse, error) {
	out := new(TicketingResponse)
	err := c.cc.Invoke(ctx, Ticketing_GetReservationsBySection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingClient) CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*TicketingResponse, error) {
	out := new(TicketingResponse)
	err := c.cc.Invoke(ctx, Ticketing_CancelReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingClient) ChangeSeat(ctx context.Context, in *ChangeSeatRequest, opts ...grpc.CallOption) (*TicketingResponse, error) {
	out := new(TicketingResponse)
	err := c.cc.Invoke(ctx, Ticketing_ChangeSeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketingServer is the server API for Ticketing service.
// All implementations must embed UnimplementedTicketingServer
// for forward compatibility
type TicketingServer interface {
	Purchase(context.Context, *PurchaseRequest) (*TicketingResponse, error)
	GetReservation(context.Context, *GetReservationRequest) (*TicketingResponse, error)
	GetReservationsBySection(context.Context, *GetReservationsBySectionRequest) (*TicketingResponse, error)
	CancelReservation(context.Context, *CancelReservationRequest) (*TicketingResponse, error)
	ChangeSeat(context.Context, *ChangeSeatRequest) (*TicketingResponse, error)
	mustEmbedUnimplementedTicketingServer()
}

// UnimplementedTicketingServer must be embedded to have forward compatible implementations.
type UnimplementedTicketingServer struct {
}

func (UnimplementedTicketingServer) Purchase(context.Context, *PurchaseRequest) (*TicketingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purchase not implemented")
}
func (UnimplementedTicketingServer) GetReservation(context.Context, *GetReservationRequest) (*TicketingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservation not implemented")
}
func (UnimplementedTicketingServer) GetReservationsBySection(context.Context, *GetReservationsBySectionRequest) (*TicketingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationsBySection not implemented")
}
func (UnimplementedTicketingServer) CancelReservation(context.Context, *CancelReservationRequest) (*TicketingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservation not implemented")
}
func (UnimplementedTicketingServer) ChangeSeat(context.Context, *ChangeSeatRequest) (*TicketingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSeat not implemented")
}
func (UnimplementedTicketingServer) mustEmbedUnimplementedTicketingServer() {}

// UnsafeTicketingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketingServer will
// result in compilation errors.
type UnsafeTicketingServer interface {
	mustEmbedUnimplementedTicketingServer()
}

func RegisterTicketingServer(s grpc.ServiceRegistrar, srv TicketingServer) {
	s.RegisterService(&Ticketing_ServiceDesc, srv)
}

func _Ticketing_Purchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServer).Purchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ticketing_Purchase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServer).Purchase(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticketing_GetReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServer).GetReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ticketing_GetReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServer).GetReservation(ctx, req.(*GetReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticketing_GetReservationsBySection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationsBySectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServer).GetReservationsBySection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ticketing_GetReservationsBySection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServer).GetReservationsBySection(ctx, req.(*GetReservationsBySectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticketing_CancelReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServer).CancelReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ticketing_CancelReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServer).CancelReservation(ctx, req.(*CancelReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticketing_ChangeSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeSeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServer).ChangeSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ticketing_ChangeSeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServer).ChangeSeat(ctx, req.(*ChangeSeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ticketing_ServiceDesc is the grpc.ServiceDesc for Ticketing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ticketing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Ticketing",
	HandlerType: (*TicketingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Purchase",
			Handler:    _Ticketing_Purchase_Handler,
		},
		{
			MethodName: "GetReservation",
			Handler:    _Ticketing_GetReservation_Handler,
		},
		{
			MethodName: "GetReservationsBySection",
			Handler:    _Ticketing_GetReservationsBySection_Handler,
		},
		{
			MethodName: "CancelReservation",
			Handler:    _Ticketing_CancelReservation_Handler,
		},
		{
			MethodName: "ChangeSeat",
			Handler:    _Ticketing_ChangeSeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/ticketing.proto",
}
