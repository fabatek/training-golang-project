// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"
	request "faba_traning_project/internal/httpbody/request"

	mock "github.com/stretchr/testify/mock"

	response "faba_traning_project/internal/httpbody/response"
)

// IOrder is an autogenerated mock type for the IOrder type
type IOrder struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, requestUser
func (_m *IOrder) Create(ctx context.Context, requestUser request.CreateOrder) (response.Order, error) {
	ret := _m.Called(ctx, requestUser)

	var r0 response.Order
	if rf, ok := ret.Get(0).(func(context.Context, request.CreateOrder) response.Order); ok {
		r0 = rf(ctx, requestUser)
	} else {
		r0 = ret.Get(0).(response.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, request.CreateOrder) error); ok {
		r1 = rf(ctx, requestUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewIOrderT interface {
	mock.TestingT
	Cleanup(func())
}

// NewIOrder creates a new instance of IOrder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIOrder(t NewIOrderT) *IOrder {
	mock := &IOrder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
