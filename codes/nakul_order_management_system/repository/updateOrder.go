package repository

import (
	"context"
	"fmt"
	"practical-assessment/model"

	"gorm.io/gorm"
)

type updateOrderRepo struct {
	DB *gorm.DB
}

type UpdateOrderRepo interface {
	UpdateOrder(ctx context.Context, orderId uint, userId int64, updateOrder model.Orders) error
}

func NewUpdateOrderRepo(db *gorm.DB) *updateOrderRepo {
	return &updateOrderRepo{
		DB: db,
	}
}

func (repo updateOrderRepo) UpdateOrder(ctx context.Context, orderId uint, userId int64, updateOrder model.Orders) error {
	result := repo.DB.WithContext(ctx).
		Table("orders").
		Where("id = ? AND user_id = ?", orderId, userId).
		Updates(&updateOrder)

	if result.Error != nil {
		return fmt.Errorf("database query error")
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("order does not exist")
	}

	return nil
}

/*err := db.WithContext(ctx).
	Table("orders").
	Where("id = ? AND user_id = ?", orderID, userID).
	Updates(map[string]interface{}{
		"product_name": productName,
		"quantity": quantity,
		"price": price,
	}).Error
if err != nil {
	return err
}*/
