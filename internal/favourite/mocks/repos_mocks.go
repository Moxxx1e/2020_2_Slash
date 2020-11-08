// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_favourite is a generated GoMock package.
package mocks

import (
	models "github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFavouriteRepository is a mock of FavouriteRepository interface
type MockFavouriteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFavouriteRepositoryMockRecorder
}

// MockFavouriteRepositoryMockRecorder is the mock recorder for MockFavouriteRepository
type MockFavouriteRepositoryMockRecorder struct {
	mock *MockFavouriteRepository
}

// NewMockFavouriteRepository creates a new mock instance
func NewMockFavouriteRepository(ctrl *gomock.Controller) *MockFavouriteRepository {
	mock := &MockFavouriteRepository{ctrl: ctrl}
	mock.recorder = &MockFavouriteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFavouriteRepository) EXPECT() *MockFavouriteRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockFavouriteRepository) Insert(favourite *models.Favourite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", favourite)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockFavouriteRepositoryMockRecorder) Insert(favourite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockFavouriteRepository)(nil).Insert), favourite)
}

// Select mocks base method
func (m *MockFavouriteRepository) Select(favourite *models.Favourite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", favourite)
	ret0, _ := ret[0].(error)
	return ret0
}

// Select indicates an expected call of Select
func (mr *MockFavouriteRepositoryMockRecorder) Select(favourite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockFavouriteRepository)(nil).Select), favourite)
}

// SelectFavouriteContent mocks base method
func (m *MockFavouriteRepository) SelectFavouriteContent(userID uint64) ([]*models.Content, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectFavouriteContent", userID)
	ret0, _ := ret[0].([]*models.Content)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectFavouriteContent indicates an expected call of SelectFavouriteContent
func (mr *MockFavouriteRepositoryMockRecorder) SelectFavouriteContent(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectFavouriteContent", reflect.TypeOf((*MockFavouriteRepository)(nil).SelectFavouriteContent), userID)
}

// Delete mocks base method
func (m *MockFavouriteRepository) Delete(favourite *models.Favourite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", favourite)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockFavouriteRepositoryMockRecorder) Delete(favourite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFavouriteRepository)(nil).Delete), favourite)
}