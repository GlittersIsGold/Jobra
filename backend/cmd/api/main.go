package main

import (
//   "net/http"
	"fmt"
	"backend/config"
//   "github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()
	
	fmt.Printf("Port: %s\n", cfg.Port)
	fmt.Printf("Environment: %s\n", cfg.Environment) 
	fmt.Printf("Database URL: %s\n", cfg.DatabaseURL)
	fmt.Printf("JWT Expiry: %d hours\n", cfg.JWTExpireHours)

	//   // Create a Gin router with default middleware (logger and recovery)
//   r := gin.Default()

//   // Define a simple GET endpoint
//   r.GET("/ping", func(c *gin.Context) {
//     // Return JSON response
//     c.JSON(http.StatusOK, gin.H{
//       "message": "pong",
//     })
//   })

//   // Start server on port 8080 (default)
//   // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
//   r.Run()
}