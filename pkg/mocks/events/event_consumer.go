// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EventConsumer is an autogenerated mock type for the EventConsumer type
type EventConsumer struct {
	mock.Mock
}

// Consume provides a mock function with given fields: _a0, _a1
func (_m *EventConsumer) Consume(_a0 string, _a1 string) ([]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]byte, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEventConsumer interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventConsumer creates a new instance of EventConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventConsumer(t mockConstructorTestingTNewEventConsumer) *EventConsumer {
	mock := &EventConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}