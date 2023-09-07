package repositories

import (
	"context"
	"database/sql"

	"github.com/ffelipelimao/booking/internal/entities/place"
)

type placeRepository struct {
	db *sql.DB
}

func NewPlaceRepository(db *sql.DB) *placeRepository {
	return &placeRepository{
		db: db,
	}
}

func (r *placeRepository) Save(ctx context.Context, place *place.Place) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO place (id, capacity, assigned_seating, created_at) 
		VALUES (?,?,?,?)`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		place.ID,
		place.Capacity,
		place.AssignedSeating,
		place.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
