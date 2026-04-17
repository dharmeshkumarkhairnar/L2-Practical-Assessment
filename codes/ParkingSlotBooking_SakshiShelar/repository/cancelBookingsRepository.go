package repository

import (
	"context"
	"practical-assessment/constant"
	"practical-assessment/model"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CancelBookingsRepoInterface interface {
	GetBookingByID(ctx context.Context, db *gorm.DB, id int) (*model.Bookings, error)
	CancelUserBookings(ctx context.Context, db *gorm.DB, booking *model.Bookings) error
}

type cancelBookingsRepository struct{}

func CancelBookingsRepository() *cancelBookingsRepository {
	return &cancelBookingsRepository{}
}

func (user *cancelBookingsRepository) GetBookingByID(ctx context.Context, db *gorm.DB, id int) (*model.Bookings, error) {

	var booking model.Bookings

	err := db.
		WithContext(ctx).
		Table(constant.BookingsTable).
		First(&booking, id).
		Error

	if err != nil {
		return nil, err
	}

	return &booking, err
}

func (user *cancelBookingsRepository) CancelUserBookings(ctx context.Context, db *gorm.DB, booking *model.Bookings) error {
	start := time.Now()
	logger := logrus.New()

	err := db.
		WithContext(ctx).
		Table(constant.BookingsTable).
		Save(booking).
		Error

	if err != nil {
		return err
	}

	logger.WithFields(logrus.Fields{
		"user":    booking.Id,
		"latency": time.Since(start).Seconds(),
	}).Info(constant.BookingCancelledSuccess)

	return nil
}
