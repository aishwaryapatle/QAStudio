package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	DBSSL  string

	JWTSecret string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", "password"),
		DBName: getEnv("DB_NAME", "test_mgmt"),
		DBSSL:  getEnv("DB_SSL", "disable"),

		JWTSecret: getEnv("JWT_SECRET", "supersecretkey"),
	}
}

func getEnv(key, defaultValue  string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue 
}