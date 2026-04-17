package repository

import (
	"context"
	"practical-assessment/constant"
	"practical-assessment/model"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LoginUserRepositoryInterface interface {
	LoginUserDB(ctx context.Context, db *gorm.DB, bffLoginUserRequest model.BFFLoginUserRequest) (*model.Users, error)
}

type loginUserRepository struct{}

// the constructor
func LoginUserRepository() *loginUserRepository {
	return &loginUserRepository{}
}

// this function returns the entire user if it exists in db else error
func (user *loginUserRepository) LoginUserDB(ctx context.Context, db *gorm.DB, bffLoginUserRequest model.BFFLoginUserRequest) (*model.Users, error) {
	start := time.Now()
	logger := logrus.New()

	var PresentUser model.Users

	resultOfLogin := db.
		WithContext(ctx).
		Table("users").
		Where(constant.EmailField, bffLoginUserRequest.Email).
		First(&PresentUser)

	if resultOfLogin.Error != nil {
		return nil, resultOfLogin.Error
	}

	logger.WithFields(logrus.Fields{
		"user":    bffLoginUserRequest.Email,
		"latency": time.Since(start).Seconds(),
	}).Info(constant.UserLoginSuccess)

	return &PresentUser, nil
}

