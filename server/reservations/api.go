package reservations

type ReservationSystem interface {
	ReserveSeat(Passenger) (*Reservation, error)
	GetReservation(Passenger) (*Reservation, error)
	CancelReservation(Passenger) error
	ChangeSeat(Passenger, Seat) (*Reservation, error)
}
