package service

import (
	"context"
	"practical-assessment/model"
	"practical-assessment/repository"

	"gorm.io/gorm"
)

type GetOrder struct {
	Repo repository.GetOrders
	DB   *gorm.DB
}

func NewGetOrder(repo repository.GetOrders, db *gorm.DB) *GetOrder {
	return &GetOrder{
		Repo: repo,
		DB:   db,
	}
}

func (service GetOrder) GetOrder(ctx context.Context, userId int64) ([]model.GetOrderResponse, error) {
	orders, err := service.Repo.GetUserOrder(ctx, userId)
	return orders, err
}
