package service

import (
	"context"
	"errors"
	"practical-assessment/constant"
	"practical-assessment/model"
	"practical-assessment/repository"
	"practical-assessment/utils/database"
)

type UserBookingsService struct {
	userBookingsRepo repository.UserBookingsRepoInterface
}

func NewUserBookingsService(userBookingsRepo repository.UserBookingsRepoInterface) *UserBookingsService {
	return &UserBookingsService{
		userBookingsRepo: userBookingsRepo,
	}
}

func (service *UserBookingsService) UserBookings(ctx context.Context, userId int) ([]model.Bookings, error) {
	client := database.GetDB().DB

	bookings, err := service.userBookingsRepo.GetUserBookings(ctx, client, userId)
	if err != nil {
		return nil, errors.New(constant.NoBookingsFound)
	}

	return bookings,nil
}
