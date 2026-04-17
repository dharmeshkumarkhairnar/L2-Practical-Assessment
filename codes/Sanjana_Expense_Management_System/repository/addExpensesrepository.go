package repository

import (
	"context"
	"errors"
	"log"
	"practical-assessment/constant"
	"practical-assessment/model"

	"gorm.io/gorm"
)

type AddExpenseRepository interface {
	AddUserExpenseRepository(ctx context.Context, db *gorm.DB, bffRequest model.BFFAddExpenseRequest) (*model.Expenses, error)
}

type addExpenseRepository struct{}

func NewAddExpenseRepository() *addExpenseRepository {
	return &addExpenseRepository{}
}

func (repository *addExpenseRepository) AddUserExpenseRepository(ctx context.Context, db *gorm.DB, bffRequest model.BFFAddExpenseRequest) (*model.Expenses, error) {
	var expenseDB model.Expenses

	var userDB model.Users

	result1 := db.WithContext(ctx).Table(constant.TableUsers).Where(constant.FieldName, bffRequest.Name).First(&userDB)
	if result1.Error != nil {
		return nil, errors.New("Query Failed")
	}

	result2 := db.WithContext(ctx).Table(constant.TableExpenses).Where("uid = ? and category = ?", userDB.Id, bffRequest.Category).First(&expenseDB)
	if result2.Error != nil {
		return nil, errors.New("Query Failed")
	}
	log.Print("Repo works")
	log.Print(expenseDB)
	return &expenseDB, nil
}
