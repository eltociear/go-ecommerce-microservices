// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	metadata "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/core/metadata"
	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/types"
)

// MessageConsumeContext is an autogenerated mock type for the MessageConsumeContext type
type MessageConsumeContext struct {
	mock.Mock
}

// ContentType provides a mock function with given fields:
func (_m *MessageConsumeContext) ContentType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// CorrelationId provides a mock function with given fields:
func (_m *MessageConsumeContext) CorrelationId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Created provides a mock function with given fields:
func (_m *MessageConsumeContext) Created() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// DeliveryTag provides a mock function with given fields:
func (_m *MessageConsumeContext) DeliveryTag() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// Message provides a mock function with given fields:
func (_m *MessageConsumeContext) Message() types.IMessage {
	ret := _m.Called()

	var r0 types.IMessage
	if rf, ok := ret.Get(0).(func() types.IMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.IMessage)
		}
	}

	return r0
}

// MessageId provides a mock function with given fields:
func (_m *MessageConsumeContext) MessageId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MessageType provides a mock function with given fields:
func (_m *MessageConsumeContext) MessageType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Metadata provides a mock function with given fields:
func (_m *MessageConsumeContext) Metadata() metadata.Metadata {
	ret := _m.Called()

	var r0 metadata.Metadata
	if rf, ok := ret.Get(0).(func() metadata.Metadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.Metadata)
		}
	}

	return r0
}

type mockConstructorTestingTNewMessageConsumeContext interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessageConsumeContext creates a new instance of MessageConsumeContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageConsumeContext(t mockConstructorTestingTNewMessageConsumeContext) *MessageConsumeContext {
	mock := &MessageConsumeContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
