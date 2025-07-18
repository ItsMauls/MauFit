package app

import (
	"user-service/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, userHandler *handler.UserHandler) {
	apiV1 := router.Group("api/v1")

	{
		userRoutes := apiV1.Group("/users")
		{
			userRoutes.POST("", userHandler.CreateUser)
			userRoutes.GET("/:id", userHandler.GetUserByID)
			userRoutes.GET("", userHandler.GetAllUsers)
			userRoutes.POST("/register", userHandler.Register)
			userRoutes.POST("/login", userHandler.Login)
			userRoutes.GET("/verify-token", userHandler.VerifyToken)
		}
	}
}
