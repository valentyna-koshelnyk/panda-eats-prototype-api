// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entity "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
)

// CartService is an autogenerated mock type for the CartService type
type CartService struct {
	mock.Mock
}

// AddItem provides a mock function with given fields: userID, itemID, quantity
func (_m *CartService) AddItem(userID string, itemID string, quantity int64) error {
	ret := _m.Called(userID, itemID, quantity)

	if len(ret) == 0 {
		panic("no return value specified for AddItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(userID, itemID, quantity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetItemsList provides a mock function with given fields: userID
func (_m *CartService) GetItemsList(userID string) ([]entity.Cart, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetItemsList")
	}

	var r0 []entity.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]entity.Cart, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) []entity.Cart); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveItem provides a mock function with given fields: userID, itemID
func (_m *CartService) RemoveItem(userID string, itemID string) error {
	ret := _m.Called(userID, itemID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, itemID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveItems provides a mock function with given fields: cartItems
func (_m *CartService) RemoveItems(cartItems []entity.Cart) error {
	ret := _m.Called(cartItems)

	if len(ret) == 0 {
		panic("no return value specified for RemoveItems")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]entity.Cart) error); ok {
		r0 = rf(cartItems)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserItem provides a mock function with given fields: userID, itemID, quantity
func (_m *CartService) UpdateUserItem(userID string, itemID string, quantity int64) error {
	ret := _m.Called(userID, itemID, quantity)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(userID, itemID, quantity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCartService creates a new instance of CartService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCartService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CartService {
	mock := &CartService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
