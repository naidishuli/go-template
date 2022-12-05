// Code generated by MockGen. DO NOT EDIT.
// Source: context.go

// Package mocks is a generated GoMock package.
package mocks

import (
	logger "go-template/pkg/logger"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// AddValue mocks base method.
func (m *MockContext) AddValue(key string, value any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddValue", key, value)
}

// AddValue indicates an expected call of AddValue.
func (mr *MockContextMockRecorder) AddValue(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddValue", reflect.TypeOf((*MockContext)(nil).AddValue), key, value)
}

// ClearTransaction mocks base method.
func (m *MockContext) ClearTransaction() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearTransaction")
}

// ClearTransaction indicates an expected call of ClearTransaction.
func (mr *MockContextMockRecorder) ClearTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearTransaction", reflect.TypeOf((*MockContext)(nil).ClearTransaction))
}

// DB mocks base method.
func (m *MockContext) DB(def *gorm.DB) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB", def)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// DB indicates an expected call of DB.
func (mr *MockContextMockRecorder) DB(def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockContext)(nil).DB), def)
}

// GetValue mocks base method.
func (m *MockContext) GetValue(key string) any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue", key)
	ret0, _ := ret[0].(any)
	return ret0
}

// GetValue indicates an expected call of GetValue.
func (mr *MockContextMockRecorder) GetValue(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockContext)(nil).GetValue), key)
}

// Log mocks base method.
func (m *MockContext) Log(def logger.Logger) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log", def)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockContextMockRecorder) Log(def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockContext)(nil).Log), def)
}

// SetLogger mocks base method.
func (m *MockContext) SetLogger(log logger.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogger", log)
}

// SetLogger indicates an expected call of SetLogger.
func (mr *MockContextMockRecorder) SetLogger(log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockContext)(nil).SetLogger), log)
}

// SetTransaction mocks base method.
func (m *MockContext) SetTransaction(tx *gorm.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTransaction", tx)
}

// SetTransaction indicates an expected call of SetTransaction.
func (mr *MockContextMockRecorder) SetTransaction(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransaction", reflect.TypeOf((*MockContext)(nil).SetTransaction), tx)
}
