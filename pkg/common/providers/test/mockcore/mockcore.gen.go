// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core (interfaces: CryptoSuiteConfig,ConfigBackend,Providers)

// Package mockcore is a generated GoMock package.
package mockcore

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
)

// MockCryptoSuiteConfig is a mock of CryptoSuiteConfig interface
type MockCryptoSuiteConfig struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoSuiteConfigMockRecorder
}

// MockCryptoSuiteConfigMockRecorder is the mock recorder for MockCryptoSuiteConfig
type MockCryptoSuiteConfigMockRecorder struct {
	mock *MockCryptoSuiteConfig
}

// NewMockCryptoSuiteConfig creates a new mock instance
func NewMockCryptoSuiteConfig(ctrl *gomock.Controller) *MockCryptoSuiteConfig {
	mock := &MockCryptoSuiteConfig{ctrl: ctrl}
	mock.recorder = &MockCryptoSuiteConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCryptoSuiteConfig) EXPECT() *MockCryptoSuiteConfigMockRecorder {
	return m.recorder
}

// Ephemeral mocks base method
func (m *MockCryptoSuiteConfig) Ephemeral() bool {
	ret := m.ctrl.Call(m, "Ephemeral")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Ephemeral indicates an expected call of Ephemeral
func (mr *MockCryptoSuiteConfigMockRecorder) Ephemeral() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ephemeral", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).Ephemeral))
}

// IsSecurityEnabled mocks base method
func (m *MockCryptoSuiteConfig) IsSecurityEnabled() bool {
	ret := m.ctrl.Call(m, "IsSecurityEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSecurityEnabled indicates an expected call of IsSecurityEnabled
func (mr *MockCryptoSuiteConfigMockRecorder) IsSecurityEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSecurityEnabled", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).IsSecurityEnabled))
}

// KeyStorePath mocks base method
func (m *MockCryptoSuiteConfig) KeyStorePath() string {
	ret := m.ctrl.Call(m, "KeyStorePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// KeyStorePath indicates an expected call of KeyStorePath
func (mr *MockCryptoSuiteConfigMockRecorder) KeyStorePath() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyStorePath", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).KeyStorePath))
}

// SecurityAlgorithm mocks base method
func (m *MockCryptoSuiteConfig) SecurityAlgorithm() string {
	ret := m.ctrl.Call(m, "SecurityAlgorithm")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityAlgorithm indicates an expected call of SecurityAlgorithm
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityAlgorithm() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityAlgorithm", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityAlgorithm))
}

// SecurityLevel mocks base method
func (m *MockCryptoSuiteConfig) SecurityLevel() int {
	ret := m.ctrl.Call(m, "SecurityLevel")
	ret0, _ := ret[0].(int)
	return ret0
}

// SecurityLevel indicates an expected call of SecurityLevel
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityLevel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityLevel", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityLevel))
}

// SecurityProvider mocks base method
func (m *MockCryptoSuiteConfig) SecurityProvider() string {
	ret := m.ctrl.Call(m, "SecurityProvider")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProvider indicates an expected call of SecurityProvider
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProvider", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityProvider))
}

// SecurityProviderLabel mocks base method
func (m *MockCryptoSuiteConfig) SecurityProviderLabel() string {
	ret := m.ctrl.Call(m, "SecurityProviderLabel")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProviderLabel indicates an expected call of SecurityProviderLabel
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityProviderLabel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProviderLabel", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityProviderLabel))
}

// SecurityProviderLibPath mocks base method
func (m *MockCryptoSuiteConfig) SecurityProviderLibPath() string {
	ret := m.ctrl.Call(m, "SecurityProviderLibPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProviderLibPath indicates an expected call of SecurityProviderLibPath
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityProviderLibPath() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProviderLibPath", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityProviderLibPath))
}

// SecurityProviderPin mocks base method
func (m *MockCryptoSuiteConfig) SecurityProviderPin() string {
	ret := m.ctrl.Call(m, "SecurityProviderPin")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProviderPin indicates an expected call of SecurityProviderPin
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityProviderPin() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProviderPin", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityProviderPin))
}

// SoftVerify mocks base method
func (m *MockCryptoSuiteConfig) SoftVerify() bool {
	ret := m.ctrl.Call(m, "SoftVerify")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SoftVerify indicates an expected call of SoftVerify
func (mr *MockCryptoSuiteConfigMockRecorder) SoftVerify() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftVerify", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SoftVerify))
}

// MockConfigBackend is a mock of ConfigBackend interface
type MockConfigBackend struct {
	ctrl     *gomock.Controller
	recorder *MockConfigBackendMockRecorder
}

// MockConfigBackendMockRecorder is the mock recorder for MockConfigBackend
type MockConfigBackendMockRecorder struct {
	mock *MockConfigBackend
}

// NewMockConfigBackend creates a new mock instance
func NewMockConfigBackend(ctrl *gomock.Controller) *MockConfigBackend {
	mock := &MockConfigBackend{ctrl: ctrl}
	mock.recorder = &MockConfigBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigBackend) EXPECT() *MockConfigBackendMockRecorder {
	return m.recorder
}

// Lookup mocks base method
func (m *MockConfigBackend) Lookup(arg0 string, arg1 ...core.LookupOption) (interface{}, bool) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Lookup", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup
func (mr *MockConfigBackendMockRecorder) Lookup(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockConfigBackend)(nil).Lookup), varargs...)
}

// MockProviders is a mock of Providers interface
type MockProviders struct {
	ctrl     *gomock.Controller
	recorder *MockProvidersMockRecorder
}

// MockProvidersMockRecorder is the mock recorder for MockProviders
type MockProvidersMockRecorder struct {
	mock *MockProviders
}

// NewMockProviders creates a new mock instance
func NewMockProviders(ctrl *gomock.Controller) *MockProviders {
	mock := &MockProviders{ctrl: ctrl}
	mock.recorder = &MockProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProviders) EXPECT() *MockProvidersMockRecorder {
	return m.recorder
}

// CryptoSuite mocks base method
func (m *MockProviders) CryptoSuite() core.CryptoSuite {
	ret := m.ctrl.Call(m, "CryptoSuite")
	ret0, _ := ret[0].(core.CryptoSuite)
	return ret0
}

// CryptoSuite indicates an expected call of CryptoSuite
func (mr *MockProvidersMockRecorder) CryptoSuite() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CryptoSuite", reflect.TypeOf((*MockProviders)(nil).CryptoSuite))
}

// SigningManager mocks base method
func (m *MockProviders) SigningManager() core.SigningManager {
	ret := m.ctrl.Call(m, "SigningManager")
	ret0, _ := ret[0].(core.SigningManager)
	return ret0
}

// SigningManager indicates an expected call of SigningManager
func (mr *MockProvidersMockRecorder) SigningManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SigningManager", reflect.TypeOf((*MockProviders)(nil).SigningManager))
}
