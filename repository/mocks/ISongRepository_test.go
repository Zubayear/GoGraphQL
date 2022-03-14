package mocks

import (
	"context"
	"github.com/Zubayear/song-ql/ent"
	. "github.com/Zubayear/song-ql/graph"
	"github.com/Zubayear/song-ql/logger"
	"github.com/google/uuid"
	"log"
	"testing"
)

func TestISongRepository_GetSongById(t *testing.T) {
	testObj := new(ISongRepository)
	song := &ent.Song{}
	testObj.On("GetSongById", context.Background(), uuid.MustParse("f5e8b74a-486e-40f2-8f6d-a5095f0e0755")).Return(song, nil)

	resolver := Resolver{
		SongRepo: testObj,
	}
	resolver.SongRepo.GetSongById(context.Background(), uuid.MustParse("f5e8b74a-486e-40f2-8f6d-a5095f0e0755"))
	testObj.AssertExpectations(t)
}

func TestISongRepository_AddArtist(t *testing.T) {
	testObj := new(ISongRepository)
	artist := &ent.Artist{
		Name: "J Hus",
		Age:  25,
	}
	ctx := context.Background()
	testObj.On("AddArtist", ctx, artist).Return(artist, nil)

	resolver := Resolver{
		SongRepo: testObj,
	}
	resolver.SongRepo.AddArtist(ctx, artist)
	testObj.AssertExpectations(t)
}

func TestISongRepository_AddSongWithArtistId(t *testing.T) {
	testObj := new(ISongRepository)
	song := &ent.Song{
		Title:       "Some Song",
		Duration:    4.56,
		LyricsExits: false,
	}
	song2 := &ent.Song{
		Title:       "Some Song",
		Duration:    4.59,
		LyricsExits: true,
	}
	ids := []uuid.UUID{
		uuid.MustParse("b198ab8d-28d5-4e79-a557-1afa7e2f4f62"),
		uuid.MustParse("c00c6959-835c-442b-9c57-3f79508a3dc7"),
	}
	ctx := context.Background()
	testObj.On("AddSongWithArtistId", ctx, song, ids).Return(song2, nil)

	resolver := Resolver{
		SongRepo: testObj,
	}
	songFromRepo, err := resolver.SongRepo.AddSongWithArtistId(ctx, song, ids)
	if err != nil {
		return
	}
	log.Printf("songFromRepo: %v", songFromRepo)
	testObj.AssertExpectations(t)
}

func TestISongRepository_GetArtistById(t *testing.T) {
	testObj := new(ISongRepository)
	id := uuid.MustParse("b198ab8d-28d5-4e79-a557-1afa7e2f4f62")
	artist := &ent.Artist{}
	ctx := context.Background()
	testObj.On("GetArtistById", ctx, id).Return(artist, nil)
	_, _ = logger.LoggerProvider()

	resolver := Resolver{
		SongRepo: testObj,
	}
	resolver.SongRepo.GetArtistById(ctx, id)
	testObj.AssertExpectations(t)
}

func TestISongRepository_AddArtists(t *testing.T) {
	testObj := new(ISongRepository)
	song := &ent.Song{
		Title:       "Some Song",
		Duration:    4.56,
		LyricsExits: false,
	}
	song2 := &ent.Song{
		Title:       "Some Song",
		Duration:    4.59,
		LyricsExits: true,
	}
	ids := []uuid.UUID{
		uuid.MustParse("b198ab8d-28d5-4e79-a557-1afa7e2f4f62"),
		uuid.MustParse("c00c6959-835c-442b-9c57-3f79508a3dc7"),
	}
	ctx := context.Background()
	testObj.On("AddSongWithArtistId", ctx, song, ids).Return(song2, nil)
	_, _ = logger.LoggerProvider()

	resolver := Resolver{
		SongRepo: testObj,
	}
	resolver.SongRepo.AddSongWithArtistId(ctx, song, ids)
	testObj.AssertExpectations(t)
}

func TestISongRepository_GetArtists(t *testing.T) {
	testObj := new(ISongRepository)
	var artists []*ent.Artist
	ctx := context.Background()
	testObj.On("GetArtists", ctx).Return(artists, nil)
	_, _ = logger.LoggerProvider()

	resolver := Resolver{
		SongRepo: testObj,
	}
	resolver.SongRepo.GetArtists(ctx)
	testObj.AssertExpectations(t)
}

func TestISongRepository_GetArtistsBySongId(t *testing.T) {
	testObj := new(ISongRepository)
	var artists []*ent.Artist
	id := uuid.MustParse("b198ab8d-28d5-4e79-a557-1afa7e2f4f62")
	ctx := context.Background()
	testObj.On("GetArtistsBySongId", ctx, id).Return(artists, nil)
	_, _ = logger.LoggerProvider()

	resolver := Resolver{
		SongRepo: testObj,
	}
	resolver.SongRepo.GetArtistsBySongId(ctx, id)
	testObj.AssertExpectations(t)

}

func TestISongRepository_GetSongs(t *testing.T) {
	testObj := new(ISongRepository)
	var songs []*ent.Song
	ctx := context.Background()
	testObj.On("GetSongs", ctx).Return(songs, nil)
	_, _ = logger.LoggerProvider()

	resolver := Resolver{
		SongRepo: testObj,
	}
	resolver.SongRepo.GetSongs(ctx)
	testObj.AssertExpectations(t)
}
