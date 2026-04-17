package service

import (
	"context"
	"errors"
	"fmt"
	"practical-assessment/model"
	"practical-assessment/repository"
	"practical-assessment/utils"
	"practical-assessment/utils/database"
	"practical-assessment/utils/redis"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LoginService struct {
	loginRepository repository.Login
}

func NewUserlogin(loginRepo repository.Login) *LoginService {
	return &LoginService{
		loginRepository: loginRepo,
	}
}

func (uLogin *LoginService) UserLogin(ctx *gin.Context, c context.Context, logger *logrus.Logger, bffLoginRequest model.BFFLoginRequest) (string, error) {
	dbClient := database.GetDB().DB
	redisClient := redis.GetRedisClient()

	user, err := uLogin.loginRepository.Login(c, dbClient, logger, bffLoginRequest.Email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("user not found")
		}
		return "", err
	}

	matched := utils.CompareHashedPassword(user.Password, bffLoginRequest.Password)

	if !matched {
		return "", errors.New("password is incorrect")
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", errors.New("token generation failed")
	}
	cacheKey := fmt.Sprintf("ACTIVE_TOKEN_%d", user.ID)

	err = redisClient.Set(c, cacheKey, token, time.Duration(3*time.Hour)).Err()

	return token, nil
}
