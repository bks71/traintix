package api

import (
	"context"
	"log"

	"github.com/bks71/traintix/pb"
	"github.com/bks71/traintix/server/inventory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TicketingServer struct {
	pb.UnimplementedTicketingServer
	inventory.InventoryService
}

func NewTicketingServer() (*TicketingServer, error) {
	return &TicketingServer{
		InventoryService: inventory.NewInventory(),
	}, nil
}

func (ts *TicketingServer) RegisterService(g grpc.ServiceRegistrar) {
	pb.RegisterTicketingServer(g, ts)
}

func (s *TicketingServer) Purchase(ctx context.Context, in *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	log.Printf("%v", in)
	resv, err := s.InventoryService.Purchase(in.User.FirstName, in.User.LastName, in.User.Email)
	if err != nil {
		return nil, err
	}

	log.Printf("%v", resv)

	resp := &pb.PurchaseResponse{
		Receipt: &pb.Receipt{
			From:  resv.From,
			To:    resv.To,
			User:  &pb.User{FirstName: resv.Passenger.FirstName, LastName: resv.Passenger.LastName, Email: resv.Passenger.Email},
			Price: resv.Price,
			// Seat:  &pb.Seat{SeatNumber: int32(resv.SeatNumber)},
			Seat: &pb.Seat{Section: resv.Section, SeatNumber: int32(resv.SeatNumber)},
		},
	}

	log.Printf("%v", resp)

	return resp, nil

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
