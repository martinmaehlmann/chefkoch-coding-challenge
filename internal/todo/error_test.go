package todo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewInvalidTodo(t *testing.T) {
	type args struct {
		toDo Todo
	}
	tests := []struct {
		name string
		args args
		want *HandlerError
	}{
		{
			name: "test provider",
			args: args{toDo: Todo{}},
			want: &HandlerError{
				Message:  fmt.Sprintf("todo %v is not valid", &Todo{}),
				HTTPCode: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInvalidTodo(&tt.args.toDo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvalidTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTodoHandlerError(t *testing.T) {
	type args struct {
		message  string
		httpCode int
	}
	tests := []struct {
		name string
		args args
		want *HandlerError
	}{
		{
			name: "test provider",
			args: args{
				message:  "test",
				httpCode: 1,
			},
			want: &HandlerError{
				Message:  "test",
				HTTPCode: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTodoHandlerError(tt.args.message, tt.args.httpCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoHandlerError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTodoInvalidIDError(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *HandlerError
	}{
		{
			name: "test provider",
			args: args{id: "-1"},
			want: &HandlerError{
				Message:  "-1 is not a valid id. IDs are positive integers",
				HTTPCode: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTodoInvalidIDError(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoInvalidIDError() = %v, want %v", got, tt.want)
			}
		})
	}
}
