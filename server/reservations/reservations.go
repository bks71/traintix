package reservations

import (
	"errors"

	"github.com/bks71/traintix/pb"
)

type ResvSys struct {
	// Available seats
	openSeats []*pb.Seat

	// Map of emails to reservations
	reservations map[string]*pb.Reservation
}

func NewReservationSystem() *ResvSys {
	return &ResvSys{
		openSeats: []*pb.Seat{
			{Section: "A", Number: 1},
			{Section: "A", Number: 2},
			{Section: "A", Number: 3},
			{Section: "A", Number: 4},
			{Section: "A", Number: 5},
			{Section: "A", Number: 6},
			{Section: "A", Number: 7},
			{Section: "A", Number: 8},
			{Section: "A", Number: 9},
			{Section: "A", Number: 10},
			{Section: "B", Number: 1},
			{Section: "B", Number: 2},
			{Section: "B", Number: 3},
			{Section: "B", Number: 4},
			{Section: "B", Number: 5},
			{Section: "B", Number: 6},
			{Section: "B", Number: 7},
			{Section: "B", Number: 8},
			{Section: "B", Number: 9},
			{Section: "B", Number: 10},
		},
		reservations: make(map[string]*pb.Reservation),
	}
}

func (resv *ResvSys) reserveAnySeat() (*pb.Seat, error) {
	if len(resv.openSeats) == 0 {
		return nil, errors.New("no seats available")
	}
	i := resv.openSeats
	s := i[len(i)-1]
	resv.openSeats = i[:len(i)-1]
	return s, nil
}

func (resv *ResvSys) ReserveSeat(p *pb.Passenger) (*pb.Reservation, error) {

	//check for an existing reservation for this passenger
	if r, exists := resv.reservations[p.Email]; exists {
		return r, errors.New("sorry this email address already has a reservation")
	}

	//Find an empty seat
	s, err := resv.reserveAnySeat()
	if err != nil {
		return nil, err
	}

	//Reserve it
	r := pb.Reservation{
		Passenger: p,
		From:      "London",
		To:        "France",
		Price:     20.00,
		Seat:      s,
	}
	resv.reservations[p.Email] = &r

	return &r, nil
}

func (resv *ResvSys) GetReservation(email string) (*pb.Reservation, error) {
	if r, exists := resv.reservations[email]; exists {
		return r, nil
	}
	return nil, errors.New("no reservation found")
}

func (resv *ResvSys) CancelReservation(email string) (*pb.Reservation, error) {
	if r, exists := resv.reservations[email]; exists {
		//TODO test this

		// put the seat back into the inventory
		resv.openSeats = append(resv.openSeats, r.Seat)

		// remove the passenger/reservation from the map
		delete(resv.reservations, email)

		// remove the passenger from the reservation and return it
		r.Passenger = nil
		return r, nil
	}

	return nil, errors.New("no reservation found")
}

func (resv *ResvSys) ChangeSeat(email string, s *pb.Seat) (*pb.Reservation, error) {

	r, exists := resv.reservations[email]
	if !exists {
		return nil, errors.New("no reservation found for this email address")
	}

	if r.Seat == s {
		return r, errors.New("passenger has already reserved this seat")
	}

	for i, v := range resv.openSeats {
		if v.Number == s.Number && v.Section == s.Section {
			// put old seat back into the inventory and switch passenger to new seat
			resv.openSeats[i] = r.Seat
			r.Seat = s
			return r, nil
		}
	}

	return r, errors.New("seat not available")
}

func (resv *ResvSys) GetReservationsBySection(sectionName string) []*pb.Reservation {
	var res []*pb.Reservation
	for _, r := range resv.reservations {
		if r.Seat.Section == sectionName {
			res = append(res, r)
		}
	}
	return res
}
