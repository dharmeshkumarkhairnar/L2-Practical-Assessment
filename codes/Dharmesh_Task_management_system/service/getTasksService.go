package service

import (
	"context"
	"errors"
	"practical-assessment/model"
	"practical-assessment/repository"
	"practical-assessment/utils/database"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GetTaskService struct {
	getTaskRepository repository.GetTasks
}

func NewGetTaskService(getTaskRepo repository.GetTasks) *GetTaskService {
	return &GetTaskService{
		getTaskRepository: getTaskRepo,
	}
}

func (gt *GetTaskService) GetTasks(ctx *gin.Context, c context.Context, logger *logrus.Logger, userId int64, status string, limit int64, offset int64) ([]*model.BFFGetTaskResponse, error) {
	dbClient := database.GetDB().DB

	tasks, err := gt.getTaskRepository.Get(c, dbClient, logger, userId, status,limit, offset)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("tasks not found")
		}
		return nil, err
	}

	return tasks, nil
}
