package booking

import (
	"errors"
	"time"

	"github.com/ffelipelimao/booking/internal/entities/ticket"
	"github.com/google/uuid"
)

var (
	ErrNumberTicketNotThanAllowed = errors.New("number of tickers is not  allowed")
	ErrEmptyPlace                 = errors.New("empty place")
	ErrInvalidBookingPayment      = errors.New("status is not enable to payment")
)

type BookingStatus string

const (
	AwaitingFunding BookingStatus = "awaiting_funding"
	Paid            BookingStatus = "paid"
	Cancelled       BookingStatus = "cancelled"
)

type Booking struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	UserID    int64            `json:"user_id"`
	Status    BookingStatus    `json:"status"`
	Tickets   []*ticket.Ticket `json:"tickets"`
	PlaceID   string           `json:"place_id"`
	Price     float64          `json:"price"`
	CreatedAt time.Time        `json:"created_at"`
}

func NewBooking(name string, userID int64, tickets []*ticket.Ticket, placeID string, price float64) (*Booking, error) {
	booking := &Booking{
		ID:        uuid.New().String(),
		Name:      name,
		UserID:    userID,
		Status:    AwaitingFunding,
		Tickets:   tickets,
		PlaceID:   placeID,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := booking.Validate(); err != nil {
		return nil, err
	}

	return booking, nil
}

func (b *Booking) Validate() error {
	if len(b.Tickets) > 4 || len(b.Tickets) <= 0 {
		return ErrNumberTicketNotThanAllowed
	}
	if b.PlaceID == "" {
		return ErrEmptyPlace
	}
	return nil
}

func (b *Booking) SetPaid() error {
	err := b.ValidatePayment()
	if err != nil {
		return err
	}

	b.Status = Paid
	return nil
}

func (b *Booking) ValidatePayment() error {
	if b.Status == Cancelled {
		return ErrInvalidBookingPayment
	}

	return nil
}
