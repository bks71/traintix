package api_server

import (
	"context"
	"log"

	"github.com/bks71/traintix/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TicketingServer struct {
	pb.UnimplementedTicketingServer
}

func (s *TicketingServer) Purchase(ctx context.Context, in *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	log.Printf("Received: %v", in)
	return nil, status.Errorf(codes.PermissionDenied, "denied")
}

func (s *TicketingServer) ViewReservation(ctx context.Context, in *pb.ViewReservationRequest) (*pb.ViewReservationResponse, error) {
	log.Printf("Received: %v", in)
	return nil, status.Errorf(codes.PermissionDenied, "denied")
}

func (s *TicketingServer) ViewAllReservations(ctx context.Context, in *pb.ViewAllReservationsRequest) (*pb.ViewAllReservationsResponse, error) {
	log.Printf("Received: %v", in)
	return nil, status.Errorf(codes.PermissionDenied, "denied")
}

func (s *TicketingServer) ChangeSeat(ctx context.Context, in *pb.ChangeSeatRequest) (*pb.ChangeSeatResponse, error) {
	log.Printf("Received: %v", in)
	return nil, status.Errorf(codes.PermissionDenied, "denied")
}

func (s *TicketingServer) CancelReservation(ctx context.Context, in *pb.CancelReservationRequest) (*pb.CancelReservationResponse, error) {
	log.Printf("Received: %v", in)
	return nil, status.Errorf(codes.PermissionDenied, "denied")
}
