package handler

import (
	"OAuth/auth"
	"OAuth/helper"
	"OAuth/user"
	"net/http"

	"github.com/gin-gonic/gin"
)


type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService} //! passing userService
}

func(h *userHandler) RegisterUser(c *gin.Context) {
	/*
		tangkap input dari user
	 	map input dari user ke struct RegisterUserInput
	 	struct di atas passing sebagai parameter service
	*/

	var input user.RegisterUserInput
	
	err := c.ShouldBindJSON(&input) //! validasi di lakukan di sini

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage) //! entity tidak bisa di proses 

		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, _ := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}


func(h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login account failed", http.StatusUnprocessableEntity, "error", errorMessage) //! entity tidak bisa di proses 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login account failed", http.StatusUnprocessableEntity, "error", errorMessage) //! entity tidak bisa di proses, karena format email salah, id dan email tidak di temukan dan password salah
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.ApiResponse("Login account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.ApiResponse("Login has been Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}


func(h *userHandler) CheckEmailHasBeenRegister(c *gin.Context) {

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage) //! entity tidak bisa di proses 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.ApiResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage) //! entity tidak bisa di proses 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func(h *userHandler) FetchUser(c *gin.Context) { //! mengambil data user yang sekarang sedang login
	currentUser := c.MustGet("currentUser").(user.User)
	formatter := user.FormatFetchUser(currentUser)

	response := helper.ApiResponse("Fetch user data successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}