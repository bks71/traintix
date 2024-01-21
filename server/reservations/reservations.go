package reservations

import (
	"errors"
)

var inventory = []Seat{
	{section: "A", number: 1},
	{section: "A", number: 2},
	{section: "A", number: 3},
	{section: "A", number: 4},
	{section: "A", number: 5},
	{section: "A", number: 6},
	{section: "A", number: 7},
	{section: "A", number: 8},
	{section: "A", number: 9},
	{section: "A", number: 10},
	{section: "B", number: 1},
	{section: "B", number: 2},
	{section: "B", number: 3},
	{section: "B", number: 4},
	{section: "B", number: 5},
	{section: "B", number: 6},
	{section: "B", number: 7},
	{section: "B", number: 8},
	{section: "B", number: 9},
	{section: "B", number: 10},
}

type Seat struct {
	section string
	number  int
}

type Reservation struct {
	From  string
	To    string
	Price float32
	Seat  Seat
}

type Passenger struct {
	FirstName string
	LastName  string
	Email     string
}

type ResvSys struct {
	inventory  []Seat
	passengers map[Passenger]Reservation
}

func NewReservationSystem() *ResvSys {
	return &ResvSys{
		inventory:  inventory,
		passengers: make(map[Passenger]Reservation),
	}
}

func (resv *ResvSys) reserveAnySeat() (Seat, error) {
	if len(resv.inventory) == 0 {
		return Seat{}, errors.New("no seats available")
	}
	i := resv.inventory
	s := i[len(i)-1]
	resv.inventory = i[:len(i)-1]
	return s, nil
}

func (resv *ResvSys) ReserveSeat(p Passenger) (*Reservation, error) {

	//check for an existing reservation for this passenger
	if r, exists := resv.passengers[p]; exists {
		return &r, errors.New("sorry you already have a reservation")
	}

	//Find an empty seat
	s, err := resv.reserveAnySeat()
	if err != nil {
		return nil, err
	}

	//Reserve it
	r := Reservation{From: "London", To: "France", Price: 20.00, Seat: s}
	resv.passengers[p] = r

	return &r, nil
}

func (resv *ResvSys) GetReservation(p Passenger) (*Reservation, error) {
	if r, exists := resv.passengers[p]; exists {
		return &r, nil
	}
	return nil, errors.New("no reservation found")
}

func (resv *ResvSys) CancelReservation(p Passenger) error {
	if _, exists := resv.passengers[p]; exists {
		delete(resv.passengers, p)
		return nil
	}
	return errors.New("no reservation found")
}

func (resv *ResvSys) ChangeSeat(p Passenger, s Seat) (*Reservation, error) {

	r, exists := resv.passengers[p]
	if !exists {
		return nil, errors.New("no reservation found for this passenger")
	}

	if r.Seat == s {
		return &r, errors.New("passenger has already reserved this seat")
	}

	for i, v := range resv.inventory {
		if v == s {
			// put old seat back into the inventory and switch passenger to new seat
			resv.inventory[i] = r.Seat
			r.Seat = s
			return &r, nil
		}
	}

	return nil, errors.New("seat not available")
}
