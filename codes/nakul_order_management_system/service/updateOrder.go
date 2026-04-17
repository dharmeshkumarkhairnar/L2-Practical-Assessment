package service

import (
	"context"
	"practical-assessment/model"
	"practical-assessment/repository"

	"gorm.io/gorm"
)

type UpdateOrderSerivce struct {
	Repo repository.UpdateOrderRepo
	DB   *gorm.DB
}

func NewUpdateOrderSerivce(repo repository.UpdateOrderRepo, db *gorm.DB) *UpdateOrderSerivce {
	return &UpdateOrderSerivce{
		Repo: repo,
		DB:   db,
	}
}

func (service UpdateOrderSerivce) UpdateOrderService(ctx context.Context, updateReq model.UpdateOrderRequest, orderId uint, userId int64) error {

	order := model.Orders{
		UserID:      userId,
		ProductName: updateReq.ProductName,
		Quantity:    updateReq.Quantity,
		Price:       updateReq.Price,
		Status:      "CREATED",
	}

	return service.Repo.UpdateOrder(ctx, orderId, userId, order)
}
