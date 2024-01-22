package reservations

import (
	"fmt"
	"testing"

	"github.com/bks71/traintix/pb"
	"github.com/stretchr/testify/assert"
	rand "github.com/xyproto/randomstring"
)

func randomPassenger() *pb.Passenger {
	return &pb.Passenger{
		FirstName: rand.HumanFriendlyString(5),
		LastName:  rand.HumanFriendlyString(5),
		Email:     fmt.Sprintf("%s@trainz.com", rand.HumanFriendlyString(5)),
	}
}

func TestResvSystem(t *testing.T) {

	t.Run("Fill the train", func(t *testing.T) {
		resv := NewReservationSystem()

		// add 20 passengers
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
			assert.True(t, r.Seat.Number >= 1 && r.Seat.Number <= 10)
			assert.True(t, r.Seat.Section == "A" || r.Seat.Section == "B")

			assert.Equal(t, i+1, len(resv.GetReservationsBySection("A"))+len(resv.GetReservationsBySection("B")))
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

		savedRes, err := resv.GetReservation(p.Email)
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

		r, err := resv.CancelReservation(p.Email)
		if err != nil {
			t.Errorf("Error canceling a reservation: %v", err)
		}

		assert.True(t, r.Passenger == nil)
	})

	t.Run("Reservation not found", func(t *testing.T) {
		resv := NewReservationSystem()

		// check for error on get
		_, errGet := resv.GetReservation(randomPassenger().Email)
		assert.Error(t, errGet)
		t.Log(errGet)

		// check for error on cancel
		_, errCancel := resv.CancelReservation(randomPassenger().Email)
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
		r2, err := resv.ChangeSeat(p.Email, r.Seat)
		assert.Error(t, err)
		assert.Equal(t, r, r2)
		t.Log(err)

		// change the seat
		newSeat := &pb.Seat{Number: 2, Section: "A"}
		newRes, err := resv.ChangeSeat(p.Email, newSeat)
		if err != nil {
			t.Errorf("Error changing a seat: %v", err)
		}

		assert.Equal(t, newSeat.Number, newRes.Seat.Number)
		assert.Equal(t, newSeat.Section, newRes.Seat.Section)
	})

	t.Run("Fill the train, do some cancels, etc.", func(t *testing.T) {
		resv := NewReservationSystem()

		passengers := make([]*pb.Passenger, 0)

		// add 20 passengers
		for i := 0; i < 20; i++ {
			passengers = append(passengers, randomPassenger())
			_, err := resv.ReserveSeat(passengers[i])
			if err != nil {
				t.Errorf("Error reserving a seat: %v", err)
			}
		}

		// do some cancels
		_, err := resv.CancelReservation(passengers[0].Email)
		if err != nil {
			t.Errorf("Error cancelling: %v", err)
		}

		_, err = resv.CancelReservation(passengers[1].Email)
		if err != nil {
			t.Errorf("Error cancelling: %v", err)
		}

		_, err = resv.CancelReservation(passengers[9].Email)
		if err != nil {
			t.Errorf("Error cancelling: %v", err)
		}

		_, err = resv.CancelReservation(passengers[10].Email)
		if err != nil {
			t.Errorf("Error cancelling: %v", err)
		}

		_, err = resv.CancelReservation(passengers[19].Email)
		if err != nil {
			t.Errorf("Error cancelling: %v", err)
		}

		r, err := resv.ReserveSeat(passengers[0])
		if err != nil {
			t.Errorf("Error reserving a seat: %v", err)
		}
		assert.True(t, r.Seat != nil)

	})

}
