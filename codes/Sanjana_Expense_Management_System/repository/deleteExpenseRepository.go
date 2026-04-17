package repository

import (
	"context"
	"errors"
	"log"
	"practical-assessment/constant"
	"practical-assessment/model"

	"gorm.io/gorm"
)

type DelExpenseRepository interface {
	DelUserExpenseRepository(ctx context.Context, db *gorm.DB, bffRequest model.BFFDelExpenseRequest) error
}

type delExpenseRepository struct{}

func NewDelExpenseRepository() *delExpenseRepository {
	return &delExpenseRepository{}
}

func (repository *delExpenseRepository) DelUserExpenseRepository(ctx context.Context, db *gorm.DB, bffRequest model.BFFDelExpenseRequest) error {
	var expenseDB model.Expenses

	var userDB model.Users

	result1 := db.WithContext(ctx).Table(constant.TableUsers).Where(constant.FieldName, bffRequest.Name).First(&userDB)
	if result1.Error != nil {
		return errors.New("No data found")
	}

	result2 := db.WithContext(ctx).Table(constant.TableExpenses).Where("uid = ? and category = ?", userDB.Id, bffRequest.Category).Update("activity", "inactive")
	if result2.Error != nil {
		return errors.New("Delete Unseuccessful")
	}
	log.Print("Repo works")
	log.Print(expenseDB)
	return nil
}
