package api

import (
	"context"
	"testing"

	"github.com/bks71/traintix/pb"
	"github.com/bks71/traintix/server/inventory"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPurchase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a context for the test
	ctx := context.Background()

	// Set up the mocks
	user := pb.User{
		FirstName: "Brandon",
		LastName:  "Stewart",
		Email:     "bkstewart@gmail.com",
	}

	inv := inventory.NewMockInventoryService(ctrl)

	server := &TicketingServer{
		InventoryService: inv,
	}

	resv := inventory.Reservation{
		From:       "Here",
		To:         "There",
		Passenger:  inventory.Passenger{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email},
		Section:    "X",
		SeatNumber: 100,
		Price:      2000.00}

	inv.EXPECT().Purchase(user.FirstName, user.LastName, user.Email).Return(resv, nil)

	// Create a purchase request
	req := &pb.PurchaseRequest{
		User: &user,
	}

	// Call the Purchase method
	resp, err := server.Purchase(ctx, req)

	// Check for errors
	assert.NoError(t, err)

	// Check the response
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Receipt)
	assert.Equal(t, resv.From, resp.Receipt.From)
	assert.Equal(t, resv.To, resp.Receipt.To)
	assert.Equal(t, user.FirstName, resp.Receipt.User.FirstName)
	assert.Equal(t, user.LastName, resp.Receipt.User.LastName)
	assert.Equal(t, user.Email, resp.Receipt.User.Email)
	assert.Equal(t, resv.Price, resp.Receipt.Price)
	assert.Equal(t, int32(resv.SeatNumber), resp.Receipt.Seat.SeatNumber)
	assert.Equal(t, resv.Section, resp.Receipt.Seat.Section)
}
