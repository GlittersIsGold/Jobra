package handlers

import (
	"backend/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine, 
	authHandler *AuthHandler, 
	userHandler *UserHandler, 
	jwtManager *auth.JWTManager, 
	orderHandler *OrderHandler,
	orderUserHandler *OrderUserHandler,
	) {
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	protected := router.Group("/api")
	protected.Use(auth.AuthMiddleware(jwtManager))
	{
		// User routes
		protected.GET("/profile", userHandler.GetProfile)
		protected.GET("/users", userHandler.GetUsers)
		protected.PUT("/profile", userHandler.UpdateProfile)
		protected.PUT("/profile/password", userHandler.ChangePassword)

		// Order routes
		protected.GET("/orders", orderHandler.GetUserOrders)
		protected.GET("/orders/all", orderHandler.GetAllOrders)
		protected.POST("/orders", orderHandler.CreateOrder)
		protected.GET("/orders/:id", orderHandler.GetOrder)
		protected.PUT("/orders/:id", orderHandler.UpdateOrder)
		protected.DELETE("/orders/:id", orderHandler.DeleteOrder)

		// Order to user routes
		protected.POST("/order-users", orderUserHandler.CreateOrderUser)
		protected.GET("/order-users", orderUserHandler.GetOrderUsers)
		protected.GET("/check-order-user", orderUserHandler.CheckOrderUserExists)
		protected.POST("/assign", orderUserHandler.AssignExecutor)
		protected.POST("/remove-executor", orderUserHandler.RemoveExecutor)
		protected.DELETE("/order-users/:orderUserId", orderUserHandler.DeleteOrderUser)
	}
}