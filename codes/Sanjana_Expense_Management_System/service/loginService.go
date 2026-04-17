package service

import (
	"context"
	"errors"
	"log"
	"practical-assessment/model"
	"practical-assessment/repository"
	"practical-assessment/utils"
	"practical-assessment/utils/database"
	"practical-assessment/utils/redis"
	"time"
)

type LoginService struct {
	loginRepository repository.LoginRepository
}

func NewLoginService(loginRepository repository.LoginRepository) *LoginService {
	return &LoginService{loginRepository: loginRepository}
}

func (service *LoginService) LoginUserService(ctx context.Context, spanCtx context.Context, bffRequest model.BFFLoginRequest) (string, error) {
	postgresClient := database.GetDB().DB
	userDB, err := service.loginRepository.LoginUserRepository(ctx, postgresClient, bffRequest)
	if err != nil {
		return "", errors.New("Service Error")
	}

	userId := userDB.Id
	match := utils.CompareHashPassword(userDB.Password, bffRequest.Password)
	if !match {
		return "", errors.New("Password Mismatched, Login Failed")
	}

	token, _, err := utils.GenerateToken(userDB.Name)
	if err != nil {
		return "", errors.New("Error while generating token")
	}

	redisClient := redis.GetRedisClient()
	if redisClient == nil {
		return "", errors.New("Redis Connection Failed")
	}

	expiry, _ := utils.GetExpiry(token)

	ttl := time.Until(time.Unix(expiry, 0))

	err = redisClient.Set(ctx, token, userId, ttl).Err()
	if err != nil {
		return "", errors.New("Redis Set error")
	}
	// return token, nil
	// log.Print("Service works")
	// cacheKey := fmt.Sprintf("JWT_TOKEN:%s", token)
	// redisClient := redis.GetRedisClient()

	// log.Print(redisClient)
	// if redisClient != nil {
	// 	log.Print("Redis")
	// 	redisData, err := redisClient.Set(ctx, cacheKey, userId, ttl).Result()
	// 	log.Print(redisData)
	// 	if err != nil {
	// 		return errors.New("Redis set operation error")
	// 	}
	// }
	log.Print("Service Done")
	return token, nil
}
