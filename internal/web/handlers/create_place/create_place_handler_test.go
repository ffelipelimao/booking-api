package web_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ffelipelimao/booking/internal/entities/place"
	"github.com/ffelipelimao/booking/internal/usecases/create_place"
	"github.com/ffelipelimao/booking/internal/web"
	web_handlers "github.com/ffelipelimao/booking/internal/web/handlers/create_place"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreatePlace_Handler(t *testing.T) {
	input := &create_place.CreatePlaceInputDTO{
		Capacity:        1,
		AssignedSeating: false,
	}

	invalidInput := &create_place.CreatePlaceInputDTO{
		Capacity:        -1,
		AssignedSeating: false,
	}

	tests := []struct {
		name               string
		inputBody          string
		expectedStatusCode int
		executeMock        func(m *web.MockCreatePlaceUseCase)
	}{
		{
			name:               "Should create place with success",
			inputBody:          `{"capacity": 1,"assigned_seating": false}`,
			expectedStatusCode: http.StatusOK,
			executeMock: func(m *web.MockCreatePlaceUseCase) {
				m.EXPECT().Create(gomock.Any(), input).Return(&create_place.CreatePlaceOutputDTO{}, nil)
			},
		},
		{
			name:               "Should be bad request",
			inputBody:          `{`,
			expectedStatusCode: http.StatusBadRequest,
			executeMock:        func(m *web.MockCreatePlaceUseCase) {},
		},
		{
			name:               "Should be error by invalid capacity",
			inputBody:          `{"capacity": -1,"assigned_seating": false}`,
			expectedStatusCode: http.StatusBadRequest,
			executeMock: func(m *web.MockCreatePlaceUseCase) {
				m.EXPECT().Create(gomock.Any(), invalidInput).Return(nil, place.ErrInvalidCapacity)
			},
		},
		{
			name:               "Should be error unknowing error",
			inputBody:          `{"capacity": 1,"assigned_seating": false}`,
			expectedStatusCode: http.StatusInternalServerError,
			executeMock: func(m *web.MockCreatePlaceUseCase) {
				m.EXPECT().Create(gomock.Any(), input).Return(nil, assert.AnError)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			createPlaceUseCaseMock := web.NewMockCreatePlaceUseCase(ctrl)

			tt.executeMock(createPlaceUseCaseMock)

			endpoint := "/v1/places"
			request := httptest.NewRequest("POST", endpoint, strings.NewReader(tt.inputBody))
			request.Header.Set("Content-Type", "application/json")

			appFake := fiber.New()

			handlerCreatePlace := web_handlers.NewCreatePlaceHandler(createPlaceUseCaseMock)
			appFake.Post(endpoint, handlerCreatePlace.Handle)

			defer appFake.Shutdown()
			res, _ := appFake.Test(request)

			assert.Equal(t, tt.expectedStatusCode, res.StatusCode)
		})
	}
}
