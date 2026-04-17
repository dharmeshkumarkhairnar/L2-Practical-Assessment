package repository

import (
	"context"
	"errors"
	"log"
	"practical-assessment/constant"
	"practical-assessment/model"

	"gorm.io/gorm"
)

type GetExpenseRepository interface {
	GetUserExpenseRepository(ctx context.Context, db *gorm.DB, bffRequest model.BFFGetExpenseRequest) (*model.Expenses, error)
}

type getExpenseRepository struct{}

func NewGetExpenseRepository() *getExpenseRepository {
	return &getExpenseRepository{}
}

func (repository *getExpenseRepository) GetUserExpenseRepository(ctx context.Context, db *gorm.DB, bffRequest model.BFFGetExpenseRequest) (*model.Expenses, error) {
	var expenseDB model.Expenses

	var userDB model.Users

	result1 := db.WithContext(ctx).Table(constant.TableUsers).Where(constant.FieldName, bffRequest.Name).First(&userDB)
	if result1.Error != nil {
		return nil, errors.New("No Data Found")
	}

	result2 := db.WithContext(ctx).Table(constant.TableExpenses).Where("uid = ? and category = ? and activity = ?", userDB.Id, bffRequest.Category, "active").First(&expenseDB)
	if result2.Error != nil {
		return nil, errors.New("No data Found")
	}
	log.Print("Repo works")
	log.Print(expenseDB)
	return &expenseDB, nil
}
