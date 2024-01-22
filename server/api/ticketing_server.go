package api

import (
	"context"
	"log"

	"github.com/bks71/traintix/pb"
	"github.com/bks71/traintix/server/reservations"
	"google.golang.org/grpc"
)

type TicketingServer struct {
	pb.UnimplementedTicketingServer
	reservations.ReservationSystem
}

func NewTicketingServer() (*TicketingServer, error) {
	return &TicketingServer{
		ReservationSystem: reservations.NewReservationSystem(),
	}, nil
}

func (ts *TicketingServer) RegisterService(g grpc.ServiceRegistrar) {
	pb.RegisterTicketingServer(g, ts)
}

func (s *TicketingServer) Purchase(ctx context.Context, in *pb.PurchaseRequest) (*pb.TicketingResponse, error) {
	log.Printf("%v", in)

	r, err := s.ReservationSystem.ReserveSeat(in.Passenger)
	if err != nil {
		return nil, err
	}

	resp := &pb.TicketingResponse{
		Reservation: []*pb.Reservation{r},
	}

	log.Printf("%v", resp)

	return resp, nil

}

func (s *TicketingServer) GetReservation(ctx context.Context, in *pb.GetReservationRequest) (*pb.TicketingResponse, error) {
	log.Printf("%v", in)

	r, err := s.ReservationSystem.GetReservation(in.Passenger)
	if err != nil {
		return nil, err
	}

	resp := &pb.TicketingResponse{
		Reservation: []*pb.Reservation{r},
	}

	log.Printf("%v", resp)

	return resp, nil
}

func (s *TicketingServer) GetReservationsBySection(ctx context.Context, in *pb.GetReservationsBySectionRequest) (*pb.TicketingResponse, error) {
	log.Printf("%v", in)

	r := s.ReservationSystem.GetReservationsBySection(in.Section)

	resp := &pb.TicketingResponse{
		Reservation: r,
	}

	log.Printf("%v", resp)

	return resp, nil
}

func (s *TicketingServer) ChangeSeat(ctx context.Context, in *pb.ChangeSeatRequest) (*pb.TicketingResponse, error) {
	log.Printf("%v", in)

	r, err := s.ReservationSystem.ChangeSeat(in.Passenger, in.Seat)
	if err != nil {
		return nil, err
	}

	resp := &pb.TicketingResponse{
		Reservation: []*pb.Reservation{r},
	}

	log.Printf("%v", resp)

	return resp, nil
}

func (s *TicketingServer) CancelReservation(ctx context.Context, in *pb.CancelReservationRequest) (*pb.TicketingResponse, error) {
	log.Printf("%v", in)

	r, err := s.ReservationSystem.CancelReservation(in.Passenger)
	if err != nil {
		return nil, err
	}

	resp := &pb.TicketingResponse{
		Reservation: []*pb.Reservation{r},
	}

	log.Printf("%v", resp)

	return resp, nil
}
