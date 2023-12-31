// Code generated by MockGen. DO NOT EDIT.
// Source: receipt_interfaces.go

// Package receipt_test is a generated GoMock package.
package receipt_test

import (
	context "context"
	reflect "reflect"

	entity "github.com/adriansabvr/receipt_processor/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCaseContract is a mock of UseCaseContract interface.
type MockUseCaseContract struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseContractMockRecorder
}

// MockUseCaseContractMockRecorder is the mock recorder for MockUseCaseContract.
type MockUseCaseContractMockRecorder struct {
	mock *MockUseCaseContract
}

// NewMockUseCaseContract creates a new mock instance.
func NewMockUseCaseContract(ctrl *gomock.Controller) *MockUseCaseContract {
	mock := &MockUseCaseContract{ctrl: ctrl}
	mock.recorder = &MockUseCaseContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCaseContract) EXPECT() *MockUseCaseContractMockRecorder {
	return m.recorder
}

// GetPoints mocks base method.
func (m *MockUseCaseContract) GetPoints(arg0 context.Context, arg1 uint64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoints", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPoints indicates an expected call of GetPoints.
func (mr *MockUseCaseContractMockRecorder) GetPoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoints", reflect.TypeOf((*MockUseCaseContract)(nil).GetPoints), arg0, arg1)
}

// Process mocks base method.
func (m *MockUseCaseContract) Process(arg0 context.Context, arg1 entity.Receipt) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Process indicates an expected call of Process.
func (mr *MockUseCaseContractMockRecorder) Process(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockUseCaseContract)(nil).Process), arg0, arg1)
}

// MockRepoContract is a mock of RepoContract interface.
type MockRepoContract struct {
	ctrl     *gomock.Controller
	recorder *MockRepoContractMockRecorder
}

// MockRepoContractMockRecorder is the mock recorder for MockRepoContract.
type MockRepoContractMockRecorder struct {
	mock *MockRepoContract
}

// NewMockRepoContract creates a new mock instance.
func NewMockRepoContract(ctrl *gomock.Controller) *MockRepoContract {
	mock := &MockRepoContract{ctrl: ctrl}
	mock.recorder = &MockRepoContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoContract) EXPECT() *MockRepoContractMockRecorder {
	return m.recorder
}

// GetReceipt mocks base method.
func (m *MockRepoContract) GetReceipt(arg0 context.Context, arg1 uint64) (entity.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipt", arg0, arg1)
	ret0, _ := ret[0].(entity.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipt indicates an expected call of GetReceipt.
func (mr *MockRepoContractMockRecorder) GetReceipt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipt", reflect.TypeOf((*MockRepoContract)(nil).GetReceipt), arg0, arg1)
}

// InsertReceipt mocks base method.
func (m *MockRepoContract) InsertReceipt(arg0 context.Context, arg1 entity.Receipt) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertReceipt", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// InsertReceipt indicates an expected call of InsertReceipt.
func (mr *MockRepoContractMockRecorder) InsertReceipt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertReceipt", reflect.TypeOf((*MockRepoContract)(nil).InsertReceipt), arg0, arg1)
}
