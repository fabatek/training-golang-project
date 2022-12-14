// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"
	models "faba_traning_project/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *Repository) Create(ctx context.Context, user models.Product) (models.Product, error) {
	ret := _m.Called(ctx, user)

	var r0 models.Product
	if rf, ok := ret.Get(0).(func(context.Context, models.Product) models.Product); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Product) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListProductByListID provides a mock function with given fields: ctx, productIDs
func (_m *Repository) GetListProductByListID(ctx context.Context, productIDs []string) ([]models.Product, error) {
	ret := _m.Called(ctx, productIDs)

	var r0 []models.Product
	if rf, ok := ret.Get(0).(func(context.Context, []string) []models.Product); ok {
		r0 = rf(ctx, productIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, productIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductById provides a mock function with given fields: ctx, productID
func (_m *Repository) GetProductById(ctx context.Context, productID string) (models.Product, error) {
	ret := _m.Called(ctx, productID)

	var r0 models.Product
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Product); ok {
		r0 = rf(ctx, productID)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t NewRepositoryT) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
