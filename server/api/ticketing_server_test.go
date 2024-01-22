package api

import (
	"context"
	"testing"

	"github.com/bks71/traintix/pb"
	resv "github.com/bks71/traintix/server/reservations"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPurchase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a context for the test
	ctx := context.Background()

	// Set up the mocks
	p := pb.Passenger{
		FirstName: "Brandon",
		LastName:  "Stewart",
		Email:     "bkstewart@gmail.com",
	}

	mockRes := resv.NewMockReservationSystem(ctrl)

	server := &TicketingServer{
		ReservationSystem: mockRes,
	}

	r := pb.Reservation{
		From:      "Here",
		To:        "There",
		Seat:      &pb.Seat{Section: "A", Number: 1},
		Passenger: &p,
		Price:     2000.00,
	}

	mockRes.EXPECT().
		ReserveSeat(&p).
		Return(&r, nil)

	// Create a purchase request
	req := &pb.PurchaseRequest{
		Passenger: &p,
	}

	// Call the Purchase method
	resp, err := server.Purchase(ctx, req)

	// Check for errors
	assert.NoError(t, err)

	// Check the response
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Reservation)
}
