package inventory

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	rand "github.com/xyproto/randomstring"
)

func TestPurchase(t *testing.T) {

	t.Run("Fill the train", func(t *testing.T) {
		inv := NewInventory()

		for i := 0; i < 20; i++ {
			fname := rand.HumanFriendlyString(5)
			lname := rand.HumanFriendlyString(5)
			email := fmt.Sprintf("%s@trainz.com", rand.HumanFriendlyString(5))
			resv, err := inv.Purchase(fname, lname, email)
			if err != nil {
				t.Errorf("Error purchasing seat: %v", err)
			}

			t.Log(resv)

			// check for correct passenger info
			assert.Equal(t, fname, resv.Passenger.FirstName)
			assert.Equal(t, lname, resv.Passenger.LastName)
			assert.Equal(t, email, resv.Passenger.Email)

			// check from/to
			assert.Equal(t, "London", resv.From)
			assert.Equal(t, "France", resv.To)

			// check for correct seat and section
			assert.True(t, resv.SeatNumber >= 0 && resv.SeatNumber <= 9)
			if i < 10 {
				assert.Equal(t, "A", resv.Section)
			} else {
				assert.Equal(t, "B", resv.Section)
			}

		}

		// check for error when train is full
		_, err := inv.Purchase("B", "S", "dfd")
		assert.Error(t, err)

		t.Log(err)
	})
}
