package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Zubayear/song-ql/ent"
	"github.com/Zubayear/song-ql/graph/generated"
	"github.com/Zubayear/song-ql/graph/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (r *mutationResolver) CreateSong(ctx context.Context, input model.NewSong) (*model.Song, error) {
	r._logger.Info("Received request for CreateSong", zap.Any("request", input))
	songToSave := &ent.Song{
		Title:       input.Title,
		Duration:    input.Duration,
		LyricsExits: input.LyricsExists,
	}
	songToFromRepo, err := r.Resolver.songRepo.AddSong(ctx, songToSave)
	if err != nil {
		return nil, err
	}
	songToReturn := &model.Song{}
	songMapper(songToFromRepo, songToReturn)
	return songToReturn, nil
}

func (r *mutationResolver) CreateArtist(ctx context.Context, input model.NewArtist) (*model.Artist, error) {
	r._logger.Info("Received request for CreateArtist", zap.Any("request", input))
	artistToSave := &ent.Artist{
		Name: input.Name,
		Age:  input.Age,
	}
	songId, err := uuid.Parse(input.SongID)
	if err != nil {
		return nil, fmt.Errorf("failed parsing songId: %w", err)
	}
	artistFromRepo, err := r.songRepo.AddArtist(ctx, artistToSave, songId)
	if err != nil {
		return nil, err
	}
	artistToReturn := &model.Artist{
		ID:   artistFromRepo.ID.String(),
		Name: artistFromRepo.Name,
		Age:  artistFromRepo.Age,
	}
	return artistToReturn, nil
}

func (r *queryResolver) Songs(ctx context.Context) ([]*model.Song, error) {
	r._logger.Info("Received request for all songs")
	songsFromRepo, err := r.songRepo.GetSongs(ctx)
	if err != nil {
		return nil, err
	}
	var songsToReturn []*model.Song
	for _, song := range songsFromRepo {
		allArtistsBySongId, err := r.GetArtistsBySongID(ctx, song.ID.String())
		if err != nil {
			return nil, err
		}
		mappedSong := &model.Song{
			ID:           song.ID.String(),
			Title:        song.Title,
			Duration:     song.Duration,
			LyricsExists: song.LyricsExits,
			Artists:      allArtistsBySongId,
		}
		songMapper(song, mappedSong)
		songsToReturn = append(songsToReturn, mappedSong)
	}
	return songsToReturn, nil
}

func (r *queryResolver) GetSongByID(ctx context.Context, input string) (*model.Song, error) {
	r._logger.Info("Received request for all GetSongByID", zap.Any("request", input))
	songId, err := uuid.Parse(input)
	if err != nil {
		return nil, fmt.Errorf("failed parsing songId: %w", err)
	}
	songFromRepo, err := r.songRepo.GetSongById(ctx, songId)
	if err != nil {
		return nil, err
	}
	allArtistsBySongId, err := r.GetArtistsBySongID(ctx, input)
	if err != nil {
		return nil, err
	}
	songs := &model.Song{
		ID:           songFromRepo.ID.String(),
		Title:        songFromRepo.Title,
		Duration:     songFromRepo.Duration,
		LyricsExists: songFromRepo.LyricsExits,
		Artists:      allArtistsBySongId,
	}
	return songs, nil
}

func (r *queryResolver) GetArtistsBySongID(ctx context.Context, input string) ([]*model.Artist, error) {
	r._logger.Info("Received request for all GetArtistsBySongID", zap.Any("request", input))
	songId, err := uuid.Parse(input)
	if err != nil {
		return nil, fmt.Errorf("failed parsing songId: %w", err)
	}
	artistsFromRepo, err := r.songRepo.GetArtistsBySongId(ctx, songId)
	if err != nil {
		return nil, err
	}
	var artistsToReturn []*model.Artist
	for _, artist := range artistsFromRepo {
		mappedArtist := &model.Artist{}
		mapArtist(artist, mappedArtist)
		artistsToReturn = append(artistsToReturn, mappedArtist)
	}
	return artistsToReturn, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) GetArtistsBySongID(ctx context.Context, input string) ([]*model.Artist, error) {
	r._logger.Info("Received request for all GetArtistsBySongID", zap.Any("request", input))
	songId, err := uuid.Parse(input)
	if err != nil {
		return nil, fmt.Errorf("failed parsing songId: %w", err)
	}
	artistsFromRepo, err := r.songRepo.GetArtistsBySongId(ctx, songId)
	if err != nil {
		return nil, err
	}
	var artistsToReturn []*model.Artist
	for _, artist := range artistsFromRepo {
		mappedArtist := &model.Artist{}
		mapArtist(artist, mappedArtist)
		artistsToReturn = append(artistsToReturn, mappedArtist)
	}
	return artistsToReturn, nil
}
func (r *mutationResolver) GetSongByID(ctx context.Context, input string) (*model.Song, error) {
	r._logger.Info("Received request for all GetSongByID", zap.Any("request", input))
	songId, err := uuid.Parse(input)
	if err != nil {
		return nil, fmt.Errorf("failed parsing songId: %w", err)
	}
	songFromRepo, err := r.songRepo.GetSongById(ctx, songId)
	if err != nil {
		return nil, err
	}
	allArtistsBySongId, err := r.GetArtistsBySongID(ctx, input)
	if err != nil {
		return nil, err
	}
	songs := &model.Song{
		ID:           songFromRepo.ID.String(),
		Title:        songFromRepo.Title,
		Duration:     songFromRepo.Duration,
		LyricsExists: songFromRepo.LyricsExits,
		Artists:      allArtistsBySongId,
	}
	return songs, nil
}
func mapArtist(src *ent.Artist, dst *model.Artist) {
	dst.ID = src.ID.String()
	dst.Name = src.Name
	dst.Age = src.Age
}
func songMapper(src *ent.Song, dst *model.Song) {
	dst.ID = src.ID.String()
	dst.Title = src.Title
	dst.Duration = src.Duration
	//dst.Artists = src.Artist
	dst.LyricsExists = src.LyricsExits
}