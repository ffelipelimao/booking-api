package web

import (
	"context"

	"github.com/ffelipelimao/booking/internal/usecases/create_place"
)

//go:generate mockgen -destination=./mocks.go -source=./interfaces.go -package=web
type CreatePlaceUseCase interface {
	Create(ctx context.Context, InputPlace *create_place.CreatePlaceInputDTO) (*create_place.CreatePlaceOutputDTO, error)
}
