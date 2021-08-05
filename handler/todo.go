package handler

import (
	"OAuth/helper"
	"OAuth/todo"
	"OAuth/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	service todo.Service
}

func NewTodoHandler(service todo.Service) *todoHandler {
	return &todoHandler{service}
}

func(h *todoHandler) GetTodos(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User) //! get id ddari user yg login melalui jwt
	userID := currentUser.ID

	todos, err := h.service.GetTodos(userID)
	if err != nil {
		response := helper.ApiResponse("Error to get todos", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List of todos", http.StatusOK, "success", todo.FormatTodos(todos))
	
	c.JSON(http.StatusOK, response)
}

func(h *todoHandler) GetTodo(c *gin.Context) {

	var input todo.GetTodoDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(user.User) //! get id ddari user yg login melalui jwt
	userID := currentUser.ID

	todoDetail, err := h.service.GetTodoById(input, userID)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	} 

	response := helper.ApiResponse("Success get todo detail", http.StatusOK, "success", todo.FormatTodoDetail(todoDetail))
	c.JSON(http.StatusOK, response)
}

func(h *todoHandler) CreateTodo(c *gin.Context) {
	var input todo.CreateTodoInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ApiResponse("Failed Create Todo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTodo, err := h.service.CreateTodo(input)
	if err != nil {
		response := helper.ApiResponse("Failed Create Todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	response := helper.ApiResponse("Success Create Todo", http.StatusOK, "success", todo.FormatTodo(newTodo))

	c.JSON(http.StatusOK, response)
}

func(h *todoHandler) UpdateTodo(c *gin.Context) {

	var inputID todo.GetTodoDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to update Todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	var inputData todo.UpdateTodoInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ApiResponse("Failed Update Todo", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(user.User) //! melakukan auth user, hanya user yang memiliki item tsb bisa melakukabn update
	inputData.User = currentUser

	updateTodo, err := h.service.UpdateTodo(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed Update Todo	", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	response := helper.ApiResponse("Success Update Todo	", http.StatusOK, "success", todo.FormatTodo(updateTodo))

	c.JSON(http.StatusOK, response)
}

func(h *todoHandler) DeleteTodo(c *gin.Context) {
	var inputID todo.GetTodoDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to delete Todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(user.User) //! get id ddari user yg login melalui jwt
	userID := currentUser.ID

	deleteTodo, err := h.service.DeleteTodo(inputID, userID)
	if err != nil {
		response := helper.ApiResponse("Failed Delete Todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	response := helper.ApiResponse("Success Delete Todo", http.StatusOK, "success", todo.FormatTodo(deleteTodo))

	c.JSON(http.StatusOK, response)
}