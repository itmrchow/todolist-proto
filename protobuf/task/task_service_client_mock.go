// Code generated by mockery v2.53.2. DO NOT EDIT.

package task

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockTaskServiceClient is an autogenerated mock type for the TaskServiceClient type
type MockTaskServiceClient struct {
	mock.Mock
}

type MockTaskServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTaskServiceClient) EXPECT() *MockTaskServiceClient_Expecter {
	return &MockTaskServiceClient_Expecter{mock: &_m.Mock}
}

// CreateTask provides a mock function with given fields: ctx, in, opts
func (_m *MockTaskServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 *CreateTaskResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *CreateTaskRequest, ...grpc.CallOption) (*CreateTaskResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *CreateTaskRequest, ...grpc.CallOption) *CreateTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreateTaskResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *CreateTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTaskServiceClient_CreateTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTask'
type MockTaskServiceClient_CreateTask_Call struct {
	*mock.Call
}

// CreateTask is a helper method to define mock.On call
//   - ctx context.Context
//   - in *CreateTaskRequest
//   - opts ...grpc.CallOption
func (_e *MockTaskServiceClient_Expecter) CreateTask(ctx interface{}, in interface{}, opts ...interface{}) *MockTaskServiceClient_CreateTask_Call {
	return &MockTaskServiceClient_CreateTask_Call{Call: _e.mock.On("CreateTask",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTaskServiceClient_CreateTask_Call) Run(run func(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption)) *MockTaskServiceClient_CreateTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*CreateTaskRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTaskServiceClient_CreateTask_Call) Return(_a0 *CreateTaskResponse, _a1 error) *MockTaskServiceClient_CreateTask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTaskServiceClient_CreateTask_Call) RunAndReturn(run func(context.Context, *CreateTaskRequest, ...grpc.CallOption) (*CreateTaskResponse, error)) *MockTaskServiceClient_CreateTask_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTask provides a mock function with given fields: ctx, in, opts
func (_m *MockTaskServiceClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 *DeleteTaskResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteTaskRequest, ...grpc.CallOption) (*DeleteTaskResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteTaskRequest, ...grpc.CallOption) *DeleteTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeleteTaskResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *DeleteTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTaskServiceClient_DeleteTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTask'
type MockTaskServiceClient_DeleteTask_Call struct {
	*mock.Call
}

// DeleteTask is a helper method to define mock.On call
//   - ctx context.Context
//   - in *DeleteTaskRequest
//   - opts ...grpc.CallOption
func (_e *MockTaskServiceClient_Expecter) DeleteTask(ctx interface{}, in interface{}, opts ...interface{}) *MockTaskServiceClient_DeleteTask_Call {
	return &MockTaskServiceClient_DeleteTask_Call{Call: _e.mock.On("DeleteTask",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTaskServiceClient_DeleteTask_Call) Run(run func(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption)) *MockTaskServiceClient_DeleteTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*DeleteTaskRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTaskServiceClient_DeleteTask_Call) Return(_a0 *DeleteTaskResponse, _a1 error) *MockTaskServiceClient_DeleteTask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTaskServiceClient_DeleteTask_Call) RunAndReturn(run func(context.Context, *DeleteTaskRequest, ...grpc.CallOption) (*DeleteTaskResponse, error)) *MockTaskServiceClient_DeleteTask_Call {
	_c.Call.Return(run)
	return _c
}

// FindTask provides a mock function with given fields: ctx, in, opts
func (_m *MockTaskServiceClient) FindTask(ctx context.Context, in *FindTaskRequest, opts ...grpc.CallOption) (*FindTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindTask")
	}

	var r0 *FindTaskResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *FindTaskRequest, ...grpc.CallOption) (*FindTaskResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *FindTaskRequest, ...grpc.CallOption) *FindTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*FindTaskResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *FindTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTaskServiceClient_FindTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindTask'
type MockTaskServiceClient_FindTask_Call struct {
	*mock.Call
}

// FindTask is a helper method to define mock.On call
//   - ctx context.Context
//   - in *FindTaskRequest
//   - opts ...grpc.CallOption
func (_e *MockTaskServiceClient_Expecter) FindTask(ctx interface{}, in interface{}, opts ...interface{}) *MockTaskServiceClient_FindTask_Call {
	return &MockTaskServiceClient_FindTask_Call{Call: _e.mock.On("FindTask",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTaskServiceClient_FindTask_Call) Run(run func(ctx context.Context, in *FindTaskRequest, opts ...grpc.CallOption)) *MockTaskServiceClient_FindTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*FindTaskRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTaskServiceClient_FindTask_Call) Return(_a0 *FindTaskResponse, _a1 error) *MockTaskServiceClient_FindTask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTaskServiceClient_FindTask_Call) RunAndReturn(run func(context.Context, *FindTaskRequest, ...grpc.CallOption) (*FindTaskResponse, error)) *MockTaskServiceClient_FindTask_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTask provides a mock function with given fields: ctx, in, opts
func (_m *MockTaskServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 *UpdateTaskResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateTaskRequest, ...grpc.CallOption) (*UpdateTaskResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateTaskRequest, ...grpc.CallOption) *UpdateTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateTaskResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *UpdateTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTaskServiceClient_UpdateTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTask'
type MockTaskServiceClient_UpdateTask_Call struct {
	*mock.Call
}

// UpdateTask is a helper method to define mock.On call
//   - ctx context.Context
//   - in *UpdateTaskRequest
//   - opts ...grpc.CallOption
func (_e *MockTaskServiceClient_Expecter) UpdateTask(ctx interface{}, in interface{}, opts ...interface{}) *MockTaskServiceClient_UpdateTask_Call {
	return &MockTaskServiceClient_UpdateTask_Call{Call: _e.mock.On("UpdateTask",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTaskServiceClient_UpdateTask_Call) Run(run func(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption)) *MockTaskServiceClient_UpdateTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*UpdateTaskRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTaskServiceClient_UpdateTask_Call) Return(_a0 *UpdateTaskResponse, _a1 error) *MockTaskServiceClient_UpdateTask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTaskServiceClient_UpdateTask_Call) RunAndReturn(run func(context.Context, *UpdateTaskRequest, ...grpc.CallOption) (*UpdateTaskResponse, error)) *MockTaskServiceClient_UpdateTask_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTaskServiceClient creates a new instance of MockTaskServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTaskServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTaskServiceClient {
	mock := &MockTaskServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
