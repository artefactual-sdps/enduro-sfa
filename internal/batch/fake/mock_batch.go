// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-labs/enduro/internal/batch (interfaces: Service)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	batch "github.com/artefactual-labs/enduro/internal/api/gen/batch"
	package_ "github.com/artefactual-labs/enduro/internal/package_"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Hints mocks base method.
func (m *MockService) Hints(arg0 context.Context) (*batch.BatchHintsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hints", arg0)
	ret0, _ := ret[0].(*batch.BatchHintsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hints indicates an expected call of Hints.
func (mr *MockServiceMockRecorder) Hints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hints", reflect.TypeOf((*MockService)(nil).Hints), arg0)
}

// InitProcessingWorkflow mocks base method.
func (m *MockService) InitProcessingWorkflow(arg0 context.Context, arg1 *package_.ProcessingWorkflowRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitProcessingWorkflow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitProcessingWorkflow indicates an expected call of InitProcessingWorkflow.
func (mr *MockServiceMockRecorder) InitProcessingWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitProcessingWorkflow", reflect.TypeOf((*MockService)(nil).InitProcessingWorkflow), arg0, arg1)
}

// Status mocks base method.
func (m *MockService) Status(arg0 context.Context) (*batch.BatchStatusResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0)
	ret0, _ := ret[0].(*batch.BatchStatusResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockServiceMockRecorder) Status(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockService)(nil).Status), arg0)
}

// Submit mocks base method.
func (m *MockService) Submit(arg0 context.Context, arg1 *batch.SubmitPayload) (*batch.BatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Submit", arg0, arg1)
	ret0, _ := ret[0].(*batch.BatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Submit indicates an expected call of Submit.
func (mr *MockServiceMockRecorder) Submit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Submit", reflect.TypeOf((*MockService)(nil).Submit), arg0, arg1)
}
