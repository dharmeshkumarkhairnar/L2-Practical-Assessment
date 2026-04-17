package service

import (
	"context"
	"errors"
	"practical-assessment/constant"
	"practical-assessment/utils/redis"
)

type LogoutUserService struct{}

func NewLogoutUserService() *LogoutUserService {
	return &LogoutUserService{}
}

func (service *LogoutUserService) LogoutUser(ctx context.Context, tokenString string) error {
	//get redisclient
	redisClient, err := redis.GetRedisClient()
	if redisClient == nil && err != nil {
		return errors.New(constant.RedisInitFailed)
	}

	err = redisClient.Del(ctx, tokenString).Err()
	if err != nil {
		return errors.New(constant.FailedToDelFromRedis)
	}

	return nil
}
