package config

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Type string 			`validate:"required,oneof=cli os-vars env"`

	Environment    string   `validate:"required,oneof=debug release test"`
	Port           string   `validate:"required,min=1,max=65535"`
	TrustedProxies []string `validate:"dive,ip"`
	DatabaseURL    string   `validate:"required,url"`

	JWTSecret      string   `validate:"required,min=32"`
	JWTExpireHours int      `validate:"required,min=1,max=8760"`
}

var (
	typeFlag           = flag.String("configuration", "cli", "Configuration mode defines the way to take application work params (cli|os-vars|env)")
	portFlag           = flag.Int("port", 8000, "Server port, 8000 by default, integer")
	envFlag            = flag.String("environment", "debug", "Project environment mode (development|production)")
	dbUrlFlag          = flag.String("db-url", "postgresql://username:password@localhost:5432/jobra?sslmode=disable", "Database connection url, string")
	jwtSecretFlag      = flag.String("jwt-secret", "your-super-secret-jwt-key-change-in-production", "Auth JWT secret key, string")
	jwtExpireFlag      = flag.Int("jwt-expire-hours", 1, "JWT expire time in hours, integer")
	trustedProxiesFlag = flag.String("trusted-proxies", "127.0.0.1,::1", "Enumeration of gin trusted proxies, as comma-sep string")
)

func Load() (Config, error) {
	logConfig()

	var config *Config
	var err error

	switch *typeFlag {
	case "cli":
		config, err = ParseCliParams()
	case "os-var":
		config, err = ParseOsVars()
	case "env":
		config, err = ParseEnv()
	default:
		log.Println("configuration flag is incorrect")
	}

	if err != nil {
		return Config{}, fmt.Errorf("failed to load config: %w", err)
	}

	return *config, nil
}

func logConfig() {
	log.Println("=== Loaded Configuration ===")
	log.Printf("Configuration: %s", *typeFlag)
	log.Printf("Port: %d", *portFlag)
	log.Printf("Environment: %s", *envFlag)
	log.Printf("Database URL: %s", *dbUrlFlag)
	log.Printf("JWT secret: %s", *jwtSecretFlag)
	log.Printf("JWT: %d hour(s)", *jwtExpireFlag)
	log.Printf("Trusted proxies: %#v\n", strings.Split(*trustedProxiesFlag, ","))
	log.Println("============================")
}

var validate = validator.New()

func Validate(c *Config) error {
	return validate.Struct(c)
}

func ParseCliParams() (*Config, error) {
	log.Println("cli configuration mode selected, retreiving configuration params from cli")

	config := &Config{
		Type:           *typeFlag,
		Environment:    *envFlag,
		Port:           strconv.Itoa(*portFlag),
		TrustedProxies: strings.Split(*trustedProxiesFlag, ","),
		DatabaseURL:    *dbUrlFlag,
		JWTSecret:      *jwtSecretFlag,
		JWTExpireHours: *jwtExpireFlag,
	}

	if err := Validate(config); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return config, nil
}

func ParseOsVars() (*Config, error) {
	log.Println("os-vars configuration mode selected, retreivng configuration params from OS variables")

	portStr := getEnv("PORT", "8000")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT: %w", err)
	}

	jwtExpireStr := getEnv("JWT_EXPIRE_HOURS", "1")
	jwtExpire, err := strconv.Atoi(jwtExpireStr)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRE_HOURS: %w", err)
	}

	config := &Config{
		Type:           getEnv("CONFIGURATION", "os-vars"),
		Environment:    getEnv("ENVIRONMENT", "debug"),
		Port:           strconv.Itoa(port),
		TrustedProxies: strings.Split(getEnv("TRUSTED_PROXIES", "127.0.0.1,::1"), ","),
		DatabaseURL:    getEnv("DATABASE_URL", "postgresql://username:password@localhost:5432/jobra?sslmode=disable"),
		JWTSecret:      getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-in-production"),
		JWTExpireHours: jwtExpire,
	}

	if err := validate.Struct(config); err != nil {
		return nil, fmt.Errorf("environment variables validation failed: %w", err)
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func ParseEnv() (*Config, error) {
	log.Println("env configuration mode selected, retreiving configuration params from .env")

	envVars, err := loadEnvFile(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	portStr := envVars["PORT"]
	if portStr == "" {
		return nil, fmt.Errorf("PORT is required in .env file")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT in .env file: %w", err)
	}

	jwtExpireStr := envVars["JWT_EXPIRE_HOURS"]
	if jwtExpireStr == "" {
		return nil, fmt.Errorf("JWT_EXPIRE_HOURS is required in .env file")
	}
	jwtExpire, err := strconv.Atoi(jwtExpireStr)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRE_HOURS in .env file: %w", err)
	}

	config := &Config{
		Type:           envVars["CONFIGURATION"],
		Environment:    envVars["ENVIRONMENT"],
		Port:           strconv.Itoa(port),
		TrustedProxies: strings.Split(envVars["TRUSTED_PROXIES"], ","),
		DatabaseURL:    envVars["DATABASE_URL"],
		JWTSecret:      envVars["JWT_SECRET"],
		JWTExpireHours: jwtExpire,
	}

	return config, nil
}

func loadEnvFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	envVars := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		value = strings.Trim(value, `"'`)

		envVars[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envVars, nil
}
