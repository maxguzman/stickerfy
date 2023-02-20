// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "stickerfy/app/models"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ProductService is an autogenerated mock type for the ProductService type
type ProductService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: product
func (_m *ProductService) Delete(product models.Product) error {
	ret := _m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ProductService) GetAll() ([]models.Product, error) {
	ret := _m.Called()

	var r0 []models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Product, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *ProductService) GetByID(id uuid.UUID) (models.Product, error) {
	ret := _m.Called(id)

	var r0 models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (models.Product, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) models.Product); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: product
func (_m *ProductService) Post(product models.Product) error {
	ret := _m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: product
func (_m *ProductService) Update(product models.Product) error {
	ret := _m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewProductService interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductService creates a new instance of ProductService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductService(t mockConstructorTestingTNewProductService) *ProductService {
	mock := &ProductService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}