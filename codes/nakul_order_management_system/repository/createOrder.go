package repository

import (
	"context"
	"fmt"
	"practical-assessment/model"

	"gorm.io/gorm"
)

type createOrderRepo struct {
	DB *gorm.DB
}

type CreateOrderRepo interface {
	CreateUserOrder(ctx context.Context, order model.Orders, userId int64) error
}

func NewCreateOrderRepo(db *gorm.DB) *createOrderRepo {
	return &createOrderRepo{
		DB: db,
	}
}

func (repo *createOrderRepo) CreateUserOrder(ctx context.Context, order model.Orders, userId int64) error {

	result := repo.DB.WithContext(ctx).
		Table("orders").
		Create(&order)

	if result.Error != nil {
		fmt.Println("create query error:", result.Error)
		return fmt.Errorf("database query error")
	}
	return nil
}
