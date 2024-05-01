// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// AuthInterceptor is an autogenerated mock type for the AuthInterceptor type
type AuthInterceptor struct {
	mock.Mock
}

type AuthInterceptor_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthInterceptor) EXPECT() *AuthInterceptor_Expecter {
	return &AuthInterceptor_Expecter{mock: &_m.Mock}
}

// StreamInterceptor provides a mock function with given fields: srv, ss, info, handler
func (_m *AuthInterceptor) StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ret := _m.Called(srv, ss, info, handler)

	if len(ret) == 0 {
		panic("no return value specified for StreamInterceptor")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, grpc.ServerStream, *grpc.StreamServerInfo, grpc.StreamHandler) error); ok {
		r0 = rf(srv, ss, info, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthInterceptor_StreamInterceptor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StreamInterceptor'
type AuthInterceptor_StreamInterceptor_Call struct {
	*mock.Call
}

// StreamInterceptor is a helper method to define mock.On call
//   - srv interface{}
//   - ss grpc.ServerStream
//   - info *grpc.StreamServerInfo
//   - handler grpc.StreamHandler
func (_e *AuthInterceptor_Expecter) StreamInterceptor(srv interface{}, ss interface{}, info interface{}, handler interface{}) *AuthInterceptor_StreamInterceptor_Call {
	return &AuthInterceptor_StreamInterceptor_Call{Call: _e.mock.On("StreamInterceptor", srv, ss, info, handler)}
}

func (_c *AuthInterceptor_StreamInterceptor_Call) Run(run func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler)) *AuthInterceptor_StreamInterceptor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(grpc.ServerStream), args[2].(*grpc.StreamServerInfo), args[3].(grpc.StreamHandler))
	})
	return _c
}

func (_c *AuthInterceptor_StreamInterceptor_Call) Return(_a0 error) *AuthInterceptor_StreamInterceptor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthInterceptor_StreamInterceptor_Call) RunAndReturn(run func(interface{}, grpc.ServerStream, *grpc.StreamServerInfo, grpc.StreamHandler) error) *AuthInterceptor_StreamInterceptor_Call {
	_c.Call.Return(run)
	return _c
}

// UnaryInterceptor provides a mock function with given fields: ctx, req, info, handler
func (_m *AuthInterceptor) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ret := _m.Called(ctx, req, info, handler)

	if len(ret) == 0 {
		panic("no return value specified for UnaryInterceptor")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error)); ok {
		return rf(ctx, req, info, handler)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) interface{}); ok {
		r0 = rf(ctx, req, info, handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) error); ok {
		r1 = rf(ctx, req, info, handler)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthInterceptor_UnaryInterceptor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnaryInterceptor'
type AuthInterceptor_UnaryInterceptor_Call struct {
	*mock.Call
}

// UnaryInterceptor is a helper method to define mock.On call
//   - ctx context.Context
//   - req interface{}
//   - info *grpc.UnaryServerInfo
//   - handler grpc.UnaryHandler
func (_e *AuthInterceptor_Expecter) UnaryInterceptor(ctx interface{}, req interface{}, info interface{}, handler interface{}) *AuthInterceptor_UnaryInterceptor_Call {
	return &AuthInterceptor_UnaryInterceptor_Call{Call: _e.mock.On("UnaryInterceptor", ctx, req, info, handler)}
}

func (_c *AuthInterceptor_UnaryInterceptor_Call) Run(run func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)) *AuthInterceptor_UnaryInterceptor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(*grpc.UnaryServerInfo), args[3].(grpc.UnaryHandler))
	})
	return _c
}

func (_c *AuthInterceptor_UnaryInterceptor_Call) Return(_a0 interface{}, _a1 error) *AuthInterceptor_UnaryInterceptor_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthInterceptor_UnaryInterceptor_Call) RunAndReturn(run func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error)) *AuthInterceptor_UnaryInterceptor_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthInterceptor creates a new instance of AuthInterceptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthInterceptor(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthInterceptor {
	mock := &AuthInterceptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}