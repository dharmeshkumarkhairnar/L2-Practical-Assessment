package repository

import (
	"context"
	"practical-assessment/constant"
	"practical-assessment/model"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AvailableSlotsRepositoryInterface interface {
	GetAvailableSlotsDB(ctx context.Context, db *gorm.DB) ([]int, error)
}

type availableSlotsRepository struct{}

func AvailableSlotsRepository() *availableSlotsRepository {
	return &availableSlotsRepository{}
}

func (user *availableSlotsRepository) GetAvailableSlotsDB(ctx context.Context, db *gorm.DB) ([]int, error) {
	start := time.Now()
	logger := logrus.New()

	var availableSlots []model.Slots

	err := db.
		WithContext(ctx).
		Table(constant.SlotsTable).
		Where(constant.StatusField, constant.SlotFree).
		Find(&availableSlots).
		Error

	if err != nil {
		return nil, err
	}

	var slotNumber []int
	for _, val := range availableSlots {
		slotNumber = append(slotNumber, val.SlotNumber)
	}

	logger.WithFields(logrus.Fields{
		"latency": time.Since(start).Seconds(),
	}).Info(constant.SLotsFetchedSuccess)

	return slotNumber, nil
}
