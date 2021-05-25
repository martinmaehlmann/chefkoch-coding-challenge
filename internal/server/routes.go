package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) findAllTodo(c *gin.Context) {
	todos := s.todoHandler.FindAll()
	c.JSON(http.StatusOK, todos)
}

func (s *Server) findTodo(c *gin.Context) {
	id := c.Param("id")

	toDo, err := s.todoHandler.Find(id)
	if err != nil {
		c.JSON(err.HTTPCode, err.Message)

		return
	}

	c.JSON(http.StatusOK, toDo)
}

func (s *Server) createTodo(c *gin.Context) {
	bodyData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s.logger.Error(fmt.Sprintf("could not read body data: %v", err))
	}

	result, serviceError := s.todoHandler.Create(bodyData)
	if serviceError != nil {
		c.JSON(serviceError.HTTPCode, serviceError.Message)

		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *Server) updateTodo(c *gin.Context) {
	id := c.Param("id")

	bodyData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s.logger.Error(fmt.Sprintf("could not read body data: %v", err))
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	result, serviceError := s.todoHandler.Update(bodyData, id)
	if serviceError != nil {
		c.JSON(serviceError.HTTPCode, serviceError.Message)

		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *Server) deleteTodo(c *gin.Context) {
	id := c.Param("id")

	serviceError := s.todoHandler.Delete(id)
	if serviceError != nil {
		c.JSON(serviceError.HTTPCode, serviceError.Message)

		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("deleted todo with id %s", id))
}
