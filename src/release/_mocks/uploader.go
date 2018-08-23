// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import io "io"
import mock "github.com/stretchr/testify/mock"

// uploader is an autogenerated mock type for the uploader type
type Uploader struct {
	mock.Mock
}

// file provides a mock function with given fields: _a0, _a1
func (_m *Uploader) File(_a0 string, _a1 io.ReadSeeker) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, io.ReadSeeker) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, io.ReadSeeker) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// json provides a mock function with given fields: _a0, _a1
func (_m *Uploader) JSON(_a0 string, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}