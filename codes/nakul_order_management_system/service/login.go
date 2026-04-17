package service

import (
	"context"
	"fmt"
	"practical-assessment/model"
	"practical-assessment/repository"
	"practical-assessment/utils"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LoginService struct {
	Repo        repository.LoginRepository
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewLoginService(repo repository.LoginRepository, db *gorm.DB, redisClient *redis.Client) *LoginService {
	return &LoginService{
		Repo:        repo,
		DB:          db,
		RedisClient: redisClient,
	}
}

func (Service *LoginService) Login(ctx context.Context, req model.LoginRequest) (string, string, error) {

	user, err := Service.Repo.UserLogin(ctx, req)
	if err != nil {
		fmt.Println("db error:", err)
		return "", "", err
	}

	if !utils.CompareHashPassword(user.Password, req.Password) {
		return "", "", fmt.Errorf("incorrect password")
	}

	//call create jwt, return the jwt to user and save it in redis
	accessToken, _, err := utils.CreateToken(user.Id)
	if err != nil {
		fmt.Println("token generation error:", err)
		return "", "", fmt.Errorf("failed to generate token")
	}

	//caching valid token in redis, deleted once expired or loggedout
	key := accessToken
	err = Service.RedisClient.Set(ctx, key, user.Id, 12*time.Hour).Err()
	if err != nil {
		return "", "", fmt.Errorf("error saving session in redis")
	}

	//for concurrent session control - a pair of userId:token is saved in redis, until it is valid (not expired)
	sessionKey := fmt.Sprintf("userId:%d", user.Id)
	err = Service.RedisClient.Set(ctx, sessionKey, accessToken, 12*time.Hour).Err()
	if err != nil {
		return "", "", fmt.Errorf("error saving session in redis")
	}

	return accessToken, "", nil
}
