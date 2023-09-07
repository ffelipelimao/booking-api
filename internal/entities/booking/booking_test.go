package booking

import (
	"testing"

	"github.com/ffelipelimao/booking/internal/entities/ticket"
	"github.com/stretchr/testify/assert"
)

func Test_Booking_New(t *testing.T) {
	tickets := []*ticket.Ticket{
		{
			ID:   1,
			Name: "name",
		},
	}

	booking, err := NewBooking("name", 123, tickets, "place_id", 12.34)

	assert.Nil(t, err)
	assert.Equal(t, booking.Name, "name")
}

func Test_Booking_InvalidNumbersTickets(t *testing.T) {
	tickets := []*ticket.Ticket{
		{
			ID:   1,
			Name: "name",
		},
		{
			ID:   1,
			Name: "name",
		},
		{
			ID:   1,
			Name: "name",
		},
		{
			ID:   1,
			Name: "name",
		},
		{
			ID:   1,
			Name: "name",
		},
	}

	_, err := NewBooking("name", 123, tickets, "place_id", 12.34)

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNumberTicketNotThanAllowed)
}

func Test_Booking_InvalidPlaceID(t *testing.T) {
	tickets := []*ticket.Ticket{
		{
			ID:   1,
			Name: "name",
		},
	}

	_, err := NewBooking("name", 123, tickets, "", 12.34)

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrEmptyPlace)
}
