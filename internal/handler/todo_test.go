package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	repository_mock "gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/repository/mock"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/todo"
	"go.uber.org/zap"
)

func Test_todoHandler_CreateNoError(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler and mockRepository
	handler, mockRepository := newHandler(t, ctrl)

	// set the mocked function calls
	mockRepository.EXPECT().Create(gomock.Any()).Return(defaultTestReturnTodo())

	// create test data
	bodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	// create the todo
	toDo, err := handler.Create(bodyData)
	assert.Nil(t, err)

	// check the result
	assert.Equal(t, defaultTestReturnTodo(), toDo)
}

func Test_todoHandler_CreateUnmarshalError(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler
	handler, _ := newHandler(t, ctrl)

	// create test data
	bodyData, err := json.Marshal("invalid")
	assert.NoError(t, err)

	// create the todo
	toDo, err := handler.Create(bodyData)

	// check the result
	assert.Error(t, err)
	assert.Nil(t, toDo)
	assert.Equal(t, todo.NewTodoHandlerError(fmt.Sprintf(
		"body data was malformed %s", string(bodyData)), http.StatusBadRequest), err)
}

func Test_todoHandler_CreateInvalidTodo(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler
	handler, _ := newHandler(t, ctrl)

	// create test data
	invalidTodo := defaultTestReturnTodo()
	invalidTodo.Name = ""

	bodyData, err := json.Marshal(invalidTodo)
	assert.NoError(t, err)

	// create the todo
	toDo, err := handler.Create(bodyData)

	// check the result
	assert.Error(t, err)
	assert.Nil(t, toDo)
}

func Test_todoHandler_DeleteNoError(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler and mockRepository
	handler, mockRepository := newHandler(t, ctrl)

	// set the mocked function calls
	mockRepository.EXPECT().Delete(gomock.Any()).Return(int64(1))

	// call delete
	err := handler.Delete("1")
	assert.Nil(t, err)
}

func Test_todoHandler_DeleteInvalidID(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler
	handler, _ := newHandler(t, ctrl)

	// call delete
	err := handler.Delete("-1")
	assert.Equal(t, todo.NewTodoInvalidIDError("-1"), err)
}

func Test_todoHandler_DeleteNotFound(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler and mockRepository
	handler, mockRepository := newHandler(t, ctrl)

	// set the mocked function calls
	mockRepository.EXPECT().Delete(gomock.Any()).Return(int64(0))

	// call delete
	err := handler.Delete("1")
	assert.Equal(t, todo.NewTodoInvalidIDError("1"), err)
}

func Test_todoHandler_FindNoError(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler and mockRepository
	handler, mockRepository := newHandler(t, ctrl)

	// set the mocked function calls
	mockRepository.EXPECT().Find(gomock.Any()).Return(defaultTestReturnTodo())

	// call find
	toDo, err := handler.Find("1")
	assert.Nil(t, err)

	// check the result
	assert.Equal(t, defaultTestReturnTodo(), toDo)
}

func Test_todoHandler_FindInvalidID(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler
	handler, _ := newHandler(t, ctrl)

	// call find
	toDo, err := handler.Find("-1")

	// check the result
	assert.Nil(t, toDo)
	assert.Equal(t, todo.NewTodoInvalidIDError("-1"), err)
}

func Test_todoHandler_FindAll(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler and mockRepository
	handler, mockRepository := newHandler(t, ctrl)

	// create test data
	returnValue := make([]*todo.Todo, 2)
	returnValue[0] = defaultTestReturnTodo()
	returnValue[1] = defaultTestReturnTodo()

	// set the mocked function calls
	mockRepository.EXPECT().FindAll().Return(returnValue)

	// call findAll
	toDos := handler.FindAll()
	assert.Equal(t, 2, len(toDos))

	// check the result
	for i := range returnValue {
		assert.Equal(t, returnValue[i], toDos[i])
	}
}

func Test_todoHandler_FindAllNoneExist(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler and mockRepository
	handler, mockRepository := newHandler(t, ctrl)

	// create test data
	returnValue := make([]*todo.Todo, 0)

	// set the mocked function calls
	mockRepository.EXPECT().FindAll().Return(returnValue)

	// call findAll
	toDos := handler.FindAll()
	assert.Equal(t, 0, len(toDos))
}

func Test_todoHandler_UpdateNoError(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler and mockRepository
	handler, mockRepository := newHandler(t, ctrl)

	// create test data
	bodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	// set the mocked function calls
	mockRepository.EXPECT().Update(gomock.Any()).Return(defaultTestReturnTodo())

	// call update
	toDo, err := handler.Update(bodyData, "1")
	assert.Nil(t, err)

	// check the result
	assert.Equal(t, defaultTestReturnTodo(), toDo)
}

func Test_todoHandler_UpdateInvalidID(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler
	handler, _ := newHandler(t, ctrl)

	// create test data
	bodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	// call update
	toDo, err := handler.Update(bodyData, "-1")
	assert.Nil(t, toDo)

	// check the result
	assert.Equal(t, todo.NewTodoInvalidIDError("-1"), err)
}

func Test_todoHandler_UpdateMalformedBodyData(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler
	handler, _ := newHandler(t, ctrl)

	// create test data
	bodyData, err := json.Marshal("invalid")
	assert.NoError(t, err)

	// call update
	toDo, err := handler.Update(bodyData, "1")
	assert.Nil(t, toDo)

	// check the result
	assert.Equal(t, todo.NewTodoHandlerError(
		fmt.Sprintf("body data was malformed %s", string(bodyData)), http.StatusBadRequest), err)
}

func Test_todoHandler_UpdateInvalidTodo(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler
	handler, _ := newHandler(t, ctrl)

	// create test data
	invalidTodo := defaultTestReturnTodo()
	invalidTodo.Name = ""

	bodyData, err := json.Marshal(invalidTodo)
	assert.NoError(t, err)

	// call update
	toDo, err := handler.Update(bodyData, "1")
	assert.Nil(t, toDo)

	// check the result
	assert.Equal(t, todo.NewInvalidTodo(invalidTodo), err)
}

func Test_todoHandler_UpdateEntryNotFound(t *testing.T) {
	// get new test setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get a new handler and mockRepository
	handler, mockRepository := newHandler(t, ctrl)

	// set the mocked function calls
	mockRepository.EXPECT().Update(gomock.Any()).Return(nil)

	// create test data
	bodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	// call update
	toDo, err := handler.Update(bodyData, "1")
	assert.Nil(t, toDo)

	// check the result
	assert.Equal(t, todo.NewTodoHandlerError(
		fmt.Sprintf("todo %v does not exist", defaultTestReturnTodo()), http.StatusBadRequest), err)
}

// newHandler returns a new handler and a mocked repository.
func newHandler(t *testing.T, ctrl *gomock.Controller) (TodoHandler, *repository_mock.MockTodoRepository) {
	t.Helper()

	logger, err := zap.NewProduction()
	assert.NoError(t, err)

	mockRepository := repository_mock.NewMockTodoRepository(ctrl)

	return NewTodoHandler(mockRepository, logger), mockRepository
}

// defaultTestReturnTodo conviniernce function to return a stuct to test with.
func defaultTestReturnTodo() *todo.Todo {
	return &todo.Todo{
		ID:          1,
		Name:        "test",
		Description: "test",
		Tasks:       nil,
	}
}
