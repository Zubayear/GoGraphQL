package graph

import (
	"github.com/Zubayear/song-ql/repository"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	_logger  *zap.Logger
	songRepo repository.ISongRepository
}

func ResolverProvider(_logger *zap.Logger, repo repository.ISongRepository) (*Resolver, error) {
	return &Resolver{_logger: _logger, songRepo: repo}, nil
}
