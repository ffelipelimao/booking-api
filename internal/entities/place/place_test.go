package place_test

import (
	"testing"

	"github.com/ffelipelimao/booking/internal/entities/place"
	"github.com/stretchr/testify/assert"
)

func Test_Place_Validate(t *testing.T) {
	invalidPlace := &place.Place{Capacity: 0}
	err := invalidPlace.Validate()
	assert.NotNil(t, err)

	validPlace := &place.Place{Capacity: 60000}
	err = validPlace.Validate()
	assert.Nil(t, err)
}
