package repository

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UpdateTasks interface {
	Update(c context.Context, db *gorm.DB, logger *logrus.Logger, taskId int64, updateReq map[string]interface{}, userId int64) error
}

type updateTasksRepository struct{}

func NewUpdateTasks() *updateTasksRepository {
	return &updateTasksRepository{}
}

func (ut *updateTasksRepository) Update(c context.Context, db *gorm.DB, logger *logrus.Logger, taskId int64, updateReq map[string]interface{},userId int64) error {

	var count int64
	checkExistance := db.WithContext(c).Table("tasks").Where("id = ? and user_id = ? and not is_deleted", taskId,userId).Count(&count)

	if checkExistance.Error != nil {
		return checkExistance.Error
	}

	if count == 0 || checkExistance.Error == gorm.ErrRecordNotFound {
		return errors.New("tasks not found")
	}

	result := db.WithContext(c).Table("tasks").Where("id = ?", taskId).Updates(updateReq)

	if result.Error != nil {
		return result.Error
	}

	logger.Info("Data updated successfully")

	return nil
}
