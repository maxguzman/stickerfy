// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
	mock "github.com/stretchr/testify/mock"
)

// Metrics is an autogenerated mock type for the Metrics type
type Metrics struct {
	mock.Mock
}

// DestroyMetrics provides a mock function with given fields:
func (_m *Metrics) DestroyMetrics() {
	_m.Called()
}

// GetTimer provides a mock function with given fields: _a0, _a1
func (_m *Metrics) GetTimer(_a0 string, _a1 ...string) *prometheus.Timer {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *prometheus.Timer
	if rf, ok := ret.Get(0).(func(string, ...string) *prometheus.Timer); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*prometheus.Timer)
		}
	}

	return r0
}

// IncrementCounter provides a mock function with given fields: _a0, _a1
func (_m *Metrics) IncrementCounter(_a0 string, _a1 ...string) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// InitMetrics provides a mock function with given fields: _a0
func (_m *Metrics) InitMetrics(_a0 map[string]interface{}) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewMetrics interface {
	mock.TestingT
	Cleanup(func())
}

// NewMetrics creates a new instance of Metrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMetrics(t mockConstructorTestingTNewMetrics) *Metrics {
	mock := &Metrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
