package repository

import (
	"context"
	"fmt"
	"github.com/Zubayear/song-ql/ent"
	"github.com/Zubayear/song-ql/ent/migrate"
	"github.com/Zubayear/song-ql/ent/song"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
)

type ISongRepository interface {
	AddSongWithArtistId(ctx context.Context, song *ent.Song, artistId []uuid.UUID) (*ent.Song, error)
	GetSongs(ctx context.Context) ([]*ent.Song, error)
	AddArtist(ctx context.Context, artist *ent.Artist) (*ent.Artist, error)
	GetSongById(ctx context.Context, songId uuid.UUID) (*ent.Song, error)
	GetArtistsBySongId(ctx context.Context, songId uuid.UUID) ([]*ent.Artist, error)
	GetArtistById(ctx context.Context, artistIds uuid.UUID) (*ent.Artist, error)
	GetArtists(ctx context.Context) ([]*ent.Artist, error)
}

type SongRepository struct {
	songClient *ent.Client
}

func (s *SongRepository) GetArtists(ctx context.Context) ([]*ent.Artist, error) {
	allArtists, err := s.songClient.Artist.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting all artists: %w", err)
	}
	return allArtists, nil
}

func (s *SongRepository) GetArtistById(ctx context.Context, artistId uuid.UUID) (*ent.Artist, error) {
	artist, err := s.songClient.Artist.Get(ctx, artistId)
	if err != nil {
		return nil, fmt.Errorf("failed getting artist: %v", err)
	}
	return artist, nil
}

func (s *SongRepository) GetArtistsBySongId(ctx context.Context, songId uuid.UUID) ([]*ent.Artist, error) {
	allArtists, err := s.songClient.Song.Query().Where(song.ID(songId)).QueryArtists().All(ctx)
	log.Println(allArtists)
	if err != nil {
		return nil, fmt.Errorf("failed getting all artists: %w", err)
	}
	return allArtists, nil
}

func (s *SongRepository) GetSongById(ctx context.Context, songId uuid.UUID) (*ent.Song, error) {
	songFromRepo, err := s.songClient.Song.Get(ctx, songId)
	if err != nil {
		return nil, fmt.Errorf("failed getting song: %w", err)
	}
	return songFromRepo, nil
}

func (s *SongRepository) AddArtist(ctx context.Context, artist *ent.Artist) (*ent.Artist, error) {
	savedArtist, err := s.songClient.Artist.Create().
		SetName(artist.Name).
		SetAge(artist.Age).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating artist: %w", err)
	}
	return savedArtist, err
}

func (s *SongRepository) AddSongWithArtistId(ctx context.Context, song *ent.Song, artistIds []uuid.UUID) (*ent.Song, error) {
	savedSong, err := s.songClient.Song.Create().
		SetTitle(song.Title).
		SetDuration(song.Duration).
		SetLyricsExits(song.LyricsExits).AddArtistIDs(artistIds...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed saving song with artist id: %w", err)
	}
	return savedSong, nil
}

func (s *SongRepository) GetSongs(ctx context.Context) ([]*ent.Song, error) {
	allSongs, err := s.songClient.Song.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting all songs: %w", err)
	}
	return allSongs, nil
}

// NewSongRepository Utility function to auto generate interface methods with IDE help
func NewSongRepository() ISongRepository {
	return &SongRepository{}
}

func DatabaseImlProvider() (*SongRepository, error) {
	client, err := ent.Open("mysql", "root:root@tcp(localhost:3306)/SongsDB?parseTime=true")
	//client, err := ent.Open(h.Type, connString)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	//defer client.Close()
	ctx := context.Background()
	// Run migration.
	if err := client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &SongRepository{songClient: client}, nil
}
