package mocks

import (
	"context"
	"github.com/Zubayear/song-ql/ent"
	. "github.com/Zubayear/song-ql/graph"
	"github.com/Zubayear/song-ql/logger"
	"github.com/google/uuid"
	"testing"
)

func TestISongRepository_GetSongById(t *testing.T) {
	testObj := new(ISongRepository)
	song := &ent.Song{}
	testObj.On("GetSongById", context.Background(), uuid.MustParse("f5e8b74a-486e-40f2-8f6d-a5095f0e0755")).Return(song, nil)
	loggerProvider := logger.LoggerProvider()

	resolver := Resolver{
		Logger:   loggerProvider,
		SongRepo: testObj,
	}
	resolver.SongRepo.GetSongById(context.Background(), uuid.MustParse("f5e8b74a-486e-40f2-8f6d-a5095f0e0755"))
	testObj.AssertExpectations(t)
}

func TestISongRepository_AddSong(t *testing.T) {
	testObj := new(ISongRepository)
	song := &ent.Song{
		Title:       "Demo Song",
		Duration:    0,
		LyricsExits: false,
	}
	testObj.On("AddSong", context.Background(), song).Return(nil, nil)
	loggerProvider := logger.LoggerProvider()

	resolver := Resolver{
		Logger:   loggerProvider,
		SongRepo: testObj,
	}
	resolver.SongRepo.AddSong(context.Background(), song)
	testObj.AssertExpectations(t)
}
