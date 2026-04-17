package repository

import (
	"context"
	"practical-assessment/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GetTasks interface {
	Get(c context.Context, db *gorm.DB, logger *logrus.Logger, userId int64, status string, limit int64, offset int64) ([]*model.BFFGetTaskResponse, error)
}

type getTasksRepository struct{}

func NewGetTasks() *getTasksRepository {
	return &getTasksRepository{}
}

func (gt *getTasksRepository) Get(c context.Context, db *gorm.DB, logger *logrus.Logger, userId int64, status string, limit int64, offset int64) ([]*model.BFFGetTaskResponse, error) {
	var tasks []*model.BFFGetTaskResponse
	var result *gorm.DB

	if status=="" && limit==0 {
		result=db.WithContext(c).Table("tasks").Offset(int(offset)).Where("user_id = ? and not is_deleted", userId).Scan(&tasks)
	} else if status!="" && limit!=0 {
		result=db.WithContext(c).Table("tasks").Limit(int(limit)).Offset(int(offset)).Where("user_id = ? and not is_deleted and status = ?", userId,status).Scan(&tasks)
	}  else if status=="" && limit!=0 {
		result=db.WithContext(c).Table("tasks").Limit(int(limit)).Offset(int(offset)).Where("user_id = ? and not is_deleted", userId).Scan(&tasks)
	}  else if status!="" && limit==0 {
		result=db.WithContext(c).Table("tasks").Offset(int(offset)).Where("user_id = ? and not is_deleted and status = ?", userId,status).Scan(&tasks)
	}
	
	

	if result.Error!=nil {
		return nil,result.Error
	}

	if len(tasks)==0 {
		return nil, gorm.ErrRecordNotFound
	}

	logger.Info("Data fetched successfully")

	return tasks, nil
}
