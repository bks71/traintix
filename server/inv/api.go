package inv

type InventoryService interface {
	ReserveSeat(firstName string, lastName string, email string) (Reservation, error)
}
