package main

import (
	"backend/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.SetTrustedProxies(config.TrustedProxies)

	gin.SetMode(config.Environment)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(config.Port)
}
