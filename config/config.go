package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DatabaseURL string

	// Redis
	// RedisAddr     string
	// RedisPassword string
	// RedisDB       int

	// Kafka
	// KafkaBrokers  []string
	// KafkaClientID string
}

func New() *Config {
	// Load .env file
	_ = godotenv.Load()

	// PostgreSQL
	dbUser := getEnv("DB_USER", "user")
	dbPass := getEnv("DB_PASSWORD", "password")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "dbname")
	serverPort := getEnv("SERVER_PORT", "8080")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	// Redis
	// redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
	// redisPassword := getEnv("REDIS_PASSWORD", "")
	// redisDB := getEnvInt("REDIS_DB", 0)

	// Kafka
	// kafkaBrokers := strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ",")
	// kafkaClientID := getEnv("KAFKA_CLIENT_ID", "pt-xyz-app")

	return &Config{
		ServerPort:  serverPort,
		DatabaseURL: dbURL,

		// RedisAddr:     redisAddr,
		// RedisPassword: redisPassword,
		// RedisDB:       redisDB,

		// KafkaBrokers:  kafkaBrokers,
		// KafkaClientID: kafkaClientID,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// func getEnvInt(key string, fallback int) int {
// 	if value, exists := os.LookupEnv(key); exists {
// 		var i int
// 		_, err := fmt.Sscanf(value, "%d", &i)
// 		if err == nil {
// 			return i
// 		}
// 	}
// 	return fallback
// }
