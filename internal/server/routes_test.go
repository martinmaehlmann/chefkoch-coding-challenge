package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/config"
	handler_mock "gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/handler/mock"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/todo"
	"go.uber.org/zap"
)

func TestServer_createTodo(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)

	// create test data
	requestBodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	// setup mocked function calls
	handlerMock.EXPECT().Create(gomock.Eq(requestBodyData)).Return(defaultTestReturnTodo(), nil)

	// run http request
	// nolint:bodyclose // is closed in cleanup
	response, err := http.Post(fmt.Sprintf("%s/todos", ts.URL), "application/json", bytes.NewBuffer(requestBodyData))
	assert.NoError(t, err)

	defer cleanup(t, ctrl, ts, response)

	// check status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// read response body
	bodyData, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// check response
	var toDO *todo.Todo
	err = json.Unmarshal(bodyData, &toDO)
	assert.NoError(t, err)
	assert.Equal(t, defaultTestReturnTodo(), toDO)
}

func TestServer_createTodoInvalidTodo(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)

	// create test data
	invalidTodo := defaultTestReturnTodo()
	invalidTodo.Name = ""

	requestBodyData, err := json.Marshal(invalidTodo)
	assert.NoError(t, err)

	// setup mocked function calls
	handlerMock.EXPECT().Create(gomock.Eq(requestBodyData)).Return(nil, todo.NewInvalidTodo(invalidTodo))

	// run http request
	// nolint:bodyclose // is closed in cleanup
	response, err := http.Post(fmt.Sprintf("%s/todos", ts.URL), "application/json", bytes.NewBuffer(requestBodyData))
	assert.NoError(t, err)

	defer cleanup(t, ctrl, ts, response)

	// check status code
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	// read response body
	bodyData, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// check response
	message, err := json.Marshal(
		fmt.Sprintf("todo %v is not valid. The field name is required for tasks and todos.", invalidTodo))
	assert.NoError(t, err)

	assert.Equal(t, message, bodyData)
}

func TestServer_deleteTodo(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)

	// setup mocked function calls
	handlerMock.EXPECT().Delete(gomock.Eq("1")).Return(nil)

	// initialize http client
	client := &http.Client{}

	url := fmt.Sprintf("%s/todos/%d", ts.URL, 1)

	// setup http request
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	// run http request
	// nolint:bodyclose // is closed in cleanup
	response, err := client.Do(request)
	assert.NoError(t, err)

	defer cleanup(t, ctrl, ts, response)

	// check status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// read response body
	bodyData, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// setup expected response
	message, err := json.Marshal(fmt.Sprintf("deleted todo with id %s", "1"))
	assert.NoError(t, err)

	// check response
	assert.Equal(t, message, bodyData)
}

func TestServer_deleteTodoInvalidID(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)

	// setup mocked function calls
	handlerMock.EXPECT().Delete(gomock.Eq("-1")).Return(todo.NewTodoInvalidIDError("-1"))

	// initialize http client
	client := &http.Client{}

	url := fmt.Sprintf("%s/todos/%d", ts.URL, -1)

	// setup http request
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	// run http request
	// nolint:bodyclose // is closed in cleanup
	response, err := client.Do(request)
	assert.NoError(t, err)

	defer cleanup(t, ctrl, ts, response)

	// check status code
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	// read response body
	bodyData, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// setup expected response
	message, err := json.Marshal(fmt.Sprintf("%s is not a valid id. IDs are positive integers", "-1"))
	assert.NoError(t, err)

	// check response
	assert.Equal(t, message, bodyData)
}

func TestServer_findAllTodo(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)
	defer cleanup(t, ctrl, ts, nil)

	// setup mocked function calls
	handlerMock.EXPECT().FindAll().Return([]*todo.Todo{defaultTestReturnTodo()})

	// run http request
	// nolint:bodyclose // body is closed in cleanup
	response, err := http.Get(fmt.Sprintf("%s/todos", ts.URL))
	assert.NoError(t, err)

	// check status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// read response body
	bodyData, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// check response
	var toDO []*todo.Todo
	err = json.Unmarshal(bodyData, &toDO)
	assert.NoError(t, err)
	assert.Equal(t, []*todo.Todo{defaultTestReturnTodo()}, toDO)
}

func TestServer_findTodoNoError(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)

	// setup mocked function calls
	handlerMock.EXPECT().Find(gomock.Eq("1")).Return(defaultTestReturnTodo(), nil)

	// run http request
	// nolint:bodyclose // body is closed in cleanup
	response, err := http.Get(fmt.Sprintf("%s/todos/%s", ts.URL, "1"))
	assert.NoError(t, err)

	defer cleanup(t, ctrl, ts, response)

	// check status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// read response body
	bodyData, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// check response
	var toDo *todo.Todo
	err = json.Unmarshal(bodyData, &toDo)
	assert.NoError(t, err)

	expected := defaultTestReturnTodo()
	assert.Equal(t, expected, toDo)
}

func TestServer_findTodoInvalidID(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)

	// setup mocked function calls
	handlerMock.EXPECT().Find(gomock.Eq("-1")).Return(nil, todo.NewTodoInvalidIDError("-1"))

	// run http request
	// nolint:bodyclose // body is closed in cleanup
	response, err := http.Get(fmt.Sprintf("%s/todos/%s", ts.URL, "-1"))
	assert.NoError(t, err)

	defer cleanup(t, ctrl, ts, response)

	// check status code
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	// read response body
	bodyData, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// setup expected response
	message, err := json.Marshal("-1 is not a valid id. IDs are positive integers")
	assert.NoError(t, err)

	// check response
	assert.Equal(t, message, bodyData)
}

