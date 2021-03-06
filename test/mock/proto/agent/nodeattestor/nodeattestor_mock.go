// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/agent/nodeattestor (interfaces: NodeAttestor)

// Package mock_nodeattestor is a generated GoMock package.
package mock_nodeattestor

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	nodeattestor "github.com/spiffe/spire/proto/agent/nodeattestor"
	plugin "github.com/spiffe/spire/proto/common/plugin"
	reflect "reflect"
)

// MockNodeAttestor is a mock of NodeAttestor interface
type MockNodeAttestor struct {
	ctrl     *gomock.Controller
	recorder *MockNodeAttestorMockRecorder
}

// MockNodeAttestorMockRecorder is the mock recorder for MockNodeAttestor
type MockNodeAttestorMockRecorder struct {
	mock *MockNodeAttestor
}

// NewMockNodeAttestor creates a new mock instance
func NewMockNodeAttestor(ctrl *gomock.Controller) *MockNodeAttestor {
	mock := &MockNodeAttestor{ctrl: ctrl}
	mock.recorder = &MockNodeAttestorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeAttestor) EXPECT() *MockNodeAttestorMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockNodeAttestor) Configure(arg0 context.Context, arg1 *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockNodeAttestorMockRecorder) Configure(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockNodeAttestor)(nil).Configure), arg0, arg1)
}

// FetchAttestationData mocks base method
func (m *MockNodeAttestor) FetchAttestationData(arg0 context.Context, arg1 *nodeattestor.FetchAttestationDataRequest) (*nodeattestor.FetchAttestationDataResponse, error) {
	ret := m.ctrl.Call(m, "FetchAttestationData", arg0, arg1)
	ret0, _ := ret[0].(*nodeattestor.FetchAttestationDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAttestationData indicates an expected call of FetchAttestationData
func (mr *MockNodeAttestorMockRecorder) FetchAttestationData(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttestationData", reflect.TypeOf((*MockNodeAttestor)(nil).FetchAttestationData), arg0, arg1)
}

// GetPluginInfo mocks base method
func (m *MockNodeAttestor) GetPluginInfo(arg0 context.Context, arg1 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0, arg1)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockNodeAttestorMockRecorder) GetPluginInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockNodeAttestor)(nil).GetPluginInfo), arg0, arg1)
}
