package service

import (
	"context"
	"practical-assessment/model"
	"practical-assessment/repository"
	"practical-assessment/utils/database"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateTaskService struct {
	createTaskRepository repository.CreateTask
}

func NewCreateTask(createTaskRepo repository.CreateTask) *CreateTaskService {
	return &CreateTaskService{
		createTaskRepository: createTaskRepo,
	}
}

func (cu *CreateTaskService) CreateTask(ctx *gin.Context, c context.Context, logger *logrus.Logger, bffCreateUserRequest model.BFFCreateTaskRequest,userId int64) error {
	dbClient := database.GetDB().DB

	err := cu.createTaskRepository.Create(c, dbClient, logger, bffCreateUserRequest,userId)

	if err != nil {
		return err
	}

	return nil
}
