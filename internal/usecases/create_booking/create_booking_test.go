package create_booking_test

import (
	"context"
	"testing"

	"github.com/ffelipelimao/booking/internal/entities/booking"
	"github.com/ffelipelimao/booking/internal/entities/ticket"

	"github.com/ffelipelimao/booking/internal/usecases/create_booking"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateBooking(t *testing.T) {
	input := &create_booking.CreateBookingInputDTO{
		Name:    "test",
		PlaceID: "1234",
		Tickets: []*ticket.Ticket{
			{
				ID:   1,
				Name: "1",
			},
		},
	}

	invalidInput := &create_booking.CreateBookingInputDTO{
		Name:    "test",
		PlaceID: "1234",
	}

	tests := []struct {
		name          string
		input         *create_booking.CreateBookingInputDTO
		expectedError error
		executeMock   func(m *create_booking.MockBookingRepository)
	}{
		{
			name:          "Should create place with success",
			expectedError: nil,
			input:         input,
			executeMock: func(m *create_booking.MockBookingRepository) {
				// TODO: Refactor to be able to know if mock was execute with right parameters
				m.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:          "Should not create booking invalid",
			expectedError: booking.ErrNumberTicketNotThanAllowed,
			input:         invalidInput,
			executeMock:   func(m *create_booking.MockBookingRepository) {},
		},
		{
			name:          "Should return an error in repository",
			expectedError: assert.AnError,
			input:         input,
			executeMock: func(m *create_booking.MockBookingRepository) {
				// TODO: Refactor to be able to know if mock was execute with right parameters
				m.EXPECT().Save(gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bookingRepositoryMock := create_booking.NewMockBookingRepository(ctrl)

			tt.executeMock(bookingRepositoryMock)

			createBookingUseCase := create_booking.NewCreateBookingUseCase(bookingRepositoryMock)

			output, err := createBookingUseCase.Create(context.Background(), tt.input)

			assert.Equal(t, tt.expectedError, err)

			if output != nil {
				assert.NotNil(t, output)
				assert.NotEmpty(t, output.ID)
			}
		})
	}
}
