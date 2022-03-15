package mocks

import (
	"context"
	"github.com/Zubayear/song-ql/ent"
	"github.com/google/uuid"
	"testing"
)

func TestISongRepository_AddArtist(t *testing.T) {
	testObj := new(ISongRepository)
	ctx := context.Background()
	var input *ent.Artist
	expecting := &ent.Artist{
		ID:   uuid.New(),
		Name: "Artist",
		Age:  40,
	}
	testObj.On("AddArtist", ctx, input).Return(expecting, nil)
	testObj.AddArtist(ctx, input)
	testObj.AssertExpectations(t)
}

func TestISongRepository_AddArtists(t *testing.T) {
	testObj := new(ISongRepository)
	ctx := context.Background()
	var input []*ent.Artist
	var expecting []*ent.Artist
	testObj.On("AddArtists", ctx, input).Return(expecting, nil)
	testObj.AddArtists(ctx, input)
	testObj.AssertExpectations(t)
}

func TestISongRepository_AddSongWithArtistId(t *testing.T) {
	testObj := new(ISongRepository)
	ctx := context.Background()
	var input1 *ent.Song
	var input2 []uuid.UUID
	expecting := &ent.Song{
		ID:          uuid.New(),
		Title:       "Song",
		Duration:    4.50,
		LyricsExits: false,
	}
	testObj.On("AddSongWithArtistId", ctx, input1, input2).Return(expecting, nil)
	testObj.AddSongWithArtistId(ctx, input1, input2)
	testObj.AssertExpectations(t)
}

func TestISongRepository_GetArtistById(t *testing.T) {
	testObj := new(ISongRepository)
	ctx := context.Background()
	input := uuid.New()
	expecting := &ent.Artist{
		ID:   uuid.New(),
		Name: "Artist",
		Age:  40,
	}
	testObj.On("GetArtistById", ctx, input).Return(expecting, nil)
	testObj.GetArtistById(ctx, input)
	testObj.AssertExpectations(t)
}

func TestISongRepository_GetArtists(t *testing.T) {
	testObj := new(ISongRepository)
	ctx := context.Background()
	var expecting []*ent.Artist
	testObj.On("GetArtists", ctx).Return(expecting, nil)
	testObj.GetArtists(ctx)
	testObj.AssertExpectations(t)
}

func TestISongRepository_GetArtistsBySongId(t *testing.T) {
	testObj := new(ISongRepository)
	ctx := context.Background()
	input := uuid.New()
	var expecting []*ent.Artist
	testObj.On("GetArtistsBySongId", ctx, input).Return(expecting, nil)
	testObj.GetArtistsBySongId(ctx, input)
	testObj.AssertExpectations(t)
}

func TestISongRepository_GetSongById(t *testing.T) {
	testObj := new(ISongRepository)
	ctx := context.Background()
	input := uuid.New()
	var expecting *ent.Song
	testObj.On("GetSongById", ctx, input).Return(expecting, nil)
	testObj.GetSongById(ctx, input)
	testObj.AssertExpectations(t)
}

func TestISongRepository_GetSongs(t *testing.T) {
	testObj := new(ISongRepository)
	ctx := context.Background()
	var expecting []*ent.Song
	testObj.On("GetSongs", ctx).Return(expecting, nil)
	testObj.GetSongs(ctx)
	testObj.AssertExpectations(t)
}
