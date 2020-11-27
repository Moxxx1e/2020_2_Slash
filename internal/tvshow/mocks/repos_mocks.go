// Code generated by MockGen. DO NOT EDIT.
// Source: internal/tvshow/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTVShowRepository is a mock of TVShowRepository interface
type MockTVShowRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTVShowRepositoryMockRecorder
}

// MockTVShowRepositoryMockRecorder is the mock recorder for MockTVShowRepository
type MockTVShowRepositoryMockRecorder struct {
	mock *MockTVShowRepository
}

// NewMockTVShowRepository creates a new mock instance
func NewMockTVShowRepository(ctrl *gomock.Controller) *MockTVShowRepository {
	mock := &MockTVShowRepository{ctrl: ctrl}
	mock.recorder = &MockTVShowRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTVShowRepository) EXPECT() *MockTVShowRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockTVShowRepository) Insert(tvshow *models.TVShow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", tvshow)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockTVShowRepositoryMockRecorder) Insert(tvshow interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTVShowRepository)(nil).Insert), tvshow)
}

// SelectByID mocks base method
func (m *MockTVShowRepository) SelectByID(tvshowID uint64) (*models.TVShow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByID", tvshowID)
	ret0, _ := ret[0].(*models.TVShow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByID indicates an expected call of SelectByID
func (mr *MockTVShowRepositoryMockRecorder) SelectByID(tvshowID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByID", reflect.TypeOf((*MockTVShowRepository)(nil).SelectByID), tvshowID)
}

// SelectFullByID mocks base method
func (m *MockTVShowRepository) SelectFullByID(tvshowID, curUserID uint64) (*models.TVShow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectFullByID", tvshowID, curUserID)
	ret0, _ := ret[0].(*models.TVShow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectFullByID indicates an expected call of SelectFullByID
func (mr *MockTVShowRepositoryMockRecorder) SelectFullByID(tvshowID, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectFullByID", reflect.TypeOf((*MockTVShowRepository)(nil).SelectFullByID), tvshowID, curUserID)
}

// SelectByContentID mocks base method
func (m *MockTVShowRepository) SelectByContentID(contentID uint64) (*models.TVShow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByContentID", contentID)
	ret0, _ := ret[0].(*models.TVShow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByContentID indicates an expected call of SelectByContentID
func (mr *MockTVShowRepositoryMockRecorder) SelectByContentID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByContentID", reflect.TypeOf((*MockTVShowRepository)(nil).SelectByContentID), contentID)
}

// SelectWhereNameLike mocks base method
func (m *MockTVShowRepository) SelectWhereNameLike(name string, pgnt *models.Pagination, curUserID uint64) ([]*models.TVShow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectWhereNameLike", name, pgnt, curUserID)
	ret0, _ := ret[0].([]*models.TVShow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectWhereNameLike indicates an expected call of SelectWhereNameLike
func (mr *MockTVShowRepositoryMockRecorder) SelectWhereNameLike(name, pgnt, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectWhereNameLike", reflect.TypeOf((*MockTVShowRepository)(nil).SelectWhereNameLike), name, pgnt, curUserID)
}

// SelectByParams mocks base method
func (m *MockTVShowRepository) SelectByParams(params *models.ContentFilter, pgnt *models.Pagination, curUserID uint64) ([]*models.TVShow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByParams", params, pgnt, curUserID)
	ret0, _ := ret[0].([]*models.TVShow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByParams indicates an expected call of SelectByParams
func (mr *MockTVShowRepositoryMockRecorder) SelectByParams(params, pgnt, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByParams", reflect.TypeOf((*MockTVShowRepository)(nil).SelectByParams), params, pgnt, curUserID)
}

// SelectLatest mocks base method
func (m *MockTVShowRepository) SelectLatest(pgnt *models.Pagination, curUserID uint64) ([]*models.TVShow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectLatest", pgnt, curUserID)
	ret0, _ := ret[0].([]*models.TVShow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectLatest indicates an expected call of SelectLatest
func (mr *MockTVShowRepositoryMockRecorder) SelectLatest(pgnt, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectLatest", reflect.TypeOf((*MockTVShowRepository)(nil).SelectLatest), pgnt, curUserID)
}

// SelectByRating mocks base method
func (m *MockTVShowRepository) SelectByRating(pgnt *models.Pagination, curUserID uint64) ([]*models.TVShow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByRating", pgnt, curUserID)
	ret0, _ := ret[0].([]*models.TVShow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByRating indicates an expected call of SelectByRating
func (mr *MockTVShowRepositoryMockRecorder) SelectByRating(pgnt, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByRating", reflect.TypeOf((*MockTVShowRepository)(nil).SelectByRating), pgnt, curUserID)
}
