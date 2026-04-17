package service

import (
	"context"
	"practical-assessment/repository"

	"gorm.io/gorm"
)

type DelOrderSerivce struct {
	Repo repository.DelOrderRepo
	DB   *gorm.DB
}

func NewDelOrderService(repo repository.DelOrderRepo, db *gorm.DB) *DelOrderSerivce {
	return &DelOrderSerivce{
		Repo: repo,
		DB:   db,
	}
}

func (service DelOrderSerivce) DeleteOrder(ctx context.Context, orderID uint, userId int64) error {

	return service.Repo.DeleteOrder(ctx, orderID, userId)
}
