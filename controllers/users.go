package controllers

import (
	"backend-crowd-funding/helper"
	"backend-crowd-funding/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userControllers struct {
	userService users.Service
}

func NewUserControllers(userService users.Service) *userControllers {
	return &userControllers{userService}
}

func (h *userControllers) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(newUser, "")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}
