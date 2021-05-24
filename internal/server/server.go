package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/config"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/handler"
	"go.uber.org/zap"
)

// Server the gin-gonic server.
type Server struct {
	logger      *zap.Logger
	engine      *gin.Engine
	port        int
	todoHandler handler.TodoHandler
}

func (s *Server) addRoutes() {
	s.engine.GET("/todos", s.findAllTodo)
	s.engine.POST("/todos", s.createTodo)
	s.engine.GET("/todos/:id", s.findTodo)
	s.engine.PUT("/todos/:id", s.updateTodo)
	s.engine.DELETE("/todos/:id", s.deleteTodo)
}

// Run starts the server and listens.
func (s *Server) Run() {
	s.addRoutes()

	err := s.engine.Run(fmt.Sprintf(":%d", s.port))
	if err != nil {
		s.logger.Fatal("could not start server")
	}
}

// NewServer returns a new Server.
func NewServer(config *config.Registry, todoHandler handler.TodoHandler, logger *zap.Logger) *Server {
	return &Server{
		logger:      logger,
		engine:      gin.New(),
		port:        config.GinConfig.Port,
		todoHandler: todoHandler,
	}
}
