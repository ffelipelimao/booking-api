package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ffelipelimao/booking/internal/entities/booking"
)

var ErrBookingNotFound = errors.New("booking not find")

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *bookingRepository {
	return &bookingRepository{
		db: db,
	}
}

func (r *bookingRepository) Save(ctx context.Context, booking *booking.Booking) error {
	stmt, err := r.db.Prepare(`INSERT INTO booking (id, name, user_id, status, place_id, price, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = r.db.Exec(booking.ID,
		booking.Name,
		booking.UserID,
		booking.Status,
		booking.PlaceID,
		booking.Price,
		booking.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *bookingRepository) Update(booking *booking.Booking) error {
	stmt, err := r.db.Prepare(`UPDATE booking
		SET name = ?, user_id = ?, status = ?, place_id = ?, price = ?, created_at = ?
		WHERE id = ?
	`)

	defer stmt.Close()

	_, err = r.db.Exec(
		booking.Name,
		booking.UserID,
		booking.Status,
		booking.PlaceID,
		booking.Price,
		booking.CreatedAt,
		booking.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *bookingRepository) Get(ctx context.Context, bookingID string) (*booking.Booking, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, name, user_id, status, place_id, price, created_at
		FROM booking
		WHERE id = ?
	`)

	defer stmt.Close()

	booking := &booking.Booking{}

	err = r.db.QueryRow(bookingID).Scan(
		&booking.ID,
		&booking.Name,
		&booking.UserID,
		&booking.Status,
		&booking.PlaceID,
		&booking.Price,
		&booking.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrBookingNotFound
		}
		return nil, err
	}

	return booking, nil
}
