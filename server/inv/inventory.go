package inv

import (
	"errors"
)

// Inventory is a map of section name (A/B) to Section
type Inventory struct {
	// sections []*Section
	sections map[string]*Section
}

// Define a sections as an array of seats. If nil, seat is not reserved. Array index is equivalent to seat number.
type Section struct {
	seats []*Reservation
}

type Reservation struct {
	From      string
	To        string
	Passenger Passenger
	Price     float32
}

type Passenger struct {
	FirstName string
	LastName  string
	Email     string
}

func NewInventory() *Inventory {
	s := make(map[string]*Section)
	s["A"] = &Section{seats: make([]*Reservation, 10)}
	s["B"] = &Section{seats: make([]*Reservation, 10)}
	return &Inventory{sections: s}
}

// TODO signature - use Passenger struct
func (inv *Inventory) ReserveSeat(p Passenger) (Reservation, error) {

	//TODO check for existing reservation first

	for _, section := range inv.sections {
		for j, resv := range section.seats {
			if resv != nil && resv.Passenger == p {
				return Reservation{}, errors.New("sorry you already have a reservation")
			} else if resv == nil {
				r := Reservation{
					From:      "London",
					To:        "France",
					Passenger: p,
					Price:     20.00,
				}

				section.seats[j] = &r
				return r, nil
			}
		}
	}
	return Reservation{}, errors.New("no seats available")
}
