package service

import (
	"context"
	"errors"
	"practical-assessment/constant"
	"practical-assessment/repository"
	"practical-assessment/utils/database"
	"practical-assessment/utils/redis"
)

type CancelBookingsService struct {
	cancelBookingsRepo repository.CancelBookingsRepoInterface
}

func NewCancelBookingsService(cancelBookingsRepo repository.CancelBookingsRepoInterface) *CancelBookingsService {
	return &CancelBookingsService{
		cancelBookingsRepo: cancelBookingsRepo,
	}
}

func (service *CancelBookingsService) CancelBookings(ctx context.Context, userId int, bookingsId int) error {
	client := database.GetDB().DB

	bookings, err := service.cancelBookingsRepo.GetBookingByID(ctx, client, bookingsId)
	if err != nil {
		return errors.New(constant.BookingNotFound)
	}

	if bookings.UserId != userId {
		return errors.New(constant.UserUnauthorized)
	}

	if bookings.Status == constant.SlotsCancelled {
		return errors.New(constant.AlreadyCancelled)
	}

	bookings.Status = constant.SlotsCancelled

	err = service.cancelBookingsRepo.CancelUserBookings(ctx, client, bookings)
	if err != nil {
		return errors.New(constant.FailedToCancel)
	}

	redisClient, _ := redis.GetRedisClient()
	redisClient.Del(ctx, "available_slots")

	return nil
}
