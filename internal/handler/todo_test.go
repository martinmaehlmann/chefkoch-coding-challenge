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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, mockRepository := newHandler(t, ctrl)

	mockRepository.EXPECT().Create(gomock.Any()).Return(defaultTestReturnTodo())

	bodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	toDo, err := handler.Create(bodyData)
	assert.Nil(t, err)

	assert.Equal(t, defaultTestReturnTodo(), toDo)
}

func Test_todoHandler_CreateUnmarshalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, _ := newHandler(t, ctrl)

	bodyData, err := json.Marshal("invalid")
	assert.NoError(t, err)

	toDo, err := handler.Create(bodyData)

	assert.Error(t, err)
	assert.Nil(t, toDo)
	assert.Equal(t, todo.NewTodoHandlerError(fmt.Sprintf(
		"body data was malformed %s", string(bodyData)), http.StatusBadRequest), err)
}

func Test_todoHandler_CreateInvalidTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, _ := newHandler(t, ctrl)

	invalidTodo := defaultTestReturnTodo()
	invalidTodo.Name = ""

	bodyData, err := json.Marshal(invalidTodo)
	assert.NoError(t, err)

	toDo, err := handler.Create(bodyData)

	assert.Error(t, err)
	assert.Nil(t, toDo)
}

func Test_todoHandler_DeleteNoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, mockRepository := newHandler(t, ctrl)

	mockRepository.EXPECT().Delete(gomock.Any()).Return(int64(1))

	err := handler.Delete("1")
	assert.Nil(t, err)
}

func Test_todoHandler_DeleteInvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, _ := newHandler(t, ctrl)

	err := handler.Delete("-1")
	assert.Equal(t, todo.NewTodoInvalidIDError("-1"), err)
}

func Test_todoHandler_DeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, mockRepository := newHandler(t, ctrl)

	mockRepository.EXPECT().Delete(gomock.Any()).Return(int64(0))

	err := handler.Delete("1")
	assert.Equal(t, todo.NewTodoInvalidIDError("1"), err)
}

func Test_todoHandler_FindNoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, mockRepository := newHandler(t, ctrl)

	mockRepository.EXPECT().Find(gomock.Any()).Return(defaultTestReturnTodo())

	toDo, err := handler.Find("1")
	assert.Nil(t, err)

	assert.Equal(t, defaultTestReturnTodo(), toDo)
}

func Test_todoHandler_FindInvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, _ := newHandler(t, ctrl)

	toDo, err := handler.Find("-1")

	assert.Nil(t, toDo)
	assert.Equal(t, todo.NewTodoInvalidIDError("-1"), err)
}

func Test_todoHandler_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, mockRepository := newHandler(t, ctrl)

	returnValue := make([]*todo.Todo, 2)
	returnValue[0] = defaultTestReturnTodo()
	returnValue[1] = defaultTestReturnTodo()

	mockRepository.EXPECT().FindAll().Return(returnValue)

	toDos := handler.FindAll()
	assert.Equal(t, 2, len(toDos))

	for i := range returnValue {
		assert.Equal(t, returnValue[i], toDos[i])
	}
}

func Test_todoHandler_FindAllNoneExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, mockRepository := newHandler(t, ctrl)

	returnValue := make([]*todo.Todo, 0)

	mockRepository.EXPECT().FindAll().Return(returnValue)

	toDos := handler.FindAll()
	assert.Equal(t, 0, len(toDos))
}

func Test_todoHandler_UpdateNoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, mockRepository := newHandler(t, ctrl)

	bodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	mockRepository.EXPECT().Update(gomock.Any()).Return(defaultTestReturnTodo())

	toDo, err := handler.Update(bodyData, "1")
	assert.Nil(t, err)

	assert.Equal(t, defaultTestReturnTodo(), toDo)
}

func Test_todoHandler_UpdateInvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, _ := newHandler(t, ctrl)

	bodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	toDo, err := handler.Update(bodyData, "-1")
	assert.Nil(t, toDo)

	assert.Equal(t, todo.NewTodoInvalidIDError("-1"), err)
}

func Test_todoHandler_UpdateMalformedBodyData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, _ := newHandler(t, ctrl)

	bodyData, err := json.Marshal("invalid")
	assert.NoError(t, err)

	toDo, err := handler.Update(bodyData, "1")
	assert.Nil(t, toDo)

	assert.Equal(t, todo.NewTodoHandlerError(
		fmt.Sprintf("body data was malformed %s", string(bodyData)), http.StatusBadRequest), err)
}

func Test_todoHandler_UpdateInvalidTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, _ := newHandler(t, ctrl)

	invalidTodo := defaultTestReturnTodo()
	invalidTodo.Name = ""

	bodyData, err := json.Marshal(invalidTodo)
	assert.NoError(t, err)

	toDo, err := handler.Update(bodyData, "1")
	assert.Nil(t, toDo)

	assert.Equal(t, todo.NewInvalidTodo(invalidTodo), err)
}

func Test_todoHandler_UpdateEntryNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler, mockRepository := newHandler(t, ctrl)

	mockRepository.EXPECT().Update(gomock.Any()).Return(nil)

	bodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	toDo, err := handler.Update(bodyData, "1")
	assert.Nil(t, toDo)

	assert.Equal(t, todo.NewTodoHandlerError(
		fmt.Sprintf("todo %v does not exist", defaultTestReturnTodo()), http.StatusBadRequest), err)
}

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
