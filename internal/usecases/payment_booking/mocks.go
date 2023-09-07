// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package payment_booking is a generated GoMock package.
package payment_booking

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

// Get mocks base method.
func (m *MockBookingRepository) Get(ctx context.Context, bookingID string) (*booking.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, bookingID)
	ret0, _ := ret[0].(*booking.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBookingRepositoryMockRecorder) Get(ctx, bookingID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBookingRepository)(nil).Get), ctx, bookingID)
}

// Update mocks base method.
func (m *MockBookingRepository) Update(ctx context.Context, bookingID string, booking *booking.Booking) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, bookingID, booking)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockBookingRepositoryMockRecorder) Update(ctx, bookingID, booking interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookingRepository)(nil).Update), ctx, bookingID, booking)
}