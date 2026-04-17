package service

import (
	"context"
	"practical-assessment/model"
	"practical-assessment/repository"
	"practical-assessment/utils/database"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UpdateTaskService struct {
	updateTasksRepository repository.UpdateTasks
}

func NewUpdateTaskService(updateTasksRepository repository.UpdateTasks) *UpdateTaskService {
	return &UpdateTaskService{
		updateTasksRepository: updateTasksRepository,
	}
}

func (ut *UpdateTaskService) UpdateTasks(ctx *gin.Context, c context.Context, logger *logrus.Logger, taskId int64, bffUpdateTaskRequest model.BFFUpdateTaskRequest, userId int64) error {
	dbClient := database.GetDB().DB

	receivedData := map[string]interface{}{
		"title":    strings.TrimSpace(bffUpdateTaskRequest.Title),
		"priority": bffUpdateTaskRequest.Priority,
	}

	if bffUpdateTaskRequest.Status != "" {
		receivedData["status"] = bffUpdateTaskRequest.Status
	}

	if bffUpdateTaskRequest.Description != "" {
		receivedData["description"] = strings.TrimSpace(bffUpdateTaskRequest.Description)
	}

	err := ut.updateTasksRepository.Update(c, dbClient, logger, taskId, receivedData, userId)

	if err != nil {
		return err
	}

	return nil
}
