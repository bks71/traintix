package inventory

import (
	"testing"
)

func TestPurchase(t *testing.T) {
	inv := NewInventory()

	for i := 0; i < 20; i++ {
		receipt, err := inv.Purchase("B", "S", "a@b.com")
		if err != nil {
			t.Errorf("Error purchasing seat: %v", err)
		}

		t.Log(receipt)
	}

	_, err := inv.Purchase("B", "S", "dfd")
	t.Log(err)
}
