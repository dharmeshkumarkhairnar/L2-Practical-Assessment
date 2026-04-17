package repository

import (
	"context"
	"fmt"
	"practical-assessment/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type loginRepository struct {
	DB          *gorm.DB
	RedisClient *redis.Client
	// Logger      *logrus.Logger
}

type LoginRepository interface {
	UserLogin(ctx context.Context, userRequest model.LoginRequest) (model.Users, error)
}

func NewloginRepository(db *gorm.DB, redisClient *redis.Client) *loginRepository {
	return &loginRepository{
		DB:          db,
		RedisClient: redisClient,
	}
}

// will return accesstoken, refreshtoken, jti, and an error
func (repository *loginRepository) UserLogin(ctx context.Context, userRequest model.LoginRequest) (model.Users, error) {

	userFromDb := model.Users{}

	result := repository.DB.WithContext(ctx).Table("users").Where("email = ?", userRequest.Email).First(&userFromDb)
	if result.Error != nil {
		fmt.Println("error while getting user:", result.Error)

		if result.Error == gorm.ErrRecordNotFound {
			return userFromDb, fmt.Errorf("user not found")
		}
		return userFromDb, fmt.Errorf("database query error")
	}

	fmt.Println("user from database:", userFromDb)
	return userFromDb, nil
}
