package service

import (
	"context"
	"errors"
	"practical-assessment/constant"
	"practical-assessment/model"
	"practical-assessment/repository"
	"practical-assessment/utils"
	"practical-assessment/utils/database"
	"practical-assessment/utils/redis"
	"time"
)

type LoginUserService struct {
	loginUserRepository repository.LoginUserRepositoryInterface
}

func NewLoginUserService(loginUserRepository repository.LoginUserRepositoryInterface) *LoginUserService {
	return &LoginUserService{
		loginUserRepository: loginUserRepository,
	}
}

func (service *LoginUserService) LoginUser(ctx context.Context, spanctx context.Context, bffLoginUserService model.BFFLoginUserRequest) (string, error) {
	client := database.GetDB().DB

	user, err := service.loginUserRepository.LoginUserDB(spanctx, client, bffLoginUserService)
	if err != nil {
		return "", errors.New(constant.UserNotFound)
	}


	passwordMatchError := utils.CompareWithHashPassword(user.Password, bffLoginUserService.Password)
	if passwordMatchError != nil {
		return "", errors.New(constant.PasswordMismatch)
	}

	tokenString, err := utils.GenerateToken(bffLoginUserService.Email, "Login")
	if err != nil {
		return "", errors.New(constant.TokenGenerationFailed)
	}

	redisClient, err := redis.GetRedisClient()
	if redisClient == nil && err != nil {
		return "", errors.New(constant.RedisInitFailed)
	}

	tokenExpiry, _ := utils.GetExpiryFromToken(tokenString)
	ttl := time.Until(time.Unix(tokenExpiry, 0))

	err = redisClient.Set(ctx, tokenString, user.Id, ttl).Err()
	if err != nil {
		return "", errors.New(constant.FailedToSetInRedis)
	}
	return tokenString, nil
}
