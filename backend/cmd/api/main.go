package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"backend/config"
	"backend/internal/auth"
	"backend/internal/handlers"
	"backend/internal/repository/postgres"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()

	pool, err := pgxpool.New(context.Background(), cfg.Database.URL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	if cfg.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	userRepo := postgres.NewUserRepository(pool)

	jwtManager := auth.NewJWTManager(cfg.JWT.SecretKey, cfg.JWT.ExpiresIn)
	
	authService := auth.NewAuthService(userRepo, jwtManager)
	authHandler := handlers.NewAuthHandler(authService)
	
	userHandler := handlers.NewUserHandler(userRepo)

	orderRepo := postgres.NewOrderRepository(pool)
	orderHandler := handlers.NewOrderHandler(orderRepo)

	orderUserRepo := postgres.NewOrderUserRepository(pool)
	orderUserHandler := handlers.NewOrderUserHandler(orderUserRepo, orderRepo)

	router := gin.New()
	router.SetTrustedProxies(cfg.Proxy.TrustedProxies)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	handlers.SetupRoutes(
		router, 
		authHandler, 
		userHandler, 
		jwtManager, 
		orderHandler, 
		orderUserHandler,
	)

	server := &http.Server{
		Addr:         cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	log.Printf("Starting server on %s in %s mode", cfg.Server.Port, cfg.Server.Environment)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}