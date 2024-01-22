package reservations

import "github.com/bks71/traintix/pb"

type ReservationSystem interface {
	ReserveSeat(*pb.Passenger) (*pb.Reservation, error)
	GetReservation(email string) (*pb.Reservation, error)
	CancelReservation(email string) (*pb.Reservation, error)
	ChangeSeat(email string, seat *pb.Seat) (*pb.Reservation, error)
	GetReservationsBySection(sectionName string) []*pb.Reservation
}
