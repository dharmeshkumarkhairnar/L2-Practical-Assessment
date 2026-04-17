package main

import (
	"fmt"
	"log"
	"practical-assessment/constant"
	"practical-assessment/model"
	"practical-assessment/router"
	"practical-assessment/utils/database"
	"practical-assessment/utils/redis"

	"github.com/sirupsen/logrus"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return
	}

	db := database.GetDB()
	fmt.Println(db)

	err = db.DB.AutoMigrate(&model.Users{})
	if err != nil {
		log.Fatalf(constant.DbMigrationFailed, err)
		return
	}

	err = db.DB.AutoMigrate(&model.Bookings{})
	if err != nil {
		log.Fatalf(constant.DbMigrationFailed, err)
		return
	}

	err = db.DB.AutoMigrate(&model.Slots{})
	if err != nil {
		log.Fatalf(constant.DbMigrationFailed, err)
		return
	}

	//redis connection
	err = redis.InitRedis()
	if err != nil {
		log.Fatalf(constant.RedisInitFailed,err)
		return
	}

	// redisClient, _ := redis.GetRedisClient()
	// fmt.Println(redisClient)

	log.Println(constant.DbMigrationSuccess)

	startRouter()
}

func startRouter() {
	logger := logrus.New()
	router := router.GetRouter()
	logger.Info(fmt.Sprintf(constant.RunningOn, constant.PortDefaultValue))
	router.Run(fmt.Sprintf(":%d", constant.PortDefaultValue))
}
