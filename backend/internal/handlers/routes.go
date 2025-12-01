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

	api := router.Group("/api")
	api.Use(auth.AuthMiddleware(jwtManager))
	
	user := api.Group("/user")
	{
		user.GET("", userHandler.GetProfile)
		user.PUT("", userHandler.UpdateProfile)
		user.PUT("/password", userHandler.ChangePassword)
	}

	users := api.Group("/users")
	{	 
		users.GET("", userHandler.GetUsers) 
	}
	
	assign := api.Group("/assign")
	{
		assign.POST("", orderUserHandler.AssignExecutor)
		assign.DELETE("", orderUserHandler.RemoveExecutor)
	}
	
	orders := api.Group("/orders")
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
		}
	}
}