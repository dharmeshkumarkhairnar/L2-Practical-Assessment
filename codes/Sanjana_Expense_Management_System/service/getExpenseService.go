package service

import (
	"context"
	"errors"
	"log"
	"practical-assessment/model"
	"practical-assessment/repository"
	"practical-assessment/utils/database"
)

type GetExpenseService struct {
	getExpenseRepository repository.GetExpenseRepository
}

func NewGetExpenseService(getExpenseRepository repository.GetExpenseRepository) *GetExpenseService {
	return &GetExpenseService{getExpenseRepository: getExpenseRepository}
}

func (service *GetExpenseService) GetUserExpenseService(ctx context.Context, spanCtx context.Context, bffRequest model.BFFGetExpenseRequest) (*model.Expenses, error) {
	postgresClient := database.GetDB().DB
	expenseDB, err := service.getExpenseRepository.GetUserExpenseRepository(ctx, postgresClient, bffRequest)
	if err != nil {
		return nil, errors.New("No data Found")
	}

	log.Print("Service Done")
	return expenseDB, nil
}
