// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LightNode is an autogenerated mock type for the LightNode type
type LightNode struct {
	mock.Mock
}

type LightNode_Expecter struct {
	mock *mock.Mock
}

func (_m *LightNode) EXPECT() *LightNode_Expecter {
	return &LightNode_Expecter{mock: &_m.Mock}
}

// NewLightNode creates a new instance of LightNode. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLightNode(t interface {
	mock.TestingT
	Cleanup(func())
}) *LightNode {
	mock := &LightNode{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
