package create_booking

import (
	"context"
	"time"

	"github.com/ffelipelimao/booking/internal/entities/booking"
	"github.com/ffelipelimao/booking/internal/entities/ticket"
)

type CreateBookingInputDTO struct {
	Name    string           `json:"name"`
	UserID  int64            `json:"user_id"`
	Tickets []*ticket.Ticket `json:"tickets"`
	PlaceID string           `json:"place_id"`
	Price   float64          `json:"price"`
}

type CreateBookingOutputDTO struct {
	ID        string                `json:"id"`
	Name      string                `json:"name"`
	UserID    int64                 `json:"user_id"`
	Status    booking.BookingStatus `json:"status"`
	Tickets   []*ticket.Ticket      `json:"tickets"`
	PlaceID   string                `json:"place_id"`
	Price     float64               `json:"price"`
	CreatedAt time.Time             `json:"created_at"`
}

type createBookingUseCase struct {
	bookingRepository BookingRepository
}

func NewCreateBookingUseCase(bookingRepository BookingRepository) *createBookingUseCase {
	return &createBookingUseCase{
		bookingRepository: bookingRepository,
	}
}

func (c *createBookingUseCase) Create(ctx context.Context, input *CreateBookingInputDTO) (*CreateBookingOutputDTO, error) {
	booking, err := booking.NewBooking(input.Name, input.UserID, input.Tickets, input.PlaceID, input.Price)
	if err != nil {
		return nil, err
	}

	if err := c.bookingRepository.Save(ctx, booking); err != nil {
		return nil, err
	}

	return &CreateBookingOutputDTO{
		ID:        booking.ID,
		Name:      booking.Name,
		UserID:    booking.UserID,
		Status:    booking.Status,
		Tickets:   booking.Tickets,
		PlaceID:   booking.PlaceID,
		Price:     booking.Price,
		CreatedAt: booking.CreatedAt,
	}, nil
}
