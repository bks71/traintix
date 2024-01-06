// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
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

// TicketingClient is the client API for Ticketing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketingClient interface {
	// Public API where you can submit a purchase for a ticket
	Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
	// Authenticated API that shows the details of the receipt for the user
	ViewReservation(ctx context.Context, in *ViewReservationRequest, opts ...grpc.CallOption) (*ViewReservationResponse, error)
	// authenticated API that lets an admin view all the users and seats they are allocated by the requested section
	ViewAllReservations(ctx context.Context, in *ViewAllReservationsRequest, opts ...grpc.CallOption) (*ViewAllReservationsResponse, error)
	// authenticated API to allow an admin or the user to remove the user from the train
	CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*CancelReservationResponse, error)
	ChangeSeat(ctx context.Context, in *ChangeSeatRequest, opts ...grpc.CallOption) (*ChangeSeatResponse, error)
}

type ticketingClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketingClient(cc grpc.ClientConnInterface) TicketingClient {
	return &ticketingClient{cc}
}

func (c *ticketingClient) Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error) {
	out := new(PurchaseResponse)
	err := c.cc.Invoke(ctx, "/pb.Ticketing/Purchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingClient) ViewReservation(ctx context.Context, in *ViewReservationRequest, opts ...grpc.CallOption) (*ViewReservationResponse, error) {
	out := new(ViewReservationResponse)
	err := c.cc.Invoke(ctx, "/pb.Ticketing/ViewReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingClient) ViewAllReservations(ctx context.Context, in *ViewAllReservationsRequest, opts ...grpc.CallOption) (*ViewAllReservationsResponse, error) {
	out := new(ViewAllReservationsResponse)
	err := c.cc.Invoke(ctx, "/pb.Ticketing/ViewAllReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingClient) CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*CancelReservationResponse, error) {
	out := new(CancelReservationResponse)
	err := c.cc.Invoke(ctx, "/pb.Ticketing/CancelReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingClient) ChangeSeat(ctx context.Context, in *ChangeSeatRequest, opts ...grpc.CallOption) (*ChangeSeatResponse, error) {
	out := new(ChangeSeatResponse)
	err := c.cc.Invoke(ctx, "/pb.Ticketing/ChangeSeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketingServer is the server API for Ticketing service.
// All implementations must embed UnimplementedTicketingServer
// for forward compatibility
type TicketingServer interface {
	// Public API where you can submit a purchase for a ticket
	Purchase(context.Context, *PurchaseRequest) (*PurchaseResponse, error)
	// Authenticated API that shows the details of the receipt for the user
	ViewReservation(context.Context, *ViewReservationRequest) (*ViewReservationResponse, error)
	// authenticated API that lets an admin view all the users and seats they are allocated by the requested section
	ViewAllReservations(context.Context, *ViewAllReservationsRequest) (*ViewAllReservationsResponse, error)
	// authenticated API to allow an admin or the user to remove the user from the train
	CancelReservation(context.Context, *CancelReservationRequest) (*CancelReservationResponse, error)
	ChangeSeat(context.Context, *ChangeSeatRequest) (*ChangeSeatResponse, error)
	mustEmbedUnimplementedTicketingServer()
}

// UnimplementedTicketingServer must be embedded to have forward compatible implementations.
type UnimplementedTicketingServer struct {
}

func (UnimplementedTicketingServer) Purchase(context.Context, *PurchaseRequest) (*PurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purchase not implemented")
}
func (UnimplementedTicketingServer) ViewReservation(context.Context, *ViewReservationRequest) (*ViewReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewReservation not implemented")
}
func (UnimplementedTicketingServer) ViewAllReservations(context.Context, *ViewAllReservationsRequest) (*ViewAllReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAllReservations not implemented")
}
func (UnimplementedTicketingServer) CancelReservation(context.Context, *CancelReservationRequest) (*CancelReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservation not implemented")
}
func (UnimplementedTicketingServer) ChangeSeat(context.Context, *ChangeSeatRequest) (*ChangeSeatResponse, error) {
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
		FullMethod: "/pb.Ticketing/Purchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServer).Purchase(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticketing_ViewReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServer).ViewReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Ticketing/ViewReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServer).ViewReservation(ctx, req.(*ViewReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticketing_ViewAllReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewAllReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServer).ViewAllReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Ticketing/ViewAllReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServer).ViewAllReservations(ctx, req.(*ViewAllReservationsRequest))
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
		FullMethod: "/pb.Ticketing/CancelReservation",
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
		FullMethod: "/pb.Ticketing/ChangeSeat",
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
			MethodName: "ViewReservation",
			Handler:    _Ticketing_ViewReservation_Handler,
		},
		{
			MethodName: "ViewAllReservations",
			Handler:    _Ticketing_ViewAllReservations_Handler,
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
