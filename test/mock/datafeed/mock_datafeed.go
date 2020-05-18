// Code generated by MockGen. DO NOT EDIT.
// Source: internal/datafeed/datafeed.go

// Package mock_datafeed is a generated GoMock package.
package mock_datafeed

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockDataFeed is a mock of DataFeed interface
type MockDataFeed struct {
	ctrl     *gomock.Controller
	recorder *MockDataFeedMockRecorder
}

// MockDataFeedMockRecorder is the mock recorder for MockDataFeed
type MockDataFeedMockRecorder struct {
	mock *MockDataFeed
}

// NewMockDataFeed creates a new mock instance
func NewMockDataFeed(ctrl *gomock.Controller) *MockDataFeed {
	mock := &MockDataFeed{ctrl: ctrl}
	mock.recorder = &MockDataFeedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataFeed) EXPECT() *MockDataFeedMockRecorder {
	return m.recorder
}

// FindCurrentAssetPrice mocks base method
func (m *MockDataFeed) FindCurrentAssetPrice(assetID, currency string) (*float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCurrentAssetPrice", assetID, currency)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCurrentAssetPrice indicates an expected call of FindCurrentAssetPrice
func (mr *MockDataFeedMockRecorder) FindCurrentAssetPrice(assetID, currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCurrentAssetPrice", reflect.TypeOf((*MockDataFeed)(nil).FindCurrentAssetPrice), assetID, currency)
}

// FindPastAssetPrice mocks base method
func (m *MockDataFeed) FindPastAssetPrice(assetID, currency string, date time.Time) (*float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPastAssetPrice", assetID, currency, date)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPastAssetPrice indicates an expected call of FindPastAssetPrice
func (mr *MockDataFeedMockRecorder) FindPastAssetPrice(assetID, currency, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPastAssetPrice", reflect.TypeOf((*MockDataFeed)(nil).FindPastAssetPrice), assetID, currency, date)
}

// MockAssetPriceFeed is a mock of AssetPriceFeed interface
type MockAssetPriceFeed struct {
	ctrl     *gomock.Controller
	recorder *MockAssetPriceFeedMockRecorder
}

// MockAssetPriceFeedMockRecorder is the mock recorder for MockAssetPriceFeed
type MockAssetPriceFeedMockRecorder struct {
	mock *MockAssetPriceFeed
}

// NewMockAssetPriceFeed creates a new mock instance
func NewMockAssetPriceFeed(ctrl *gomock.Controller) *MockAssetPriceFeed {
	mock := &MockAssetPriceFeed{ctrl: ctrl}
	mock.recorder = &MockAssetPriceFeedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAssetPriceFeed) EXPECT() *MockAssetPriceFeedMockRecorder {
	return m.recorder
}

// FindCurrentAssetPrice mocks base method
func (m *MockAssetPriceFeed) FindCurrentAssetPrice(assetID, currency string) (*float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCurrentAssetPrice", assetID, currency)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCurrentAssetPrice indicates an expected call of FindCurrentAssetPrice
func (mr *MockAssetPriceFeedMockRecorder) FindCurrentAssetPrice(assetID, currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCurrentAssetPrice", reflect.TypeOf((*MockAssetPriceFeed)(nil).FindCurrentAssetPrice), assetID, currency)
}

// FindPastAssetPrice mocks base method
func (m *MockAssetPriceFeed) FindPastAssetPrice(assetID, currency string, date time.Time) (*float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPastAssetPrice", assetID, currency, date)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPastAssetPrice indicates an expected call of FindPastAssetPrice
func (mr *MockAssetPriceFeedMockRecorder) FindPastAssetPrice(assetID, currency, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPastAssetPrice", reflect.TypeOf((*MockAssetPriceFeed)(nil).FindPastAssetPrice), assetID, currency, date)
}