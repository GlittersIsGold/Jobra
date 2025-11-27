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
	
	user := protected.Group("/user")
	{
		user.GET("", userHandler.GetProfile)
		user.PUT("", userHandler.UpdateProfile)
		user.PUT("/password", userHandler.ChangePassword)
	}

	users := protected.Group("/users")
	{	 
		users.GET("", userHandler.GetUsers) 
	}
	
	orders := protected.Group("/orders")
	{
		orders.GET("", orderHandler.GetAllOrders)
		orders.GET("/mine", orderHandler.GetUserOrders)
		orders.POST("/create", orderHandler.CreateOrder)

		order := orders.Group("/:id")
		{
			order.GET("", orderHandler.GetOrder)
			order.PUT("", orderHandler.UpdateOrder)
			order.DELETE("", orderHandler.DeleteOrder)
		}

		users := orders.Group("/users")
		{
			users.GET("", orderUserHandler.GetOrderUsers)
			users.POST("", orderUserHandler.CreateOrderUser)
			users.DELETE("/:orderUserId", orderUserHandler.DeleteOrderUser)
			users.GET("/check", orderUserHandler.CheckOrderUserExists)

			assign := users.Group("/assign")
			{
				assign.POST("", orderUserHandler.AssignExecutor)
				assign.DELETE("", orderUserHandler.RemoveExecutor)
			}
		}
	}
}