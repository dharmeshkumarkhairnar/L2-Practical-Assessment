package repository

import (
	"context"
	"errors"
	"practical-assessment/model"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CreateTask interface {
	Create(c context.Context, db *gorm.DB, logger *logrus.Logger, userReq model.BFFCreateTaskRequest, userId int64) error
}

type createTaskRepository struct{}

func NewCreateTask() *createTaskRepository {
	return &createTaskRepository{}
}

func (ct *createTaskRepository) Create(c context.Context, db *gorm.DB, logger *logrus.Logger, userReq model.BFFCreateTaskRequest, userId int64) error {

	user := model.Tasks{
		UserID:      uint64(userId),
		Title:       strings.TrimSpace(userReq.Title),
		Description: strings.TrimSpace(userReq.Description),
		Status:      strings.ToLower(userReq.Status),
		Priority:    strings.ToLower(userReq.Priority),
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}

	result := db.WithContext(c).Table("tasks").Create(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("error in adding data in DB")
	}

	logger.Info("Data added successfully")

	return nil
}
