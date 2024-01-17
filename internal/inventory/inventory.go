package inventory

import "errors"

// "fmt"
// "os"

// Define sections as arrays of users. Array index is equivalent to seat number.
type Section struct {
	name         string
	reservations []*Reservation
}

type Inventory struct {
	sections []*Section
}

type Reservation struct {
	FirstName string
	LastName  string
	Email     string
}

type Receipt struct {
	From        string
	To          string
	Reservation Reservation
	Section     string
	Seat        int
	Price       float32
}

func NewInventory() *Inventory {
	return &Inventory{
		sections: []*Section{
			{name: "a", reservations: make([]*Reservation, 10)},
			{name: "b", reservations: make([]*Reservation, 10)},
		},
	}

}

func (inv *Inventory) Purchase(firstName string, lastName string, email string) (Receipt, error) {
	for _, section := range inv.sections {
		for j, resv := range section.reservations {
			// fmt.Fprint(os.Stdout, i, j, resv)
			if resv == nil {
				r := &Reservation{
					FirstName: firstName,
					LastName:  lastName,
					Email:     email,
				}
				section.reservations[j] = r
				return Receipt{
					From:        "London",
					To:          "France",
					Reservation: *r,
					Section:     section.name,
					Seat:        j,
					Price:       20.00,
				}, nil
			}
			// if user == nil {
			// 	section.seats[j] = user
			// 	return pb.Receipt{Seat: j, Section: section.name}, nil
			// }
		}
		// if  == nil {
		// 	i.a[i] = user
		// 	return &pb.Receipt{Seat: i, Section: "a"}, nil
		// }
	}
	return Receipt{}, errors.New("No seats available")
}
