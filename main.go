package main

import (
	"backend-crowd-funding/auth"
	"backend-crowd-funding/campaign"
	"backend-crowd-funding/controllers"
	"backend-crowd-funding/helper"
	"backend-crowd-funding/transaction"
	"backend-crowd-funding/users"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

	// Repository
	userRepository := users.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	// service
	userService := users.NewService(userRepository)
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transaction.NewService(transactionRepository, campaignRepository)

	// controllers
	userControllers := controllers.NewUserControllers(userService, authService)
	campaignControllers := controllers.NewCampaignControllers(campaignService)
	transactionsController := controllers.NewTransactionController(transactionService)

	router := gin.Default()
	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	/**
	* Auth Sevices
	 */
	api.POST("/auth/register", userControllers.RegisterUser)
	api.POST("/auth/sessions", userControllers.Login)
	api.POST("/auth/check-email", userControllers.CheckEmailExists)

	/*
	* User Services
	 */
	api.POST("/services/users/avatar", authMiddleWare(authService, userService), userControllers.UploadAvatar)

	/*
	* Campaigns Service
	 */
	api.GET("/services/campaigns", campaignControllers.GetCampaigns)
	api.GET("/services/campaigns/:id", campaignControllers.GetCampaign)
	api.POST("/services/campaigns", authMiddleWare(authService, userService), campaignControllers.CreateCampaign)
	api.PUT("/services/campaigns/:id", authMiddleWare(authService, userService), campaignControllers.UpdateCampaign)
	api.POST("/services/campaign-images", authMiddleWare(authService, userService), campaignControllers.UploadImage)

	/*
	* Transaction Campaigns
	 */
	api.GET("/services/campaigns/:id/transactions", authMiddleWare(authService, userService), transactionsController.GetCampaignTransactions)
	api.GET("/services/transactions", authMiddleWare(authService, userService), transactionsController.GetUserTransactions)

	router.Run(":8080")
}

func authMiddleWare(authService auth.Service, userService users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			tokenString = arrToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))
		user, err := userService.GetuserByID(userId)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}

}
