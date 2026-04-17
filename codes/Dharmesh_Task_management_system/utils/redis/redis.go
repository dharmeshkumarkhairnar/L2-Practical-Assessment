package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	once        sync.Once
)

const (
	REDIS_URL      = "localhost:6379"
	REDIS_PASSWORD = ""
)

// InitRedis initializes the Redis client (runs only once)
func InitRedis() error {
	var err error

	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:         REDIS_URL,
			Password:     REDIS_PASSWORD, // "" if no password
			DB:           0,              // default DB
			PoolSize:     10,
			MinIdleConns: 5,
			DialTimeout:  5 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Ping to verify connection
		_, err = redisClient.Ping(ctx).Result()
		if err != nil {
			err = fmt.Errorf("failed to connect to redis: %w", err)
			return
		}
	})

	return err
}

// GetRedisClient returns the initialized Redis client
func GetRedisClient() *redis.Client {
	return redisClient
}
