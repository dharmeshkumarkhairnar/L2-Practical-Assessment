package repository

import (
	"context"
	"fmt"
	"practical-assessment/model"

	"gorm.io/gorm"
)

type getOrders struct {
	DB *gorm.DB
}

type GetOrders interface {
	GetUserOrder(ctx context.Context, userId int64) ([]model.GetOrderResponse, error)
}

func NewGetOrders(db *gorm.DB) *getOrders {
	return &getOrders{
		DB: db,
	}
}

func (repo getOrders) GetUserOrder(ctx context.Context, userId int64) ([]model.GetOrderResponse, error) {

	// orders := []model.Orders{}
	userORder := []model.GetOrderResponse{}

	err := repo.DB.WithContext(ctx).
		Table("orders").
		Where("user_id = ?", userId).
		Find(&userORder).Error

	if err != nil {
		fmt.Println("get db err:", err)
		return userORder, fmt.Errorf("database query error")
	}
	if len(userORder) > 0 {
		return userORder, nil
	}

	return userORder, fmt.Errorf("user does not have any orders")
}
