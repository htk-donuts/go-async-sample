// Code generated by MockGen. DO NOT EDIT.
// Source: csv_presenter.go
//
// Generated by this command:
//
//	mockgen -source=csv_presenter.go -destination=mock/csv_presenter.go
//

// Package mock_presenter is a generated GoMock package.
package mock_presenter

import (
	reflect "reflect"

	model "github.com/htk-donuts/go-async-sample/internal/domain/model"
	gomock "go.uber.org/mock/gomock"
)

// MockCSVPresenter is a mock of CSVPresenter interface.
type MockCSVPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockCSVPresenterMockRecorder
	isgomock struct{}
}

// MockCSVPresenterMockRecorder is the mock recorder for MockCSVPresenter.
type MockCSVPresenterMockRecorder struct {
	mock *MockCSVPresenter
}

// NewMockCSVPresenter creates a new mock instance.
func NewMockCSVPresenter(ctrl *gomock.Controller) *MockCSVPresenter {
	mock := &MockCSVPresenter{ctrl: ctrl}
	mock.recorder = &MockCSVPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCSVPresenter) EXPECT() *MockCSVPresenterMockRecorder {
	return m.recorder
}

// OutputCSV mocks base method.
func (m *MockCSVPresenter) OutputCSV(arg0 []model.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutputCSV", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// OutputCSV indicates an expected call of OutputCSV.
func (mr *MockCSVPresenterMockRecorder) OutputCSV(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutputCSV", reflect.TypeOf((*MockCSVPresenter)(nil).OutputCSV), arg0)
}
