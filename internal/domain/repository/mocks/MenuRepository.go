// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	utils "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
)

// MenuRepository is an autogenerated mock type for the MenuRepository type
type MenuRepository struct {
	mock.Mock
}

// GetMenu provides a mock function with given fields: id, pagination
func (_m *MenuRepository) GetMenu(id int, pagination *utils.Pagination) (*utils.Pagination, error) {
	ret := _m.Called(id, pagination)

	if len(ret) == 0 {
		panic("no return value specified for GetMenu")
	}

	var r0 *utils.Pagination
	var r1 error
	if rf, ok := ret.Get(0).(func(int, *utils.Pagination) (*utils.Pagination, error)); ok {
		return rf(id, pagination)
	}
	if rf, ok := ret.Get(0).(func(int, *utils.Pagination) *utils.Pagination); ok {
		r0 = rf(id, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.Pagination)
		}
	}

	if rf, ok := ret.Get(1).(func(int, *utils.Pagination) error); ok {
		r1 = rf(id, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMenuRepository creates a new instance of MenuRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMenuRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MenuRepository {
	mock := &MenuRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
