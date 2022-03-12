//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Zubayear/song-ql/graph"
	"github.com/Zubayear/song-ql/logger"
	"github.com/Zubayear/song-ql/repository"
	"github.com/google/wire"
)

func DependencyProvider() (*graph.Resolver, error) {
	wire.Build(repository.DatabaseImlProvider, graph.ResolverProvider,
		wire.Bind(new(repository.ISongRepository), new(*repository.SongRepository)),
		logger.LoggerProvider)
	return &graph.Resolver{}, nil
}
