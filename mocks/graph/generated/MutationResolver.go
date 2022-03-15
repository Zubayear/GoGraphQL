// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/Zubayear/song-ql/graph/model"
)

// MutationResolver is an autogenerated mock type for the MutationResolver type
type MutationResolver struct {
	mock.Mock
}

// CreateArtist provides a mock function with given fields: ctx, input
func (_m *MutationResolver) CreateArtist(ctx context.Context, input model.NewArtist) (*model.Artist, error) {
	ret := _m.Called(ctx, input)

	var r0 *model.Artist
	if rf, ok := ret.Get(0).(func(context.Context, model.NewArtist) *model.Artist); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Artist)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.NewArtist) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateArtists provides a mock function with given fields: ctx, input
func (_m *MutationResolver) CreateArtists(ctx context.Context, input []*model.NewArtist) ([]*model.Artist, error) {
	ret := _m.Called(ctx, input)

	var r0 []*model.Artist
	if rf, ok := ret.Get(0).(func(context.Context, []*model.NewArtist) []*model.Artist); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Artist)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*model.NewArtist) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSong provides a mock function with given fields: ctx, input
func (_m *MutationResolver) CreateSong(ctx context.Context, input model.NewSong) (*model.Song, error) {
	ret := _m.Called(ctx, input)

	var r0 *model.Song
	if rf, ok := ret.Get(0).(func(context.Context, model.NewSong) *model.Song); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Song)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.NewSong) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}