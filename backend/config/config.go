package config

import (
	"os"
	"log"
	"flag"
)

type Config struct {
	Port        	 string
	Environment 	 string
	
	DatabaseURL      string
	
	JWTSecret   	 string
	JWTExpireHours   string 

	TrustedProxies	 string
}

var (
	portFlag           = flag.String("port", "8000", "Server port")
	configFlag         = flag.String("conf", "", "Name of the configuration file")
	envFlag            = flag.String("env", "", "Project environment mode (development or production)")
	dbURLFlag          = flag.String("db-url", "", "Database connection string url")
	jwtSecretFlag      = flag.String("jwt-secret", "", "Auth JWT secret key")
	jwtExpiryFlag      = flag.Int("jwt-expiry", 0, "JWT expire time in hours")
	trustedProxiesFlag = flag.String("trusted-proxies", "", "Enumeration of gin trusted proxies")
)

func Load() *Config {

	flag.Parse()

	return &Config{
		Port:           getEnv("PORT", "8081"),
		Environment: 	getEnv("ENVIRONMENT", "developmend"),
		
		DatabaseURL: 	getEnv("DATABASE_URL", "postgresql://postgres:password@localhost:5432/jobra?sslmode=disable"),
		
		JWTSecret: 		getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-in-production"),
		JWTExpireHours: getEnv("JWT_EXPIRY_HOURS", "24"), 

		TrustedProxies: getEnv("CORS_ALLOWED_ORIGINS", "*"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func logConfig(cfg *Config) {
	log.Println("=== Loaded Configuration ===")
	log.Printf("Port: %s", cfg.Port)
	log.Printf("Environment: %s", cfg.Environment)
	log.Printf("Database URL: %s", cfg.DatabaseURL)
	log.Printf("JWT Expiry: %d hours", cfg.JWTExpireHours)
	log.Printf("CORS Origins: %s", cfg.TrustedProxies)
	log.Println("============================")
}