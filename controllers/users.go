package controllers

import (
	"backend-crowd-funding/auth"
	"backend-crowd-funding/helper"
	"backend-crowd-funding/users"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userControllers struct {
	userService users.Service
	authService auth.Service
}

func NewUserControllers(userService users.Service, authService auth.Service) *userControllers {
	return &userControllers{userService, authService}
}

func (h *userControllers) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(newUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userControllers) Login(c *gin.Context) {
	var input users.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Successfuly Loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userControllers) CheckEmailExists(c *gin.Context) {
	var input users.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Checking email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailExists, err := h.userService.IsEmailExists(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Internal Error"}

		response := helper.APIResponse("Checking email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailExists,
	}

	var metaMessage string
	metaMessage = "Email Exists"

	if IsEmailExists {
		metaMessage = "Email Avaliable"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusUnprocessableEntity, response)
}

func (h *userControllers) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed upload avatar", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.Users)
	userID := currentUser.ID
	path := fmt.Sprintf("images/avatar/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed upload avatar", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	_, err = h.userService.SaveAvatar(userID, path)

	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed upload avatar", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}

	response := helper.APIResponse("Avatar succesfuly upload", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userControllers) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.Users)
	formatter := users.FormatUser(currentUser, "")
	response := helper.APIResponse("Successfully fetch user", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
