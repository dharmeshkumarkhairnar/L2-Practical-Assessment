package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"practical-assessment/constant"
	"practical-assessment/repository"
	"practical-assessment/utils"
	"practical-assessment/utils/database"
	"practical-assessment/utils/redis"
	"time"
)

type AvailableSlotsService struct {
	availableSlotsRepo repository.AvailableSlotsRepositoryInterface
}

func NewAvailableSlotsService(availableSlotsRepo repository.AvailableSlotsRepositoryInterface) *AvailableSlotsService {
	return &AvailableSlotsService{
		availableSlotsRepo: availableSlotsRepo,
	}
}

func (service *AvailableSlotsService) AvailableSlots(ctx context.Context, tokenString string) ([]int, error) {
	//first getting data from redis
	redisClient, err := redis.GetRedisClient()
	if redisClient == nil && err != nil {
		return nil, errors.New(constant.RedisInitFailed)
	}

	data, err := redisClient.Get(ctx, constant.AvailableSlotsKey).Result()
	if err == nil && len(data) > 0 {
		var slotsRedis []int
		err := json.Unmarshal([]byte(data), &slotsRedis)
		if err != nil {
			return nil, errors.New(constant.JsonUnmarshalFailed)
		}
		fmt.Println("!!!!!!!!!!!!!!Data from Redis!!!!!!!!!!")
		return slotsRedis, nil
	}

	//if not in redis fetch from db
	client := database.GetDB().DB

	slotsDB, err := service.availableSlotsRepo.GetAvailableSlotsDB(ctx, client)
	if err != nil {
		return nil, errors.New(constant.NoFreeSlots)
	}

	jsonData, err := json.Marshal(slotsDB)
	if err != nil {
		return nil, errors.New(constant.JsonMarshalFailed)
	}

	tokenExpiry, _ := utils.GetExpiryFromToken(tokenString)
	ttl := time.Until(time.Unix(tokenExpiry, 0))
	
	err = redisClient.Set(ctx, constant.AvailableSlotsKey, jsonData, ttl).Err()
	if err != nil {
		return nil, errors.New(constant.FailedToSetInRedis)
	}

	fmt.Println("!!!!!!!!!!!!!!Data from DB!!!!!!!!!!")
	return slotsDB, nil

}
