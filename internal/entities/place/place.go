package place

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrInvalidCapacity = errors.New("invalid capacity")

type Place struct {
	ID              string
	Capacity        int64
	AssignedSeating bool
	CreatedAt       time.Time
}

func NewPlace(capacity int64, assignedSeating bool) (*Place, error) {
	place := &Place{
		ID:              uuid.New().String(),
		Capacity:        capacity,
		AssignedSeating: assignedSeating,
		CreatedAt:       time.Now(),
	}

	if err := place.Validate(); err != nil {
		return nil, err
	}

	return place, nil
}

func (p *Place) Validate() error {
	if p.Capacity <= 0 {
		return ErrInvalidCapacity
	}

	return nil
}
