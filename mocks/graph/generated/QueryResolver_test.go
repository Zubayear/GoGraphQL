package mocks

import (
	"context"
	"github.com/Zubayear/song-ql/graph/model"
	"github.com/google/uuid"
	"testing"
)

func TestQueryResolver_Artists(t *testing.T) {
	testObj := new(QueryResolver)
	ctx := context.Background()
	var expecting []*model.Artist
	testObj.On("Artists", ctx).Return(expecting, nil)
	testObj.Artists(ctx)
	testObj.AssertExpectations(t)
}

func TestQueryResolver_ArtistsBySongID(t *testing.T) {
	testObj := new(QueryResolver)
	ctx := context.Background()
	var input uuid.UUID
	var expecting []*model.Artist
	testObj.On("ArtistsBySongID", ctx, input.String()).Return(expecting, nil)
	testObj.ArtistsBySongID(ctx, input.String())
	testObj.AssertExpectations(t)
}

func TestQueryResolver_SongByID(t *testing.T) {
	testObj := new(QueryResolver)
	ctx := context.Background()
	var input uuid.UUID
	var expecting *model.Song
	testObj.On("SongByID", ctx, input.String()).Return(expecting, nil)
	testObj.SongByID(ctx, input.String())
	testObj.AssertExpectations(t)
}

func TestQueryResolver_Songs(t *testing.T) {
	testObj := new(QueryResolver)
	ctx := context.Background()
	var expecting []*model.Song
	testObj.On("Songs", ctx).Return(expecting, nil)
	testObj.Songs(ctx)
	testObj.AssertExpectations(t)
}
