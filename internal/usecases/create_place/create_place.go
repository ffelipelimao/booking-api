package create_place

import (
	"context"
	"time"

	"github.com/ffelipelimao/booking/internal/entities/place"
)

type CreatePlaceInputDTO struct {
	Capacity        int64 `json:"capacity"`
	AssignedSeating bool  `json:"assigned_seating"`
}

type CreatePlaceOutputDTO struct {
	ID              string    `json:"id"`
	Capacity        int64     `json:"capacity"`
	AssignedSeating bool      `json:"assigned_seating"`
	CreatedAt       time.Time `json:"created_at"`
}

type createPlaceUseCase struct {
	placeRepository PlaceRepository
}

func NewCreatePlaceUseCase(placeRepository PlaceRepository) createPlaceUseCase {
	return createPlaceUseCase{
		placeRepository: placeRepository,
	}
}

func (c createPlaceUseCase) Create(ctx context.Context, input *CreatePlaceInputDTO) (*CreatePlaceOutputDTO, error) {
	place, err := place.NewPlace(input.Capacity, input.AssignedSeating)
	if err != nil {
		return nil, err
	}

	if err := c.placeRepository.Save(ctx, place); err != nil {
		return nil, err
	}

	return &CreatePlaceOutputDTO{
		ID:              place.ID,
		Capacity:        place.Capacity,
		AssignedSeating: place.AssignedSeating,
		CreatedAt:       place.CreatedAt,
	}, nil
}
