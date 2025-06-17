package pkg

// import (
// 	"context"
// 	"log"

// 	"pt-xyz-multifinance/config"

// 	"github.com/redis/go-redis/v9"
// )

// func NewRedis(cfg *config.Config) *redis.Client {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     cfg.RedisAddr,
// 		Password: cfg.RedisPassword,
// 		DB:       cfg.RedisDB,
// 	})

// 	if err := rdb.Ping(context.Background()).Err(); err != nil {
// 		log.Fatalf("Failed to connect to Redis: %v", err)
// 	}
// 	log.Println("Redis connection success")
// 	return rdb
// }
