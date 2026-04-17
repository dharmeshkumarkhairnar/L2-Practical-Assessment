package repository

import (
	"context"
	"fmt"
	"practical-assessment/model"

	"gorm.io/gorm"
)

type delOrderRepo struct {
	DB *gorm.DB
}

type DelOrderRepo interface {
	DeleteOrder(ctx context.Context, orderId uint, userId int64) error
}

func NewdelOrderRepo(db *gorm.DB) *delOrderRepo {
	return &delOrderRepo{
		DB: db,
	}
}

func (repo *delOrderRepo) DeleteOrder(ctx context.Context, orderId uint, userId int64) error {
	userOrder := model.Orders{}
	result := repo.DB.WithContext(ctx).
		Table("orders").
		Where("id = ? AND user_id = ?", orderId, userId).
		Scan(&userOrder)

	if result.Error != nil {
		fmt.Println("db error: ", result.Error)
		return fmt.Errorf("database query error")
	}
	if userOrder.Status == "CANCELLED" {
		return fmt.Errorf("order does not exist")
	}

	result = repo.DB.WithContext(ctx).
		Table("orders").
		Where("id = ? AND user_id = ?", orderId, userId).
		Update("status", "CANCELLED")

	if result.Error != nil {
		fmt.Println("db error: ", result.Error)
		return fmt.Errorf("database query error")
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("order not found")
	}

	return nil
}
