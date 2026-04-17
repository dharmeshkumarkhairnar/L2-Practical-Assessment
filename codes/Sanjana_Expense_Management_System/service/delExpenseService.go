package service

import (
	"context"
	"errors"
	"log"
	"practical-assessment/model"
	"practical-assessment/repository"
	"practical-assessment/utils/database"
)

type DelExpenseService struct {
	delExpenseRepository repository.DelExpenseRepository
}

func NewDelExpenseService(delExpenseRepository repository.DelExpenseRepository) *DelExpenseService {
	return &DelExpenseService{delExpenseRepository: delExpenseRepository}
}

func (service *DelExpenseService) DelUserExpenseService(ctx context.Context, spanCtx context.Context, bffRequest model.BFFDelExpenseRequest) error {
	postgresClient := database.GetDB().DB
	err := service.delExpenseRepository.DelUserExpenseRepository(ctx, postgresClient, bffRequest)
	if err != nil {
		return errors.New("No data Found")
	}

	log.Print("Service Done")
	return nil
}