func TestServer_updateTodo(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)

	// initialize http client
	client := &http.Client{}

	// create test data
	bodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	// setup mocked function calls
	handlerMock.EXPECT().Update(gomock.Eq(bodyData), gomock.Eq("1")).Return(defaultTestReturnTodo(), nil)

	url := fmt.Sprintf("%s/todos/%d", ts.URL, 1)

	// setup http request
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bodyData))
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	// run http request
	// nolint:bodyclose // is closed in cleanup
	response, err := client.Do(request)
	assert.NoError(t, err)

	defer cleanup(t, ctrl, ts, response)

	// check status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// read response body
	bodyData, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// check response
	var toDo *todo.Todo
	err = json.Unmarshal(bodyData, &toDo)
	assert.NoError(t, err)

	expected := defaultTestReturnTodo()
	assert.Equal(t, expected, toDo)
}

func TestServer_updateTodoInvalidID(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)

	// initialize http client
	client := &http.Client{}

	// create test data
	bodyData, err := json.Marshal(defaultTestReturnTodo())
	assert.NoError(t, err)

	// setup mocked function calls
	handlerMock.EXPECT().Update(gomock.Eq(bodyData), gomock.Eq("-1")).Return(nil, todo.NewTodoInvalidIDError("-1"))

	url := fmt.Sprintf("%s/todos/%d", ts.URL, -1)

	// setup http request
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bodyData))
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	// run http request
	// nolint:bodyclose // is closed in cleanup
	response, err := client.Do(request)
	assert.NoError(t, err)

	defer cleanup(t, ctrl, ts, response)

	// check status code
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	// read response body
	bodyData, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// check response
	assert.Equal(t, "\"-1 is not a valid id. IDs are positive integers\"", string(bodyData))
}

func TestServer_updateTodoMalformedBody(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)

	// initialize http client
	client := &http.Client{}

	// create test data
	requestBodyData, err := json.Marshal("invalid todo data")
	assert.NoError(t, err)

	// setup mocked function calls
	handlerMock.EXPECT().Update(gomock.Eq(requestBodyData), gomock.Eq("1")).Return(
		nil, todo.NewTodoHandlerError(
			fmt.Sprintf("body data was malformed %s", string(requestBodyData)), http.StatusBadRequest))

	url := fmt.Sprintf("%s/todos/%d", ts.URL, 1)

	// setup http request
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBodyData))
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	// run http request
	// nolint:bodyclose // is closed in cleanup
	response, err := client.Do(request)
	assert.NoError(t, err)

	defer cleanup(t, ctrl, ts, response)

	// check status code
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	// read response body
	bodyData, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// setup expected response
	message, err := json.Marshal(fmt.Sprintf("body data was malformed %s", string(requestBodyData)))
	assert.NoError(t, err)

	// check response
	assert.Equal(t, message, bodyData)
}

func TestServer_updateTodoInvalidTodo(t *testing.T) {
	// get new test setup
	ctrl, ts, handlerMock := newTestServer(t)

	// initialize http client
	client := &http.Client{}

	// create test data
	toDo := defaultTestReturnTodo()
	toDo.Name = ""

	bodyData, err := json.Marshal(toDo)
	assert.NoError(t, err)

	// setup mocked function calls
	handlerMock.EXPECT().Update(gomock.Eq(bodyData), gomock.Eq("1")).Return(nil, todo.NewInvalidTodo(toDo))

	url := fmt.Sprintf("%s/todos/%d", ts.URL, 1)

	// setup http request
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bodyData))
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	// run http request
	// nolint:bodyclose // is closed in cleanup
	response, err := client.Do(request)
	assert.NoError(t, err)

	defer cleanup(t, ctrl, ts, response)

	// check status code
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	// read response body
	bodyData, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	// setup expected response
	message, err := json.Marshal(todo.NewInvalidTodo(toDo).Message)
	assert.NoError(t, err)

	// check response
	assert.Equal(t, message, bodyData)
}

// newTestServer returns a new test server with all routes and configuration set.
func newTestServer(t *testing.T) (*gomock.Controller, *httptest.Server, *handler_mock.MockTodoHandler) {
	t.Helper()
	ctrl := gomock.NewController(t)
	handlerMock := handler_mock.NewMockTodoHandler(ctrl)
	logger, _ := zap.NewProduction()
	registry := config.NewRegistry(&config.GinConfig{Port: randomPort()}, nil, logger)
	server := NewServer(registry, handlerMock, logger)
	server.addRoutes()
	ts := httptest.NewServer(server.engine)

	return ctrl, ts, handlerMock
}

// cleanup cleans up all resources used by the test.
func cleanup(t *testing.T, ctrl *gomock.Controller, ts *httptest.Server, response *http.Response) {
	t.Helper()

	if response != nil {
		err := response.Body.Close()
		assert.NoError(t, err)
	}

	ts.Close()
	ctrl.Finish()
}

// randomPort returns a random port to use for the server in the range of including 8000 - 8999.
func randomPort() int {
	rand.Seed(time.Now().Unix())

	// Generate a random number x where x is in range 5<=x<=20
	rangeLower := 8000
	rangeUpper := 8999

	return rangeLower + rand.Intn(rangeUpper-rangeLower+1)
}

func defaultTestReturnTodo() *todo.Todo {
	tasks := make([]*todo.Task, 2)
	tasks[0] = &todo.Task{
		ID:          1,
		Name:        "test1",
		Description: "test1",
	}
	tasks[1] = &todo.Task{
		ID:          2,
		Name:        "test2",
		Description: "test2",
	}

	return &todo.Todo{
		ID:          1,
		Name:        "test",
		Description: "test",
		Tasks:       tasks,
	}
}
