// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package create_place is a generated GoMock package.
package create_place

import (
	context "context"
	reflect "reflect"

	place "github.com/ffelipelimao/booking/internal/entities/place"
	gomock "github.com/golang/mock/gomock"
)

// MockPlaceRepository is a mock of PlaceRepository interface.
type MockPlaceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPlaceRepositoryMockRecorder
}

// MockPlaceRepositoryMockRecorder is the mock recorder for MockPlaceRepository.
type MockPlaceRepositoryMockRecorder struct {
	mock *MockPlaceRepository
}

// NewMockPlaceRepository creates a new mock instance.
func NewMockPlaceRepository(ctrl *gomock.Controller) *MockPlaceRepository {
	mock := &MockPlaceRepository{ctrl: ctrl}
	mock.recorder = &MockPlaceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlaceRepository) EXPECT() *MockPlaceRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockPlaceRepository) Save(ctx context.Context, place *place.Place) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, place)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockPlaceRepositoryMockRecorder) Save(ctx, place interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPlaceRepository)(nil).Save), ctx, place)
}
