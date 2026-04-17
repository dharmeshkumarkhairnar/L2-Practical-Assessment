package service

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LogoutService struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewLogoutService(db *gorm.DB, redisClient *redis.Client) *LogoutService {
	return &LogoutService{
		DB:          db,
		RedisClient: redisClient,
	}
}

func (service LogoutService) Logout(ctx context.Context, userId int64, token string) error {
	//on logout invalidating the token from redis
	key := token
	err := service.RedisClient.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("redis error: %v", err)
	}

	//deleting the session pair also, since it no longer exists
	sessionKey := fmt.Sprintf("userId:%d", userId)
	err = service.RedisClient.Del(ctx, sessionKey).Err()
	if err != nil {
		return fmt.Errorf("redis error: %v", err)
	}

	return nil
}
