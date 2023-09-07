package payment_booking

import (
	"context"

	"github.com/ffelipelimao/booking/internal/entities/booking"
)

type PaymentBookingInputDTO struct {
	ID     string                `json:"id"`
	Status booking.BookingStatus `json:"status"`
	UserID int64                 `json:"user_id"`
}

type PaymentBookingOutputDTO struct {
	ID     string                `json:"id"`
	Status booking.BookingStatus `json:"status"`
	UserID int64                 `json:"user_id"`
}

type paymentBookingUseCase struct {
	bookingRepository BookingRepository
}

func NewPaymentBookingUseCase(bookingRepository BookingRepository) *paymentBookingUseCase {
	return &paymentBookingUseCase{
		bookingRepository: bookingRepository,
	}
}

func (c *paymentBookingUseCase) Execute(ctx context.Context, input *PaymentBookingInputDTO) (*PaymentBookingOutputDTO, error) {
	booking, err := c.bookingRepository.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if err := booking.SetPaid(); err != nil {
		return nil, err
	}

	if err := c.bookingRepository.Update(ctx, booking.ID, booking); err != nil {
		return nil, err
	}

	return &PaymentBookingOutputDTO{
		ID:     booking.ID,
		Status: booking.Status,
		UserID: booking.UserID,
	}, nil
}
