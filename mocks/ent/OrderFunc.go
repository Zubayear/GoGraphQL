// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	sql "entgo.io/ent/dialect/sql"
	mock "github.com/stretchr/testify/mock"
)

// OrderFunc is an autogenerated mock type for the OrderFunc type
type OrderFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *OrderFunc) Execute(_a0 *sql.Selector) {
	_m.Called(_a0)
}
