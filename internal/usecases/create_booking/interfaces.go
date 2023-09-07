package create_booking

import (
	"context"

	"github.com/ffelipelimao/booking/internal/entities/booking"
)

//go:generate mockgen -destination=./mocks.go -source=./interfaces.go -package=create_booking
type BookingRepository interface {
	Save(ctx context.Context, booking *booking.Booking) error
}
