package service

import (
	"context"
	"practical-assessment/model"
	"practical-assessment/repository"

	"gorm.io/gorm"
)

type CreateOrder struct {
	Repo repository.CreateOrderRepo
	DB   *gorm.DB
}

func NewCreateOrder(repo repository.CreateOrderRepo, db *gorm.DB) *CreateOrder {
	return &CreateOrder{
		Repo: repo,
		DB:   db,
	}
}

func (service CreateOrder) CreateOrderService(ctx context.Context, req model.CreateOrderRequest, userId int64) error {

	order := model.Orders{
		UserID:      userId,
		ProductName: req.ProductName,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Status:      "CREATED",
	}

	err := service.Repo.CreateUserOrder(ctx, order, userId)
	return err
}
