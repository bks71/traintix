package inventory

type InventoryService interface {
	Purchase(firstName string, lastName string, email string) (Reservation, error)
}
