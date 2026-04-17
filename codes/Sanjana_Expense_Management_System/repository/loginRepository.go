package repository

import (
	"context"
	"errors"
	"log"
	"practical-assessment/constant"
	"practical-assessment/model"

	"gorm.io/gorm"
)

type LoginRepository interface {
	LoginUserRepository(ctx context.Context, db *gorm.DB, bffRequest model.BFFLoginRequest) (*model.Users, error)
}

type loginRepository struct{}

func NewRepository() *loginRepository {
	return &loginRepository{}
}

func (repository *loginRepository) LoginUserRepository(ctx context.Context, db *gorm.DB, bffRequest model.BFFLoginRequest) (*model.Users, error) {
	var userDB model.Users

	result := db.WithContext(ctx).Table(constant.TableUsers).Where(constant.FieldName, bffRequest.Username).First(&userDB)
	if result.Error != nil {
		return nil, errors.New("Query Failed")
	}
	log.Print("Repo works")
	log.Print(userDB)
	return &userDB, nil
}
