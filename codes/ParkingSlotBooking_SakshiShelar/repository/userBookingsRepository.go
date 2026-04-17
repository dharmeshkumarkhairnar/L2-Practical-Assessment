package repository

import (
	"context"
	"practical-assessment/constant"
	"practical-assessment/model"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserBookingsRepoInterface interface {
	GetUserBookings(ctx context.Context, db *gorm.DB, userId int) ([]model.Bookings, error)
}

type userBookingsRepository struct{}

func UserBookingsRepository() *userBookingsRepository {
	return &userBookingsRepository{}
}

func (user *userBookingsRepository) GetUserBookings(ctx context.Context, db *gorm.DB, userId int) ([]model.Bookings, error) {
	start := time.Now()
	logger := logrus.New()

	var bookings []model.Bookings

	err := db.
		WithContext(ctx).
		Table("bookings").
		Where("user_id= ?", userId).
		Find(&bookings).
		Error

	if err != nil {
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"user":    userId,
		"latency": time.Since(start).Seconds(),
	}).Info(constant.UserLoginSuccess)

	return bookings, nil
}
