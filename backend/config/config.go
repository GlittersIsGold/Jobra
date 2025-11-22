package config

import (
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Proxy    ProxyConfig
}

type ServerConfig struct {
	Port        string
	Environment string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout time.Duration
}

type DatabaseConfig struct {
	URL string
}

type JWTConfig struct {
	SecretKey string
	ExpiresIn time.Duration
}

type ProxyConfig struct {
	TrustedProxies []string
}

func Load() Config {
	return Config{
		Server: ServerConfig{
			Port:        getEnv("PORT", ":8080"),
			Environment: getEnv("ENVIRONMENT", "development"),
			ReadTimeout: 15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout: 60 * time.Second,
		},
		Database: DatabaseConfig{
			URL: getEnv("DATABASE_URL", "postgresql://user:password@localhost:5432/backend?sslmode=disable"),
		},
		JWT: JWTConfig{
			SecretKey: getEnv("JWT_SECRET", "your-secret-key"),
			ExpiresIn: 24 * time.Hour,
		},
		Proxy: ProxyConfig{
			TrustedProxies: []string{"127.0.0.1"},
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}