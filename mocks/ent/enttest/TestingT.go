// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TestingT is an autogenerated mock type for the TestingT type
type TestingT struct {
	mock.Mock
}

// Error provides a mock function with given fields: _a0
func (_m *TestingT) Error(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *TestingT) FailNow() {
	_m.Called()
}