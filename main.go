package main

import (
	"backend-crowd-funding/controllers"
	"backend-crowd-funding/users"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userControllers := controllers.NewUserControllers(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	/**
	* Auth Sevices
	 */
	api.POST("/auth/register", userControllers.RegisterUser)
	api.POST("/auth/sessions", userControllers.Login)
	api.POST("/auth/check-email", userControllers.CheckEmailExists)

	router.Run()
}
