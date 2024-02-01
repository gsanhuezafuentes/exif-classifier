// Code generated by mockery v2.40.1. DO NOT EDIT.

package commands

import mock "github.com/stretchr/testify/mock"

// MockGroupCmdFileOperation is an autogenerated mock type for the GroupCmdFileOperation type
type MockGroupCmdFileOperation struct {
	mock.Mock
}

type MockGroupCmdFileOperation_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGroupCmdFileOperation) EXPECT() *MockGroupCmdFileOperation_Expecter {
	return &MockGroupCmdFileOperation_Expecter{mock: &_m.Mock}
}

// GetCurrentDirectory provides a mock function with given fields:
func (_m *MockGroupCmdFileOperation) GetCurrentDirectory() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentDirectory")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupCmdFileOperation_GetCurrentDirectory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrentDirectory'
type MockGroupCmdFileOperation_GetCurrentDirectory_Call struct {
	*mock.Call
}

// GetCurrentDirectory is a helper method to define mock.On call
func (_e *MockGroupCmdFileOperation_Expecter) GetCurrentDirectory() *MockGroupCmdFileOperation_GetCurrentDirectory_Call {
	return &MockGroupCmdFileOperation_GetCurrentDirectory_Call{Call: _e.mock.On("GetCurrentDirectory")}
}

func (_c *MockGroupCmdFileOperation_GetCurrentDirectory_Call) Run(run func()) *MockGroupCmdFileOperation_GetCurrentDirectory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockGroupCmdFileOperation_GetCurrentDirectory_Call) Return(_a0 string, _a1 error) *MockGroupCmdFileOperation_GetCurrentDirectory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGroupCmdFileOperation_GetCurrentDirectory_Call) RunAndReturn(run func() (string, error)) *MockGroupCmdFileOperation_GetCurrentDirectory_Call {
	_c.Call.Return(run)
	return _c
}

// GetImageFilesPathFromDirectory provides a mock function with given fields: _a0
func (_m *MockGroupCmdFileOperation) GetImageFilesPathFromDirectory(_a0 string) ([]string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetImageFilesPathFromDirectory")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupCmdFileOperation_GetImageFilesPathFromDirectory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetImageFilesPathFromDirectory'
type MockGroupCmdFileOperation_GetImageFilesPathFromDirectory_Call struct {
	*mock.Call
}

// GetImageFilesPathFromDirectory is a helper method to define mock.On call
//   - _a0 string
func (_e *MockGroupCmdFileOperation_Expecter) GetImageFilesPathFromDirectory(_a0 interface{}) *MockGroupCmdFileOperation_GetImageFilesPathFromDirectory_Call {
	return &MockGroupCmdFileOperation_GetImageFilesPathFromDirectory_Call{Call: _e.mock.On("GetImageFilesPathFromDirectory", _a0)}
}

func (_c *MockGroupCmdFileOperation_GetImageFilesPathFromDirectory_Call) Run(run func(_a0 string)) *MockGroupCmdFileOperation_GetImageFilesPathFromDirectory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockGroupCmdFileOperation_GetImageFilesPathFromDirectory_Call) Return(_a0 []string, _a1 error) *MockGroupCmdFileOperation_GetImageFilesPathFromDirectory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGroupCmdFileOperation_GetImageFilesPathFromDirectory_Call) RunAndReturn(run func(string) ([]string, error)) *MockGroupCmdFileOperation_GetImageFilesPathFromDirectory_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGroupCmdFileOperation creates a new instance of MockGroupCmdFileOperation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGroupCmdFileOperation(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGroupCmdFileOperation {
	mock := &MockGroupCmdFileOperation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}