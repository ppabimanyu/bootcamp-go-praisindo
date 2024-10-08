// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	domain "boiler-plate/internal/users/domain"
	exception "boiler-plate/pkg/exception"
	context "context"

	mock "github.com/stretchr/testify/mock"

	service "boiler-plate/internal/users/service"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Auth provides a mock function with given fields: ctx, email, password
func (_m *Service) Auth(ctx context.Context, email string, password string) (*domain.Users, *exception.Exception) {
	ret := _m.Called(ctx, email, password)

	var r0 *domain.Users
	var r1 *exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*domain.Users, *exception.Exception)); ok {
		return rf(ctx, email, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.Users); ok {
		r0 = rf(ctx, email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Users)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) *exception.Exception); ok {
		r1 = rf(ctx, email, password)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*exception.Exception)
		}
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, req
func (_m *Service) Create(ctx context.Context, req *domain.Users) *exception.Exception {
	ret := _m.Called(ctx, req)

	var r0 *exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Users) *exception.Exception); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*exception.Exception)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Service) Delete(ctx context.Context, id string) *exception.Exception {
	ret := _m.Called(ctx, id)

	var r0 *exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, string) *exception.Exception); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*exception.Exception)
		}
	}

	return r0
}

// Detail provides a mock function with given fields: ctx, id
func (_m *Service) Detail(ctx context.Context, id string) (*domain.UserResponse, *exception.Exception) {
	ret := _m.Called(ctx, id)

	var r0 *domain.UserResponse
	var r1 *exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.UserResponse, *exception.Exception)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.UserResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *exception.Exception); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*exception.Exception)
		}
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, limit, page
func (_m *Service) Find(ctx context.Context, limit string, page string) (*service.FindResponse, *exception.Exception) {
	ret := _m.Called(ctx, limit, page)

	var r0 *service.FindResponse
	var r1 *exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*service.FindResponse, *exception.Exception)); ok {
		return rf(ctx, limit, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *service.FindResponse); ok {
		r0 = rf(ctx, limit, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.FindResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) *exception.Exception); ok {
		r1 = rf(ctx, limit, page)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*exception.Exception)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, users
func (_m *Service) Update(ctx context.Context, id string, users *domain.Users) *exception.Exception {
	ret := _m.Called(ctx, id, users)

	var r0 *exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.Users) *exception.Exception); ok {
		r0 = rf(ctx, id, users)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*exception.Exception)
		}
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
