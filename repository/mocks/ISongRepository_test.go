package mocks

import (
	"context"
	"github.com/Zubayear/song-ql/ent"
	. "github.com/Zubayear/song-ql/graph"
	"github.com/Zubayear/song-ql/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"log"
	"reflect"
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

func TestISongRepository_AddArtist(t *testing.T) {
	testObj := new(ISongRepository)
	artist := &ent.Artist{
		Name: "J Hus",
		Age:  25,
	}
	ctx := context.Background()
	testObj.On("AddArtist", ctx, artist).Return(artist, nil)
	loggerProvider := logger.LoggerProvider()

	resolver := Resolver{
		Logger:   loggerProvider,
		SongRepo: testObj,
	}
	addArtist, err := resolver.SongRepo.AddArtist(ctx, artist)
	if err != nil {
		return
	}
	loggerProvider.Info("Result of mocking", zap.Any("addArtist", addArtist))
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
	loggerProvider := logger.LoggerProvider()

	resolver := Resolver{
		Logger:   loggerProvider,
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
	loggerProvider := logger.LoggerProvider()

	resolver := Resolver{
		Logger:   loggerProvider,
		SongRepo: testObj,
	}
	songFromRepo, err := resolver.SongRepo.AddSongWithArtistId(ctx, song, ids)
	if err != nil {
		return
	}
	log.Printf("songFromRepo: %v", songFromRepo)
	testObj.AssertExpectations(t)
}

func TestISongRepository_GetArtists(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ent.Artist
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ISongRepository{
				Mock: tt.fields.Mock,
			}
			got, err := _m.GetArtists(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArtists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArtists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISongRepository_GetArtistsBySongId(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx    context.Context
		songId uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ent.Artist
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ISongRepository{
				Mock: tt.fields.Mock,
			}
			got, err := _m.GetArtistsBySongId(tt.args.ctx, tt.args.songId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArtistsBySongId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArtistsBySongId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISongRepository_GetSongById1(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx    context.Context
		songId uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ent.Song
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ISongRepository{
				Mock: tt.fields.Mock,
			}
			got, err := _m.GetSongById(tt.args.ctx, tt.args.songId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSongById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSongById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISongRepository_GetSongs(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ent.Song
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ISongRepository{
				Mock: tt.fields.Mock,
			}
			got, err := _m.GetSongs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSongs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSongs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
