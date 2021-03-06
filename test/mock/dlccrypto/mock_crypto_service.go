// Code generated by MockGen. DO NOT EDIT.
// Source: internal/dlccrypto/crypto_service.go

// Package mock_dlccrypto is a generated GoMock package.
package mock_dlccrypto

import (
	gomock "github.com/golang/mock/gomock"
	dlccrypto "p2pderivatives-oracle/internal/dlccrypto"
	reflect "reflect"
)

// MockCryptoService is a mock of CryptoService interface
type MockCryptoService struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoServiceMockRecorder
}

// MockCryptoServiceMockRecorder is the mock recorder for MockCryptoService
type MockCryptoServiceMockRecorder struct {
	mock *MockCryptoService
}

// NewMockCryptoService creates a new mock instance
func NewMockCryptoService(ctrl *gomock.Controller) *MockCryptoService {
	mock := &MockCryptoService{ctrl: ctrl}
	mock.recorder = &MockCryptoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCryptoService) EXPECT() *MockCryptoServiceMockRecorder {
	return m.recorder
}

// GenerateKvalue mocks base method
func (m *MockCryptoService) GenerateKvalue() (*dlccrypto.PrivateKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateKvalue")
	ret0, _ := ret[0].(*dlccrypto.PrivateKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateKvalue indicates an expected call of GenerateKvalue
func (mr *MockCryptoServiceMockRecorder) GenerateKvalue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateKvalue", reflect.TypeOf((*MockCryptoService)(nil).GenerateKvalue))
}

// ComputeRvalue mocks base method
func (m *MockCryptoService) ComputeRvalue(nonce *dlccrypto.PrivateKey) (*dlccrypto.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeRvalue", nonce)
	ret0, _ := ret[0].(*dlccrypto.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeRvalue indicates an expected call of ComputeRvalue
func (mr *MockCryptoServiceMockRecorder) ComputeRvalue(nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeRvalue", reflect.TypeOf((*MockCryptoService)(nil).ComputeRvalue), nonce)
}

// PublicKeyFromPrivateKey mocks base method
func (m *MockCryptoService) PublicKeyFromPrivateKey(privateKey *dlccrypto.PrivateKey) (*dlccrypto.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicKeyFromPrivateKey", privateKey)
	ret0, _ := ret[0].(*dlccrypto.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicKeyFromPrivateKey indicates an expected call of PublicKeyFromPrivateKey
func (mr *MockCryptoServiceMockRecorder) PublicKeyFromPrivateKey(privateKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicKeyFromPrivateKey", reflect.TypeOf((*MockCryptoService)(nil).PublicKeyFromPrivateKey), privateKey)
}

// ComputeSchnorrSignature mocks base method
func (m *MockCryptoService) ComputeSchnorrSignature(privateKey, oneTimeSigningK *dlccrypto.PrivateKey, message string) (*dlccrypto.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeSchnorrSignature", privateKey, oneTimeSigningK, message)
	ret0, _ := ret[0].(*dlccrypto.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeSchnorrSignature indicates an expected call of ComputeSchnorrSignature
func (mr *MockCryptoServiceMockRecorder) ComputeSchnorrSignature(privateKey, oneTimeSigningK, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeSchnorrSignature", reflect.TypeOf((*MockCryptoService)(nil).ComputeSchnorrSignature), privateKey, oneTimeSigningK, message)
}

// VerifySchnorrSignature mocks base method
func (m *MockCryptoService) VerifySchnorrSignature(publicKey, rvalue *dlccrypto.PublicKey, signature *dlccrypto.Signature, message string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySchnorrSignature", publicKey, rvalue, signature, message)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifySchnorrSignature indicates an expected call of VerifySchnorrSignature
func (mr *MockCryptoServiceMockRecorder) VerifySchnorrSignature(publicKey, rvalue, signature, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySchnorrSignature", reflect.TypeOf((*MockCryptoService)(nil).VerifySchnorrSignature), publicKey, rvalue, signature, message)
}
