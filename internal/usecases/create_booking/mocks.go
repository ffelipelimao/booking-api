// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package create_booking is a generated GoMock package.
package create_booking

import (
	context "context"
	reflect "reflect"

	booking "github.com/ffelipelimao/booking/internal/entities/booking"
	gomock "github.com/golang/mock/gomock"
)

// MockBookingRepository is a mock of BookingRepository interface.
type MockBookingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookingRepositoryMockRecorder
}

// MockBookingRepositoryMockRecorder is the mock recorder for MockBookingRepository.
type MockBookingRepositoryMockRecorder struct {
	mock *MockBookingRepository
}

// NewMockBookingRepository creates a new mock instance.
func NewMockBookingRepository(ctrl *gomock.Controller) *MockBookingRepository {
	mock := &MockBookingRepository{ctrl: ctrl}
	mock.recorder = &MockBookingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookingRepository) EXPECT() *MockBookingRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockBookingRepository) Save(ctx context.Context, booking *booking.Booking) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, booking)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockBookingRepositoryMockRecorder) Save(ctx, booking interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockBookingRepository)(nil).Save), ctx, booking)
}
