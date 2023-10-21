// Code generated by MockGen. DO NOT EDIT.
// Source: models/person_store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models  "institute-person-api/src/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPersonStoreInterface is a mock of PersonStoreInterface interface.
type MockPersonStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPersonStoreInterfaceMockRecorder
}

// MockPersonStoreInterfaceMockRecorder is the mock recorder for MockPersonStoreInterface.
type MockPersonStoreInterfaceMockRecorder struct {
	mock *MockPersonStoreInterface
}

// NewMockPersonStoreInterface creates a new mock instance.
func NewMockPersonStoreInterface(ctrl *gomock.Controller) *MockPersonStoreInterface {
	mock := &MockPersonStoreInterface{ctrl: ctrl}
	mock.recorder = &MockPersonStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonStoreInterface) EXPECT() *MockPersonStoreInterfaceMockRecorder {
	return m.recorder
}

// Disconnect mocks base method.
func (m *MockPersonStoreInterface) Disconnect() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disconnect")
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockPersonStoreInterfaceMockRecorder) Disconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockPersonStoreInterface)(nil).Disconnect))
}

// FindMany mocks base method.
func (m *MockPersonStoreInterface) FindMany() ([]models.PersonShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMany")
	ret0, _ := ret[0].([]models.PersonShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMany indicates an expected call of FindMany.
func (mr *MockPersonStoreInterfaceMockRecorder) FindMany() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMany", reflect.TypeOf((*MockPersonStoreInterface)(nil).FindMany))
}

// FindOne mocks base method.
func (m *MockPersonStoreInterface) FindOne(id string) (*models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", id)
	ret0, _ := ret[0].(*models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockPersonStoreInterfaceMockRecorder) FindOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockPersonStoreInterface)(nil).FindOne), id)
}

// FindOneAndUpdate mocks base method.
func (m *MockPersonStoreInterface) FindOneAndUpdate(id string, information []byte, crumb *models.BreadCrumb) (*models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneAndUpdate", id, information, crumb)
	ret0, _ := ret[0].(*models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneAndUpdate indicates an expected call of FindOneAndUpdate.
func (mr *MockPersonStoreInterfaceMockRecorder) FindOneAndUpdate(id, information, crumb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneAndUpdate", reflect.TypeOf((*MockPersonStoreInterface)(nil).FindOneAndUpdate), id, information, crumb)
}

// Insert mocks base method.
func (m *MockPersonStoreInterface) Insert(information []byte, crumb *models.BreadCrumb) (*models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", information, crumb)
	ret0, _ := ret[0].(*models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockPersonStoreInterfaceMockRecorder) Insert(information, crumb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPersonStoreInterface)(nil).Insert), information, crumb)
}
