package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	Type string

	Environment    string
	Port           string
	TrustedProxies []string
	DatabaseURL    string

	JWTSecret      string
	JWTExpireHours int
}

var (
	configFlag         = flag.String("configuration", "none", "Configuration mode (none|os-vars|env)")
	portFlag           = flag.Int("port", 8000, "Server port, 8000 by default, integer")
	envFlag            = flag.String("environment", "development", "Project environment mode (development|production)")
	dbUrlFlag          = flag.String("db-url", "postgresql://username:password@localhost:5432/jobra?sslmode=disable", "Database connection url, string")
	jwtSecretFlag      = flag.String("jwt-secret", "your-super-secret-jwt-key-change-in-production", "Auth JWT secret key, string")
	jwtExpiryFlag      = flag.Int("jwt-expire-hours", 1, "JWT expire time in hours, integer")
	trustedProxiesFlag = flag.String("trusted-proxies", "127.0.0.1,::1", "Enumeration of gin trusted proxies, as comma-sep string")
)

func Load() (Config, error) {
	logConfig()

	var config Config
	var err error

	switch *configFlag {
	case "none":
		log.Println("none configuration mode selected, retreiving configuration params from cli")
		config, err = ParseCliParams()
	case "os-var":
		log.Println("os-vars configuration mode selected, retreivn configuration params from OS variables")
		config, err = ParseOsVars()
	case "env":
		log.Println("env configuration mode selected, retreiving configuration params from .env")
		config, err = ParseEnv()
	default:
		log.Println("configuration flag is incorrect")
	}

	if err != nil {
		return Config{}, fmt.Errorf("failed to load config: %w", err)
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func logConfig() {
	log.Println("=== Loaded Configuration ===")
	log.Printf("Configuration: %s", *configFlag)
	log.Printf("Port: %d", *portFlag)
	log.Printf("Environment: %s", *envFlag)
	log.Printf("Database URL: %s", *dbUrlFlag)
	log.Printf("JWT secret: %s", *jwtSecretFlag)
	log.Printf("JWT: %d hour(s)", *jwtExpiryFlag)
	log.Printf("Trusted proxies: %#v\n", strings.Split(*trustedProxiesFlag, ","))
	log.Println("============================")
}

func ParseCliParams() (Config, error) {
	return Config, error
}

func ParseOsVars() (Config, error) {
	return Config, error
}

func ParseEnv() (Config, error) {
	return Config, error
}
