package graph

import (
	"github.com/Zubayear/song-ql/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SongRepo repository.ISongRepository
}

func ResolverProvider(repo repository.ISongRepository) (*Resolver, error) {
	return &Resolver{SongRepo: repo}, nil
}
