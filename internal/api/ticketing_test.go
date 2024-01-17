package api

import (
	"context"
	"testing"

	"github.com/bks71/traintix/internal/inventory"
	"github.com/bks71/traintix/pb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPurchase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a context for the test
	ctx := context.Background()

	inv := inventory.NewMockInventoryService(ctrl)

	server := &TicketingServer{
		InventoryService: inv,
	}

	inv.EXPECT().Purchase(gomock.Any(), gomock.Any(), gomock.Any()) //TODO

	// Create a purchase request
	req := &pb.PurchaseRequest{
		User: &pb.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
	}

	// Call the Purchase method
	resp, err := server.Purchase(ctx, req)

	// Check for errors
	assert.NoError(t, err)

	// Check the response
	assert.NotNil(t, resp)
	// assert.NotNil(t, resp.Receipt)
	// assert.Equal(t, "London", resp.Receipt.From)
	// assert.Equal(t, "France", resp.Receipt.To)
	// assert.Equal(t, "John", resp.Receipt.User.FirstName)
	// assert.Equal(t, "Doe", resp.Receipt.User.LastName)
	// assert.Equal(t, "john.doe@example.com", resp.Receipt.User.Email)
	// assert.NotZero(t, resp.Receipt.Price)
	// assert.NotNil(t, resp.Receipt.Seat)
	// assert.NotZero(t, resp.Receipt.Seat.SeatNumber)
	// assert.NotEmpty(t, resp.Receipt.Seat.Section)
}
