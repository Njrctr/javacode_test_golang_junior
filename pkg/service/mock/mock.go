// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
)

// MockAutorization is a mock of Autorization interface.
type MockAutorization struct {
	ctrl     *gomock.Controller
	recorder *MockAutorizationMockRecorder
}

// MockAutorizationMockRecorder is the mock recorder for MockAutorization.
type MockAutorizationMockRecorder struct {
	mock *MockAutorization
}

// NewMockAutorization creates a new mock instance.
func NewMockAutorization(ctrl *gomock.Controller) *MockAutorization {
	mock := &MockAutorization{ctrl: ctrl}
	mock.recorder = &MockAutorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutorization) EXPECT() *MockAutorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAutorization) CreateUser(user models.SignUpInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAutorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAutorization)(nil).CreateUser), user)
}

// GenerateJWTToken mocks base method.
func (m *MockAutorization) GenerateJWTToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateJWTToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateJWTToken indicates an expected call of GenerateJWTToken.
func (mr *MockAutorizationMockRecorder) GenerateJWTToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateJWTToken", reflect.TypeOf((*MockAutorization)(nil).GenerateJWTToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAutorization) ParseToken(token string) (int, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAutorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAutorization)(nil).ParseToken), token)
}

// MockWallet is a mock of Wallet interface.
type MockWallet struct {
	ctrl     *gomock.Controller
	recorder *MockWalletMockRecorder
}

// MockWalletMockRecorder is the mock recorder for MockWallet.
type MockWalletMockRecorder struct {
	mock *MockWallet
}

// NewMockWallet creates a new mock instance.
func NewMockWallet(ctrl *gomock.Controller) *MockWallet {
	mock := &MockWallet{ctrl: ctrl}
	mock.recorder = &MockWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWallet) EXPECT() *MockWalletMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWallet) Create(userId int) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWalletMockRecorder) Create(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWallet)(nil).Create), userId)
}

// Delete mocks base method.
func (m *MockWallet) Delete(userId int, walletUUID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, walletUUID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWalletMockRecorder) Delete(userId, walletUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWallet)(nil).Delete), userId, walletUUID)
}

// GetAll mocks base method.
func (m *MockWallet) GetAll(userId int) ([]models.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userId)
	ret0, _ := ret[0].([]models.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockWalletMockRecorder) GetAll(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockWallet)(nil).GetAll), userId)
}

// GetBalanceByUUID mocks base method.
func (m *MockWallet) GetBalanceByUUID(walletUUID uuid.UUID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceByUUID", walletUUID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalanceByUUID indicates an expected call of GetBalanceByUUID.
func (mr *MockWalletMockRecorder) GetBalanceByUUID(walletUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceByUUID", reflect.TypeOf((*MockWallet)(nil).GetBalanceByUUID), walletUUID)
}

// GetByUUID mocks base method.
func (m *MockWallet) GetByUUID(walletUUID uuid.UUID) (models.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUUID", walletUUID)
	ret0, _ := ret[0].(models.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUUID indicates an expected call of GetByUUID.
func (mr *MockWalletMockRecorder) GetByUUID(walletUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUUID", reflect.TypeOf((*MockWallet)(nil).GetByUUID), walletUUID)
}

// Update mocks base method.
func (m *MockWallet) Update(input models.WalletUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWalletMockRecorder) Update(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWallet)(nil).Update), input)
}

// MockAdmin is a mock of Admin interface.
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin.
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance.
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// BlockWallet mocks base method.
func (m *MockAdmin) BlockWallet(input models.BlockWallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockWallet", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// BlockWallet indicates an expected call of BlockWallet.
func (mr *MockAdminMockRecorder) BlockWallet(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockWallet", reflect.TypeOf((*MockAdmin)(nil).BlockWallet), input)
}

// Update mocks base method.
func (m *MockAdmin) Update(input models.WalletUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAdminMockRecorder) Update(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdmin)(nil).Update), input)
}
