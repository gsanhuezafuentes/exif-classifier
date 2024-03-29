// Code generated by mockery v2.40.1. DO NOT EDIT.

package exif_reader

import (
	exif_reader "github.com/gsanhuezafuentes/exif-classifier/exif_reader"
	mock "github.com/stretchr/testify/mock"
)

// MockExifReader is an autogenerated mock type for the ExifReader type
type MockExifReader struct {
	mock.Mock
}

type MockExifReader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExifReader) EXPECT() *MockExifReader_Expecter {
	return &MockExifReader_Expecter{mock: &_m.Mock}
}

// Read provides a mock function with given fields: filename
func (_m *MockExifReader) Read(filename string) (*exif_reader.ExifData, error) {
	ret := _m.Called(filename)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 *exif_reader.ExifData
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*exif_reader.ExifData, error)); ok {
		return rf(filename)
	}
	if rf, ok := ret.Get(0).(func(string) *exif_reader.ExifData); ok {
		r0 = rf(filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*exif_reader.ExifData)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExifReader_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockExifReader_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - filename string
func (_e *MockExifReader_Expecter) Read(filename interface{}) *MockExifReader_Read_Call {
	return &MockExifReader_Read_Call{Call: _e.mock.On("Read", filename)}
}

func (_c *MockExifReader_Read_Call) Run(run func(filename string)) *MockExifReader_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockExifReader_Read_Call) Return(_a0 *exif_reader.ExifData, _a1 error) *MockExifReader_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExifReader_Read_Call) RunAndReturn(run func(string) (*exif_reader.ExifData, error)) *MockExifReader_Read_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExifReader creates a new instance of MockExifReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExifReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExifReader {
	mock := &MockExifReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
