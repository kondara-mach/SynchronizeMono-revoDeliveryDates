// Code generated by MockGen. DO NOT EDIT.
// Source: .\difference.go

// Package mock_compare is a generated GoMock package.
package mock_compare

import (
	compare "SynchronizeMonorevoDeliveryDates/domain/compare"
	monorevo "SynchronizeMonorevoDeliveryDates/domain/monorevo"
	orderdb "SynchronizeMonorevoDeliveryDates/domain/orderdb"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExtractor is a mock of Extractor interface.
type MockExtractor struct {
	ctrl     *gomock.Controller
	recorder *MockExtractorMockRecorder
}

// MockExtractorMockRecorder is the mock recorder for MockExtractor.
type MockExtractorMockRecorder struct {
	mock *MockExtractor
}

// NewMockExtractor creates a new mock instance.
func NewMockExtractor(ctrl *gomock.Controller) *MockExtractor {
	mock := &MockExtractor{ctrl: ctrl}
	mock.recorder = &MockExtractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtractor) EXPECT() *MockExtractorMockRecorder {
	return m.recorder
}

// ExtractForDeliveryDate mocks base method.
func (m *MockExtractor) ExtractForDeliveryDate(j []orderdb.JobBook, p []monorevo.Proposition) []compare.DifferentProposition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractForDeliveryDate", j, p)
	ret0, _ := ret[0].([]compare.DifferentProposition)
	return ret0
}

// ExtractForDeliveryDate indicates an expected call of ExtractForDeliveryDate.
func (mr *MockExtractorMockRecorder) ExtractForDeliveryDate(j, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractForDeliveryDate", reflect.TypeOf((*MockExtractor)(nil).ExtractForDeliveryDate), j, p)
}
