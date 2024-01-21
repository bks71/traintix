package reservations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	rand "github.com/xyproto/randomstring"
)

func randomPassenger() Passenger {
	return Passenger{
		FirstName: rand.HumanFriendlyString(5),
		LastName:  rand.HumanFriendlyString(5),
		Email:     fmt.Sprintf("%s@trainz.com", rand.HumanFriendlyString(5)),
	}
}

func TestPurchase(t *testing.T) {

	t.Run("Fill the train", func(t *testing.T) {
		resv := NewReservationSystem()

		for i := 0; i < 20; i++ {
			p := randomPassenger()
			r, err := resv.ReserveSeat(p)
			if err != nil {
				t.Errorf("Error reserving a seat: %v", err)
			}

			t.Log(r)

			assert.Equal(t, "London", r.From)
			assert.Equal(t, "France", r.To)
			assert.Equal(t, float32(20.00), r.Price)
			assert.True(t, r.Seat.number >= 1 && r.Seat.number <= 10)
			assert.True(t, r.Seat.section == "A" || r.Seat.section == "B")
		}

		// check for error when train is full
		_, err := resv.ReserveSeat(randomPassenger())
		assert.Error(t, err)

		t.Log(err)
	})

	t.Run("Get an existing reservation", func(t *testing.T) {
		resv := NewReservationSystem()

		p := randomPassenger()

		newRes, err := resv.ReserveSeat(p)
		if err != nil {
			t.Errorf("Error reserving a seat: %v", err)
		}

		savedRes, err := resv.GetReservation(p)
		if err != nil {
			t.Errorf("Error retrieving a reservation: %v", err)
		}

		assert.Equal(t, newRes, savedRes)

	})

	t.Run("Cancel a reservation", func(t *testing.T) {
		resv := NewReservationSystem()

		p := randomPassenger()

		_, err := resv.ReserveSeat(p)
		if err != nil {
			t.Errorf("Error reserving a seat: %v", err)
		}

		err = resv.CancelReservation(p)
		if err != nil {
			t.Errorf("Error canceling a reservation: %v", err)
		}
	})

	t.Run("Reservation not found", func(t *testing.T) {
		resv := NewReservationSystem()

		// check for error on get
		_, errGet := resv.GetReservation(randomPassenger())
		assert.Error(t, errGet)
		t.Log(errGet)

		// check for error on cancel
		errCancel := resv.CancelReservation(randomPassenger())
		assert.Error(t, errCancel)
		t.Log(errCancel)
	})

	t.Run("Change a seat", func(t *testing.T) {
		resv := NewReservationSystem()

		p := randomPassenger()

		r, err := resv.ReserveSeat(p)
		if err != nil {
			t.Errorf("Error reserving a seat: %v", err)
		}

		// test that passenger can't change to the same seat
		_, err = resv.ChangeSeat(p, r.Seat)
		assert.Error(t, err)
		t.Log(err)

		// change the seat
		newSeat := Seat{number: 5, section: "B"}
		newRes, err := resv.ChangeSeat(p, newSeat)
		if err != nil {
			t.Errorf("Error changing a seat: %v", err)
		}

		assert.Equal(t, newSeat, newRes.Seat)
	})

}
