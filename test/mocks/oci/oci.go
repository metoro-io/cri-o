// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cri-o/cri-o/internal/oci (interfaces: RuntimeImpl)

// Package ocimock is a generated GoMock package.
package ocimock

import (
	context "context"
	oci "github.com/cri-o/cri-o/internal/oci"
	gomock "github.com/golang/mock/gomock"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	io "io"
	remotecommand "k8s.io/client-go/tools/remotecommand"
	reflect "reflect"
	syscall "syscall"
)

// MockRuntimeImpl is a mock of RuntimeImpl interface
type MockRuntimeImpl struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeImplMockRecorder
}

// MockRuntimeImplMockRecorder is the mock recorder for MockRuntimeImpl
type MockRuntimeImplMockRecorder struct {
	mock *MockRuntimeImpl
}

// NewMockRuntimeImpl creates a new mock instance
func NewMockRuntimeImpl(ctrl *gomock.Controller) *MockRuntimeImpl {
	mock := &MockRuntimeImpl{ctrl: ctrl}
	mock.recorder = &MockRuntimeImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRuntimeImpl) EXPECT() *MockRuntimeImplMockRecorder {
	return m.recorder
}

// AttachContainer mocks base method
func (m *MockRuntimeImpl) AttachContainer(arg0 context.Context, arg1 *oci.Container, arg2 io.Reader, arg3, arg4 io.WriteCloser, arg5 bool, arg6 <-chan remotecommand.TerminalSize) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachContainer", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachContainer indicates an expected call of AttachContainer
func (mr *MockRuntimeImplMockRecorder) AttachContainer(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).AttachContainer), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ContainerStats mocks base method
func (m *MockRuntimeImpl) ContainerStats(arg0 context.Context, arg1 *oci.Container, arg2 string) (*oci.ContainerStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerStats", arg0, arg1, arg2)
	ret0, _ := ret[0].(*oci.ContainerStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerStats indicates an expected call of ContainerStats
func (mr *MockRuntimeImplMockRecorder) ContainerStats(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStats", reflect.TypeOf((*MockRuntimeImpl)(nil).ContainerStats), arg0, arg1, arg2)
}

// CreateContainer mocks base method
func (m *MockRuntimeImpl) CreateContainer(arg0 context.Context, arg1 *oci.Container, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateContainer indicates an expected call of CreateContainer
func (mr *MockRuntimeImplMockRecorder) CreateContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).CreateContainer), arg0, arg1, arg2)
}

// DeleteContainer mocks base method
func (m *MockRuntimeImpl) DeleteContainer(arg0 context.Context, arg1 *oci.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContainer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContainer indicates an expected call of DeleteContainer
func (mr *MockRuntimeImplMockRecorder) DeleteContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).DeleteContainer), arg0, arg1)
}

// ExecContainer mocks base method
func (m *MockRuntimeImpl) ExecContainer(arg0 context.Context, arg1 *oci.Container, arg2 []string, arg3 io.Reader, arg4, arg5 io.WriteCloser, arg6 bool, arg7 <-chan remotecommand.TerminalSize) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecContainer", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecContainer indicates an expected call of ExecContainer
func (mr *MockRuntimeImplMockRecorder) ExecContainer(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).ExecContainer), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// ExecSyncContainer mocks base method
func (m *MockRuntimeImpl) ExecSyncContainer(arg0 context.Context, arg1 *oci.Container, arg2 []string, arg3 int64) (*oci.ExecSyncResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecSyncContainer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*oci.ExecSyncResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecSyncContainer indicates an expected call of ExecSyncContainer
func (mr *MockRuntimeImplMockRecorder) ExecSyncContainer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecSyncContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).ExecSyncContainer), arg0, arg1, arg2, arg3)
}

// PauseContainer mocks base method
func (m *MockRuntimeImpl) PauseContainer(arg0 context.Context, arg1 *oci.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseContainer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PauseContainer indicates an expected call of PauseContainer
func (mr *MockRuntimeImplMockRecorder) PauseContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).PauseContainer), arg0, arg1)
}

// PortForwardContainer mocks base method
func (m *MockRuntimeImpl) PortForwardContainer(arg0 context.Context, arg1 *oci.Container, arg2 string, arg3 int32, arg4 io.ReadWriteCloser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PortForwardContainer", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// PortForwardContainer indicates an expected call of PortForwardContainer
func (mr *MockRuntimeImplMockRecorder) PortForwardContainer(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PortForwardContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).PortForwardContainer), arg0, arg1, arg2, arg3, arg4)
}

// ReopenContainerLog mocks base method
func (m *MockRuntimeImpl) ReopenContainerLog(arg0 context.Context, arg1 *oci.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReopenContainerLog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReopenContainerLog indicates an expected call of ReopenContainerLog
func (mr *MockRuntimeImplMockRecorder) ReopenContainerLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReopenContainerLog", reflect.TypeOf((*MockRuntimeImpl)(nil).ReopenContainerLog), arg0, arg1)
}

// SignalContainer mocks base method
func (m *MockRuntimeImpl) SignalContainer(arg0 context.Context, arg1 *oci.Container, arg2 syscall.Signal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignalContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignalContainer indicates an expected call of SignalContainer
func (mr *MockRuntimeImplMockRecorder) SignalContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).SignalContainer), arg0, arg1, arg2)
}

// StartContainer mocks base method
func (m *MockRuntimeImpl) StartContainer(arg0 context.Context, arg1 *oci.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartContainer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartContainer indicates an expected call of StartContainer
func (mr *MockRuntimeImplMockRecorder) StartContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).StartContainer), arg0, arg1)
}

// StopContainer mocks base method
func (m *MockRuntimeImpl) StopContainer(arg0 context.Context, arg1 *oci.Container, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopContainer indicates an expected call of StopContainer
func (mr *MockRuntimeImplMockRecorder) StopContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).StopContainer), arg0, arg1, arg2)
}

// UnpauseContainer mocks base method
func (m *MockRuntimeImpl) UnpauseContainer(arg0 context.Context, arg1 *oci.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpauseContainer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnpauseContainer indicates an expected call of UnpauseContainer
func (mr *MockRuntimeImplMockRecorder) UnpauseContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpauseContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).UnpauseContainer), arg0, arg1)
}

// UpdateContainer mocks base method
func (m *MockRuntimeImpl) UpdateContainer(arg0 context.Context, arg1 *oci.Container, arg2 *specs.LinuxResources) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContainer indicates an expected call of UpdateContainer
func (mr *MockRuntimeImplMockRecorder) UpdateContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContainer", reflect.TypeOf((*MockRuntimeImpl)(nil).UpdateContainer), arg0, arg1, arg2)
}

// UpdateContainerStatus mocks base method
func (m *MockRuntimeImpl) UpdateContainerStatus(arg0 context.Context, arg1 *oci.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContainerStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContainerStatus indicates an expected call of UpdateContainerStatus
func (mr *MockRuntimeImplMockRecorder) UpdateContainerStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContainerStatus", reflect.TypeOf((*MockRuntimeImpl)(nil).UpdateContainerStatus), arg0, arg1)
}

// WaitContainerStateStopped mocks base method
func (m *MockRuntimeImpl) WaitContainerStateStopped(arg0 context.Context, arg1 *oci.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitContainerStateStopped", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitContainerStateStopped indicates an expected call of WaitContainerStateStopped
func (mr *MockRuntimeImplMockRecorder) WaitContainerStateStopped(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitContainerStateStopped", reflect.TypeOf((*MockRuntimeImpl)(nil).WaitContainerStateStopped), arg0, arg1)
}
