// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/config"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/handler"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/repository"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/server"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func InitializeConfig(logger *zap.Logger) *config.Registry {
	ginConfig := config.NewGinConfig()
	postgresConfig := config.NewPostgresConfig()
	registry := config.NewRegistry(ginConfig, postgresConfig, logger)
	return registry
}

func InitializeServer(todoRepository repository.TodoRepository, logger *zap.Logger) *server.Server {
	ginConfig := config.NewGinConfig()
	postgresConfig := config.NewPostgresConfig()
	registry := config.NewRegistry(ginConfig, postgresConfig, logger)
	todoHandler := handler.NewTodoHandler(todoRepository, logger)
	serverServer := server.NewServer(registry, todoHandler, logger)
	return serverServer
}
