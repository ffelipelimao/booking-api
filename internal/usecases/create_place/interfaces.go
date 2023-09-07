package create_place

import (
	"context"

	"github.com/ffelipelimao/booking/internal/entities/place"
)

//go:generate mockgen -destination=./mocks.go -source=./interfaces.go -package=create_place
type PlaceRepository interface {
	Save(ctx context.Context, place *place.Place) error
}
