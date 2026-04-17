package repository

import (
	"context"
	"practical-assessment/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Login interface {
	Login(c context.Context, db *gorm.DB, logger *logrus.Logger, userEmail string) (*model.User, error)
}

type loginRepository struct{}

func Newlogin() *loginRepository {
	return &loginRepository{}
}

func (l *loginRepository) Login(c context.Context, db *gorm.DB, logger *logrus.Logger, userEmail string) (*model.User, error) {
	var user model.User
	result := db.WithContext(c).Table("users").Where("email = ?", userEmail).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}

	logger.Info("Data fetched successfully")

	return &user, nil
}
