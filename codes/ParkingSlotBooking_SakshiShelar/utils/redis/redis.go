package redis

import (
	"context"
	"fmt"
	"practical-assessment/constant"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	once        sync.Once
)

// InitRedis initializes the Redis client (runs only once)
func InitRedis() error {
	var err error

	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:         constant.RedisAddress,
			Password:     constant.RedisPassword, // "" if no password
			DB:           constant.RedisDb,       // default DB
			PoolSize:     constant.RedisPoolSize,
			MinIdleConns: constant.RedisMinIdleConns,
			DialTimeout:  constant.RedisDialTimeout,
			ReadTimeout:  constant.RedisReadTimeout,
			WriteTimeout: constant.RedisWriteTimeout,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Ping to verify connection
		_, err = redisClient.Ping(ctx).Result()
		if err != nil {
			// err = fmt.Errorf("Failed to connect tp redis: %v",err)
			err = fmt.Errorf(constant.RedisInitFailed, err)
			return
		}
	})
	return err
}

// GetRedisClient returns the initialized Redis client
func GetRedisClient() (*redis.Client, error) {
	err := InitRedis()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}
