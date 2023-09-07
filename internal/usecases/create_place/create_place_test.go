package create_place_test

import (
	"context"
	"testing"

	"github.com/ffelipelimao/booking/internal/entities/place"
	"github.com/ffelipelimao/booking/internal/usecases/create_place"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreatePlace(t *testing.T) {
	input := &create_place.CreatePlaceInputDTO{
		Capacity:        1,
		AssignedSeating: false,
	}

	invalidInput := &create_place.CreatePlaceInputDTO{
		Capacity:        -1,
		AssignedSeating: false,
	}

	tests := []struct {
		name          string
		input         *create_place.CreatePlaceInputDTO
		expectedError error
		executeMock   func(m *create_place.MockPlaceRepository)
	}{
		{
			name:          "Should create place with success",
			expectedError: nil,
			input:         input,
			executeMock: func(m *create_place.MockPlaceRepository) {
				// TODO: Refactor to be able to know if mock was execute with right parameters
				m.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:          "Should not create place invalid",
			expectedError: place.ErrInvalidCapacity,
			input:         invalidInput,
			executeMock:   func(m *create_place.MockPlaceRepository) {},
		},
		{
			name:          "Should return an error in repository",
			expectedError: assert.AnError,
			input:         input,
			executeMock: func(m *create_place.MockPlaceRepository) {
				// TODO: Refactor to be able to know if mock was execute with right parameters
				m.EXPECT().Save(gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			placeRepositoryMock := create_place.NewMockPlaceRepository(ctrl)

			tt.executeMock(placeRepositoryMock)

			createPlaceUseCase := create_place.NewCreatePlaceUseCase(placeRepositoryMock)

			output, err := createPlaceUseCase.Create(context.Background(), tt.input)

			assert.Equal(t, tt.expectedError, err)

			if output != nil {
				assert.NotNil(t, output)
				assert.NotEmpty(t, output.ID)
			}
		})
	}
}
