package mocks

import (
	"testing"
)

func TestResolverRoot_Mutation(t *testing.T) {
	testObj := new(ResolverRoot)
	testObj.On("Mutation").Return(nil)
	testObj.Mutation()
	testObj.AssertExpectations(t)
}

func TestResolverRoot_Query(t *testing.T) {
	testObj := new(ResolverRoot)
	testObj.On("Query").Return(nil)
	testObj.Query()
	testObj.AssertExpectations(t)
}
