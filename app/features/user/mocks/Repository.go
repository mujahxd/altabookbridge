// Code generated by mockery v2.23.2. DO NOT EDIT.

package mocks

import (
	user "github.com/mujahxd/altabookbridge/app/features/user"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FindByUsername provides a mock function with given fields: username
func (_m *Repository) FindByUsername(username string) (user.Core, error) {
	ret := _m.Called(username)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.Core, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) user.Core); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: _a0
func (_m *Repository) Save(_a0 user.Core) (user.Core, error) {
	ret := _m.Called(_a0)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(user.Core) (user.Core, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(user.Core) user.Core); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(user.Core) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
