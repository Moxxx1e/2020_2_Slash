// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_movie is a generated GoMock package.
package mocks

import (
	models "github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMovieRepository is a mock of MovieRepository interface
type MockMovieRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMovieRepositoryMockRecorder
}

// MockMovieRepositoryMockRecorder is the mock recorder for MockMovieRepository
type MockMovieRepositoryMockRecorder struct {
	mock *MockMovieRepository
}

// NewMockMovieRepository creates a new mock instance
func NewMockMovieRepository(ctrl *gomock.Controller) *MockMovieRepository {
	mock := &MockMovieRepository{ctrl: ctrl}
	mock.recorder = &MockMovieRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMovieRepository) EXPECT() *MockMovieRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockMovieRepository) Insert(movie *models.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockMovieRepositoryMockRecorder) Insert(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMovieRepository)(nil).Insert), movie)
}

// Update mocks base method
func (m *MockMovieRepository) Update(movie *models.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockMovieRepositoryMockRecorder) Update(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMovieRepository)(nil).Update), movie)
}

// DeleteByID mocks base method
func (m *MockMovieRepository) DeleteByID(movieID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", movieID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockMovieRepositoryMockRecorder) DeleteByID(movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockMovieRepository)(nil).DeleteByID), movieID)
}

// SelectByID mocks base method
func (m *MockMovieRepository) SelectByID(movieID uint64) (*models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByID", movieID)
	ret0, _ := ret[0].(*models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByID indicates an expected call of SelectByID
func (mr *MockMovieRepositoryMockRecorder) SelectByID(movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByID", reflect.TypeOf((*MockMovieRepository)(nil).SelectByID), movieID)
}

// SelectFullByID mocks base method
func (m *MockMovieRepository) SelectFullByID(movieID, curUserID uint64) (*models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectFullByID", movieID, curUserID)
	ret0, _ := ret[0].(*models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectFullByID indicates an expected call of SelectFullByID
func (mr *MockMovieRepositoryMockRecorder) SelectFullByID(movieID, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectFullByID", reflect.TypeOf((*MockMovieRepository)(nil).SelectFullByID), movieID, curUserID)
}

// SelectByContentID mocks base method
func (m *MockMovieRepository) SelectByContentID(contentID uint64) (*models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByContentID", contentID)
	ret0, _ := ret[0].(*models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByContentID indicates an expected call of SelectByContentID
func (mr *MockMovieRepositoryMockRecorder) SelectByContentID(contentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByContentID", reflect.TypeOf((*MockMovieRepository)(nil).SelectByContentID), contentID)
}

// SelectByParams mocks base method
func (m *MockMovieRepository) SelectByParams(params *models.ContentFilter, pgnt *models.Pagination, curUserID uint64) ([]*models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByParams", params, pgnt, curUserID)
	ret0, _ := ret[0].([]*models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByParams indicates an expected call of SelectByParams
func (mr *MockMovieRepositoryMockRecorder) SelectByParams(params, pgnt, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByParams", reflect.TypeOf((*MockMovieRepository)(nil).SelectByParams), params, pgnt, curUserID)
}

// SelectLatest mocks base method
func (m *MockMovieRepository) SelectLatest(pgnt *models.Pagination, curUserID uint64) ([]*models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectLatest", pgnt, curUserID)
	ret0, _ := ret[0].([]*models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectLatest indicates an expected call of SelectLatest
func (mr *MockMovieRepositoryMockRecorder) SelectLatest(pgnt, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectLatest", reflect.TypeOf((*MockMovieRepository)(nil).SelectLatest), pgnt, curUserID)
}

// SelectByRating mocks base method
func (m *MockMovieRepository) SelectByRating(pgnt *models.Pagination, curUserID uint64) ([]*models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByRating", pgnt, curUserID)
	ret0, _ := ret[0].([]*models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByRating indicates an expected call of SelectByRating
func (mr *MockMovieRepositoryMockRecorder) SelectByRating(pgnt, curUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByRating", reflect.TypeOf((*MockMovieRepository)(nil).SelectByRating), pgnt, curUserID)
}

// SelectWhereNameLike mocks base method
func (m *MockMovieRepository) SelectWhereNameLike(curUserID uint64, name string, limit, offset uint64) ([]*models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectWhereNameLike", curUserID, name, limit, offset)
	ret0, _ := ret[0].([]*models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectWhereNameLike indicates an expected call of SelectWhereNameLike
func (mr *MockMovieRepositoryMockRecorder) SelectWhereNameLike(curUserID, name, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectWhereNameLike", reflect.TypeOf((*MockMovieRepository)(nil).SelectWhereNameLike), curUserID, name, limit, offset)
}
