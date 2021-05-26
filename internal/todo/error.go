package todo

import (
	"fmt"
	"net/http"
)

// HandlerError an error to wrap errors in the handler.TodoHandler.
type HandlerError struct {
	Message  string
	HTTPCode int
}

// Error the error message string.
func (s HandlerError) Error() string {
	return s.Message
}

// NewTodoHandlerError returns a new HandlerError.
func NewTodoHandlerError(message string, httpCode int) *HandlerError {
	return &HandlerError{
		Message:  message,
		HTTPCode: httpCode,
	}
}

// NewTodoInvalidIDError returns an invalid id error
// Error Message: "%s is not a valid id. IDs are positive integers"
// HTTPCode: http.StatusBadRequest.
func NewTodoInvalidIDError(id string) *HandlerError {
	return &HandlerError{
		Message:  fmt.Sprintf("%s is not a valid id. IDs are positive integers", id),
		HTTPCode: http.StatusBadRequest,
	}
}

// NewInvalidTodo returns an invalid todo error
// Error Message: todo %v is not valid. The field name is required for tasks and todos.
// HTTPCode: http.StatusBadRequest.
func NewInvalidTodo(toDo *Todo) *HandlerError {
	return &HandlerError{
		Message:  fmt.Sprintf("todo %v is not valid. The field name is required for tasks and todos.", toDo),
		HTTPCode: http.StatusBadRequest,
	}
}
