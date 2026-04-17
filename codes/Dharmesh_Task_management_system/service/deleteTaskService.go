package service

import (
	"context"
	"practical-assessment/repository"
	"practical-assessment/utils/database"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DeleteTaskService struct {
	deleteTasksRepository repository.DeleteTasks
}

func NewDeleteTaskService(DeleteTaskRepo repository.DeleteTasks) *DeleteTaskService {
	return &DeleteTaskService{
		deleteTasksRepository: DeleteTaskRepo,
	}
}

func (dt *DeleteTaskService) DeleteTasks(ctx *gin.Context, c context.Context, logger *logrus.Logger, taskId int64, userId int64) error {
	dbClient := database.GetDB().DB

	err := dt.deleteTasksRepository.Delete(c, dbClient, logger, taskId, userId)

	if err != nil {
		return err
	}

	return nil
}
