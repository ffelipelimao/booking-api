package payment_booking

import (
	"context"

	"github.com/ffelipelimao/booking/internal/entities/booking"
)

//go:generate mockgen -destination=./mocks.go -source=./interfaces.go -package=payment_booking
type BookingRepository interface {
	Update(ctx context.Context, bookingID string, booking *booking.Booking) error
	Get(ctx context.Context, bookingID string) (*booking.Booking, error)
}
