//+build wireinject

package wire

import (
	"github.com/google/wire"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/config"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/handler"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/repository"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/server"
	"go.uber.org/zap"
)

func InitializeConfig(logger *zap.Logger) *config.Registry {
	wire.Build(config.NewRegistry, config.NewPostgresConfig, config.NewGinConfig)
	return &config.Registry{}
}

func InitializeServer(todoRepository repository.TodoRepository, logger *zap.Logger) *server.Server {
	wire.Build(server.NewServer, handler.NewTodoHandler, config.NewPostgresConfig, config.NewRegistry, config.NewGinConfig)
	return &server.Server{}
}
