package inventory

import "errors"

// Define sections as arrays of users. Array index is equivalent to seat number.
type Section struct {
	name         string
	reservations []*Reservation
}

type Inventory struct {
	sections []*Section
}

type Passenger struct {
	FirstName string
	LastName  string
	Email     string
}

type Reservation struct {
	From       string
	To         string
	Passenger  Passenger
	Section    string
	SeatNumber int
	Price      float32
}

func NewInventory() *Inventory {
	return &Inventory{
		sections: []*Section{
			{name: "A", reservations: make([]*Reservation, 10)},
			{name: "B", reservations: make([]*Reservation, 10)},
		},
	}

}

func (inv *Inventory) Purchase(firstName string, lastName string, email string) (Reservation, error) {
	p := Passenger{FirstName: firstName, LastName: lastName, Email: email}
	for _, section := range inv.sections {
		for j, resv := range section.reservations {
			if resv != nil && resv.Passenger == p {
				return Reservation{}, errors.New("sorry you already have a reservation")
			} else if resv == nil {
				r := Reservation{
					From:       "London",
					To:         "France",
					Passenger:  p,
					Section:    section.name,
					SeatNumber: j,
					Price:      20.00,
				}

				section.reservations[j] = &r
				return r, nil
			}
		}
	}
	return Reservation{}, errors.New("no seats available")
}
