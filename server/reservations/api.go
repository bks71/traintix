package reservations

import "github.com/bks71/traintix/pb"

type ReservationSystem interface {
	ReserveSeat(*pb.Passenger) (*pb.Reservation, error)
	GetReservation(*pb.Passenger) (*pb.Reservation, error)
	CancelReservation(*pb.Passenger) (*pb.Reservation, error)
	ChangeSeat(*pb.Passenger, *pb.Seat) (*pb.Reservation, error)
	GetReservationsBySection(sectionName string) []*pb.Reservation
}
