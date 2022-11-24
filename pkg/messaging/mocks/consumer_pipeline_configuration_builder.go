// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	pipeline "github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/pipeline"
	mock "github.com/stretchr/testify/mock"
)

// ConsumerPipelineConfigurationBuilder is an autogenerated mock type for the ConsumerPipelineConfigurationBuilder type
type ConsumerPipelineConfigurationBuilder struct {
	mock.Mock
}

// AddPipeline provides a mock function with given fields: _a0
func (_m *ConsumerPipelineConfigurationBuilder) AddPipeline(_a0 pipeline.ConsumerPipeline) pipeline.ConsumerPipelineConfigurationBuilder {
	ret := _m.Called(_a0)

	var r0 pipeline.ConsumerPipelineConfigurationBuilder
	if rf, ok := ret.Get(0).(func(pipeline.ConsumerPipeline) pipeline.ConsumerPipelineConfigurationBuilder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pipeline.ConsumerPipelineConfigurationBuilder)
		}
	}

	return r0
}

// Build provides a mock function with given fields:
func (_m *ConsumerPipelineConfigurationBuilder) Build() *pipeline.ConsumerPipelineConfiguration {
	ret := _m.Called()

	var r0 *pipeline.ConsumerPipelineConfiguration
	if rf, ok := ret.Get(0).(func() *pipeline.ConsumerPipelineConfiguration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipeline.ConsumerPipelineConfiguration)
		}
	}

	return r0
}

type mockConstructorTestingTNewConsumerPipelineConfigurationBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumerPipelineConfigurationBuilder creates a new instance of ConsumerPipelineConfigurationBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumerPipelineConfigurationBuilder(t mockConstructorTestingTNewConsumerPipelineConfigurationBuilder) *ConsumerPipelineConfigurationBuilder {
	mock := &ConsumerPipelineConfigurationBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}