package inventory

import (
	"testing"
)

func TestPurchase(t *testing.T) {
	inv := NewInventory()

	// Test case 1: Purchase a seat in section "a"
	receipt, err := inv.Purchase("Brandon", "Stewart", "bkstewart@gmail.com")
	if err != nil {
		t.Errorf("Error purchasing seat: %v", err)
	}

	t.Log(receipt)

}
