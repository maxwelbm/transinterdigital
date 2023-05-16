package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBDriver               string
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPassword             string
	DBName                 string
	KeySecret              string
	DBURL                  string
	PgAdminDefaultEmail    string
	PgAdminDefaultPassword string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		DBDriver:               getEnv("DB_DRIVER", ""),
		DBHost:                 getEnv("DB_HOST", ""),
		DBPort:                 getEnv("DB_PORT", ""),
		DBUser:                 getEnv("DB_USER", ""),
		DBPassword:             getEnv("DB_PASSWORD", ""),
		DBName:                 getEnv("DB_NAME", ""),
		KeySecret:              getEnv("KEY_SECRET", ""),
		DBURL:                  getEnv("DB_URL", ""),
		PgAdminDefaultEmail:    getEnv("PGADMIN_DEFAULT_EMAIL", ""),
		PgAdminDefaultPassword: getEnv("PGADMIN_DEFAULT_PASSWORD", ""),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
