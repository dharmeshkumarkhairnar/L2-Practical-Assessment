package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DeleteTasks interface {
	Delete(c context.Context, db *gorm.DB, logger *logrus.Logger, taskId int64, userId int64) error
}

type deleteTasksRepository struct{}

func NewDeleteTasks() *deleteTasksRepository {
	return &deleteTasksRepository{}
}

func (dt *deleteTasksRepository) Delete(c context.Context, db *gorm.DB, logger *logrus.Logger, taskId int64,userId int64) error {

	result := db.WithContext(c).Table("tasks").Where("id = ? and user_id = ?", taskId,userId).Update("is_deleted","true")

	if result.Error != nil {
		return result.Error
	}

	logger.Info("Data deleted successfully")

	return nil
}
