// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	data "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogs/write_service/internal/products/contracts/data"
	mock "github.com/stretchr/testify/mock"
)

// CatalogUnitOfWork is an autogenerated mock type for the CatalogUnitOfWork type
type CatalogUnitOfWork struct {
	mock.Mock
}

// Do provides a mock function with given fields: ctx, action
func (_m *CatalogUnitOfWork) Do(ctx context.Context, action data.CatalogUnitOfWorkActionFunc) error {
	ret := _m.Called(ctx, action)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, data.CatalogUnitOfWorkActionFunc) error); ok {
		r0 = rf(ctx, action)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCatalogUnitOfWork interface {
	mock.TestingT
	Cleanup(func())
}

// NewCatalogUnitOfWork creates a new instance of CatalogUnitOfWork. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCatalogUnitOfWork(t mockConstructorTestingTNewCatalogUnitOfWork) *CatalogUnitOfWork {
	mock := &CatalogUnitOfWork{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
