// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// Hash provides a mock function with given fields: s
func (_m *AuthService) Hash(s string) (string, error) {
	ret := _m.Called(s)

	if len(ret) == 0 {
		panic("no return value specified for Hash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(s)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyPassword provides a mock function with given fields: userPassword, providedPassword
func (_m *AuthService) VerifyPassword(userPassword string, providedPassword string) bool {
	ret := _m.Called(userPassword, providedPassword)

	if len(ret) == 0 {
		panic("no return value specified for VerifyPassword")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(userPassword, providedPassword)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
