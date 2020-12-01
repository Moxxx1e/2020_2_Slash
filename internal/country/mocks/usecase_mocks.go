// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_country is a generated GoMock package.
package mocks

import (
	errors "github.com/go-park-mail-ru/2020_2_Slash/internal/helpers/errors"
	models "github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCountryUsecase is a mock of CountryUsecase interface
type MockCountryUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCountryUsecaseMockRecorder
}

// MockCountryUsecaseMockRecorder is the mock recorder for MockCountryUsecase
type MockCountryUsecaseMockRecorder struct {
	mock *MockCountryUsecase
}

// NewMockCountryUsecase creates a new mock instance
func NewMockCountryUsecase(ctrl *gomock.Controller) *MockCountryUsecase {
	mock := &MockCountryUsecase{ctrl: ctrl}
	mock.recorder = &MockCountryUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCountryUsecase) EXPECT() *MockCountryUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCountryUsecase) Create(country *models.Country) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", country)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCountryUsecaseMockRecorder) Create(country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCountryUsecase)(nil).Create), country)
}

// Update mocks base method
func (m *MockCountryUsecase) Update(newCountryData *models.Country) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", newCountryData)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCountryUsecaseMockRecorder) Update(newCountryData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCountryUsecase)(nil).Update), newCountryData)
}

// DeleteByID mocks base method
func (m *MockCountryUsecase) DeleteByID(countryID uint64) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", countryID)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockCountryUsecaseMockRecorder) DeleteByID(countryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockCountryUsecase)(nil).DeleteByID), countryID)
}

// GetByID mocks base method
func (m *MockCountryUsecase) GetByID(countryID uint64) (*models.Country, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", countryID)
	ret0, _ := ret[0].(*models.Country)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockCountryUsecaseMockRecorder) GetByID(countryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCountryUsecase)(nil).GetByID), countryID)
}

// GetByName mocks base method
func (m *MockCountryUsecase) GetByName(name string) (*models.Country, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", name)
	ret0, _ := ret[0].(*models.Country)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName
func (mr *MockCountryUsecaseMockRecorder) GetByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockCountryUsecase)(nil).GetByName), name)
}

// List mocks base method
func (m *MockCountryUsecase) List() ([]*models.Country, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*models.Country)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockCountryUsecaseMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCountryUsecase)(nil).List))
}

// ListByID mocks base method
func (m *MockCountryUsecase) ListByID(countriesID []uint64) ([]*models.Country, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByID", countriesID)
	ret0, _ := ret[0].([]*models.Country)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// ListByID indicates an expected call of ListByID
func (mr *MockCountryUsecaseMockRecorder) ListByID(countriesID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByID", reflect.TypeOf((*MockCountryUsecase)(nil).ListByID), countriesID)
}
