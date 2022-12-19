// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	auth "capstone-alta1/features/auth"

	mock "github.com/stretchr/testify/mock"
)

// AuthRepository is an autogenerated mock type for the RepositoryInterface type
type AuthRepository struct {
	mock.Mock
}

// FindUser provides a mock function with given fields: email
func (_m *AuthRepository) FindUser(email string) (auth.Core, error) {
	ret := _m.Called(email)

	var r0 auth.Core
	if rf, ok := ret.Get(0).(func(string) auth.Core); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(auth.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthRepository creates a new instance of AuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthRepository(t mockConstructorTestingTNewAuthRepository) *AuthRepository {
	mock := &AuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
