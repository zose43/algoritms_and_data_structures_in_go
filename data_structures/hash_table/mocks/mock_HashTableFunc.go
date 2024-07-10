// Code generated by mockery v2.38.0. DO NOT EDIT.

package hash_table

import mock "github.com/stretchr/testify/mock"

// MockHashTableFunc is an autogenerated mock type for the HashTableFunc type
type MockHashTableFunc[T interface{}] struct {
	mock.Mock
}

type MockHashTableFunc_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *MockHashTableFunc[T]) EXPECT() *MockHashTableFunc_Expecter[T] {
	return &MockHashTableFunc_Expecter[T]{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: val
func (_m *MockHashTableFunc[T]) Execute(val T) {
	_m.Called(val)
}

// MockHashTableFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockHashTableFunc_Execute_Call[T interface{}] struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - val T
func (_e *MockHashTableFunc_Expecter[T]) Execute(val interface{}) *MockHashTableFunc_Execute_Call[T] {
	return &MockHashTableFunc_Execute_Call[T]{Call: _e.mock.On("Execute", val)}
}

func (_c *MockHashTableFunc_Execute_Call[T]) Run(run func(val T)) *MockHashTableFunc_Execute_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(T))
	})
	return _c
}

func (_c *MockHashTableFunc_Execute_Call[T]) Return() *MockHashTableFunc_Execute_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockHashTableFunc_Execute_Call[T]) RunAndReturn(run func(T)) *MockHashTableFunc_Execute_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewMockHashTableFunc creates a new instance of MockHashTableFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHashTableFunc[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHashTableFunc[T] {
	mock := &MockHashTableFunc[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}