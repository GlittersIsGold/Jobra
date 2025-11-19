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
	Port         string
	Environment  string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	DSN string
}

type JWTConfig struct {
	SecretKey string
	ExpiresIn time.Duration
}

type ProxyConfig struct {
	TrustedProxies []string
}

func Load() *Config {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-super-secret-key-change-in-production"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	return &Config{
		Server: ServerConfig{
			Port:         port,
			Environment:  env,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		Database: DatabaseConfig{
			DSN: "users.db",
		},
		JWT: JWTConfig{
			SecretKey: jwtSecret,
			ExpiresIn: 24 * time.Hour,
		},
		Proxy: ProxyConfig{
			TrustedProxies: []string{"127.0.0.1", "localhost"},
		},
	}
}