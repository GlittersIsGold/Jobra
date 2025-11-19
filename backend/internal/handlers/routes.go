package handlers

import (
	"backend/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authHandler *AuthHandler, userHandler *UserHandler, jwtManager *auth.JWTManager) {
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	protected := router.Group("/api")
	protected.Use(auth.AuthMiddleware(jwtManager))
	{
		protected.GET("/profile", userHandler.GetProfile)
		protected.GET("/users", userHandler.GetUsers)
	}
}