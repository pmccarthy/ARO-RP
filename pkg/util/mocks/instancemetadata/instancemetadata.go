// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/instancemetadata (interfaces: ServicePrincipalToken)

// Package mock_instancemetadata is a generated GoMock package.
package mock_instancemetadata

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServicePrincipalToken is a mock of ServicePrincipalToken interface
type MockServicePrincipalToken struct {
	ctrl     *gomock.Controller
	recorder *MockServicePrincipalTokenMockRecorder
}

// MockServicePrincipalTokenMockRecorder is the mock recorder for MockServicePrincipalToken
type MockServicePrincipalTokenMockRecorder struct {
	mock *MockServicePrincipalToken
}

// NewMockServicePrincipalToken creates a new mock instance
func NewMockServicePrincipalToken(ctrl *gomock.Controller) *MockServicePrincipalToken {
	mock := &MockServicePrincipalToken{ctrl: ctrl}
	mock.recorder = &MockServicePrincipalTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServicePrincipalToken) EXPECT() *MockServicePrincipalTokenMockRecorder {
	return m.recorder
}

// EnsureFresh mocks base method
func (m *MockServicePrincipalToken) EnsureFresh() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureFresh")
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureFresh indicates an expected call of EnsureFresh
func (mr *MockServicePrincipalTokenMockRecorder) EnsureFresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureFresh", reflect.TypeOf((*MockServicePrincipalToken)(nil).EnsureFresh))
}

// OAuthToken mocks base method
func (m *MockServicePrincipalToken) OAuthToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OAuthToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// OAuthToken indicates an expected call of OAuthToken
func (mr *MockServicePrincipalTokenMockRecorder) OAuthToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuthToken", reflect.TypeOf((*MockServicePrincipalToken)(nil).OAuthToken))
}
