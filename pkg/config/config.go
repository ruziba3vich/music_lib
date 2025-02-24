package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBSSLMode   string
	ExternalAPI string
	RedisTTL    int
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default env variables")
	}

	redisTTL, _ := strconv.Atoi(getEnv("REDIS_TTL", "3600"))

	config := &Config{
		Port:        getEnv("PORT", "7777"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", ""),
		DBName:      getEnv("DB_NAME", "music_db"),
		DBSSLMode:   getEnv("DB_SSLMODE", "disable"),
		ExternalAPI: getEnv("EXTERNAL_API_URL", "http://localhost:8000/info"),
		RedisTTL:    redisTTL,
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
