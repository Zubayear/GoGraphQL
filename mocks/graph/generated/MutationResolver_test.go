package mocks

import (
	"context"
	"github.com/Zubayear/song-ql/graph/model"
	"testing"
)

func TestMutationResolver_CreateArtist(t *testing.T) {
	testObj := new(MutationResolver)
	ctx := context.Background()
	var input model.NewArtist
	var expecting *model.Artist
	testObj.On("CreateArtist", ctx, input).Return(expecting, nil)
	testObj.CreateArtist(ctx, input)
	testObj.AssertExpectations(t)
}

func TestMutationResolver_CreateArtists(t *testing.T) {
	testObj := new(MutationResolver)
	ctx := context.Background()
	var input []*model.NewArtist
	var expecting []*model.Artist
	testObj.On("CreateArtists", ctx, input).Return(expecting, nil)
	testObj.CreateArtists(ctx, input)
	testObj.AssertExpectations(t)
}

func TestMutationResolver_CreateSong(t *testing.T) {
	testObj := new(MutationResolver)
	ctx := context.Background()
	var input model.NewSong
	var expecting *model.Song
	testObj.On("CreateSong", ctx, input).Return(expecting, nil)
	testObj.CreateSong(ctx, input)
	testObj.AssertExpectations(t)
}
