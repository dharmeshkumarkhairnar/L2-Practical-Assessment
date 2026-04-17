package service

import (
	"context"
	"errors"
	"fmt"
	"practical-assessment/utils/redis"

	"github.com/sirupsen/logrus"
)

type LogoutService struct{}

func NewUserLogout() *LogoutService {
	return &LogoutService{}
}

func (uLogout *LogoutService) UserLogout(c context.Context, logger *logrus.Logger, userID int64) error {
	redisClient := redis.GetRedisClient()

	cacheKey := fmt.Sprintf("ACTIVE_TOKEN_%d", userID)

	err := redisClient.Expire(c, cacheKey, 0).Err()

	if err != nil {
		return errors.New("error in expiring the data in radis")
	}

	logger.Info("token removed in the redis")

	return nil
}
