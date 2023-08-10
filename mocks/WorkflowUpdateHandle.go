// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by mockery v2.30.16. DO NOT EDIT.
// Modified manually to add new method.
package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// WorkflowUpdateHandle is an autogenerated mock type for the WorkflowUpdateHandle type
type WorkflowUpdateHandle struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, valuePtr
func (_m *WorkflowUpdateHandle) Get(ctx context.Context, valuePtr interface{}) error {
	ret := _m.Called(ctx, valuePtr)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, valuePtr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunID provides a mock function with given fields:
func (_m *WorkflowUpdateHandle) RunID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateID provides a mock function with given fields:
func (_m *WorkflowUpdateHandle) UpdateID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WorkflowID provides a mock function with given fields:
func (_m *WorkflowUpdateHandle) WorkflowID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewWorkflowUpdateHandle creates a new instance of WorkflowUpdateHandle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWorkflowUpdateHandle(t interface {
	mock.TestingT
	Cleanup(func())
}) *WorkflowUpdateHandle {
	mock := &WorkflowUpdateHandle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}