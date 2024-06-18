// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entity "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"

	utils "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
)

// MenuService is an autogenerated mock type for the MenuService type
type MenuService struct {
	mock.Mock
}

// GetItem provides a mock function with given fields: itemID
func (_m *MenuService) GetItem(itemID string) (*entity.Menu, error) {
	ret := _m.Called(itemID)

	if len(ret) == 0 {
		panic("no return value specified for GetItem")
	}

	var r0 *entity.Menu
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Menu, error)); ok {
		return rf(itemID)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Menu); ok {
		r0 = rf(itemID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Menu)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(itemID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestaurantMenu provides a mock function with given fields: id, limit, offset
func (_m *MenuService) GetRestaurantMenu(id int, limit int, offset int) (*utils.Pagination, error) {
	ret := _m.Called(id, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetRestaurantMenu")
	}

	var r0 *utils.Pagination
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, int) (*utils.Pagination, error)); ok {
		return rf(id, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(int, int, int) *utils.Pagination); ok {
		r0 = rf(id, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.Pagination)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, int) error); ok {
		r1 = rf(id, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMenuService creates a new instance of MenuService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMenuService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MenuService {
	mock := &MenuService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
