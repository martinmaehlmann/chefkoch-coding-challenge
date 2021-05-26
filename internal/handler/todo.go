package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/repository"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/todo"
	"go.uber.org/zap"
)

// TodoHandler handles the retrival and persisting of Todos.
type TodoHandler interface {
	// FindAll returns all available Todos, or an empty slice, if none are available.
	FindAll() []*todo.Todo

	// Find returns a Todo by its id, if it exists.
	// Else returns a nil object
	// If the id is a non valid positive integer, returns an error with http code http.StatusBadRequest.
	Find(id string) (*todo.Todo, *todo.HandlerError)

	// Update updates the specified Todo, if it exists
	// returns nil and an error with http code http.StatusBadRequest, if it does not or the body data was malformed.
	Update(bodyData []byte, id string) (*todo.Todo, *todo.HandlerError)

	// Create creates the specified Todo and returns it with the updated id
	// If the specified Todo has an id assigned, a http.StatusBadRequest error will be returned.
	Create(bodyData []byte) (*todo.Todo, *todo.HandlerError)

	// Delete soft deletes the specified Todo cascading
	// if the Todo does not exist, a new TodoHandlerError is returned.
	Delete(id string) *todo.HandlerError
}

type todoHandler struct {
	logger     *zap.Logger
	repository repository.TodoRepository
}

// FindAll returns all available Todos, or an empty slice, if none are available.
func (s *todoHandler) FindAll() []*todo.Todo {
	return s.repository.FindAll()
}

// Find returns a Todo by its id, if it exists.
// Else returns a nil object
// If the id is a non valid positive integer, returns an error with http code http.StatusBadRequest.
func (s *todoHandler) Find(id string) (*todo.Todo, *todo.HandlerError) {
	validID, err := strconv.Atoi(id)
	if err != nil || validID < 0 {
		return nil, todo.NewTodoInvalidIDError(id)
	}

	return s.repository.Find(uint(validID)), nil
}

// Update updates the specified Todo, if it exists
// returns nil and an error with http code http.StatusBadRequest, if it does not or the body data was malformed.
func (s *todoHandler) Update(bodyData []byte, id string) (*todo.Todo, *todo.HandlerError) {
	validID, err := strconv.Atoi(id)
	if err != nil || validID < 0 {
		return nil, todo.NewTodoInvalidIDError(id)
	}

	toDo, err2 := unmarshalTodo(bodyData)
	if err2 != nil {
		return nil, err2
	}

	toDo.ID = uint(validID)

	if !toDo.Valid() {
		return nil, todo.NewInvalidTodo(toDo)
	}

	result := s.repository.Update(toDo)
	if result == nil {
		return nil, todo.NewTodoHandlerError(fmt.Sprintf("todo %v does not exist", toDo), http.StatusBadRequest)
	}

	return result, nil
}

// Create creates the specified Todo and returns it with the updated id
// If the specified Todo has an id assigned, a http.StatusBadRequest error will be returned.
func (s *todoHandler) Create(bodyData []byte) (*todo.Todo, *todo.HandlerError) {
	toDo, err := unmarshalTodo(bodyData)
	if err != nil {
		return nil, err
	}

	if !toDo.Valid() {
		return nil, todo.NewInvalidTodo(toDo)
	}

	result := s.repository.Create(toDo)

	return result, nil
}

// Delete soft deletes the specified Todo cascading
// if the Todo does not exist, a new TodoHandlerError is returned.
func (s *todoHandler) Delete(id string) *todo.HandlerError {
	validID, err := strconv.Atoi(id)
	if err != nil || validID < 0 {
		return todo.NewTodoInvalidIDError(id)
	}

	rowsAffected := s.repository.Delete(uint(validID))
	if rowsAffected == 0 {
		return todo.NewTodoInvalidIDError(id)
	}

	return nil
}

// unmarshalTodo tries to unmarshal the given slice of bytes to a Todo
// if this does not succeed, a new TodoHandlerError is returned.
func unmarshalTodo(bodyData []byte) (*todo.Todo, *todo.HandlerError) {
	var toDO *todo.Todo

	err := json.Unmarshal(bodyData, &toDO)
	if err != nil {
		return nil, todo.NewTodoHandlerError(
			fmt.Sprintf("body data was malformed %s", string(bodyData)), http.StatusBadRequest)
	}

	return toDO, nil
}

// NewTodoHandler returns a new TodoHandler.
func NewTodoHandler(repository repository.TodoRepository, logger *zap.Logger) TodoHandler {
	return &todoHandler{
		logger:     logger,
		repository: repository,
	}
}
