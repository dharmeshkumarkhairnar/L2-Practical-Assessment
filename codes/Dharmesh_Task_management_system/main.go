package main

import (
	"log"
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
	}

	db := database.GetDB().DB

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("error in automigrating schema: %s", err.Error())
	}

	if err := db.AutoMigrate(&model.Tasks{}); err != nil {
		log.Fatalf("error in automigrating schema: %s", err.Error())
	}

	log.Print("Database migrated successfully")

	err = redis.InitRedis()
	if err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}

	// redisClient := redis.GetRedisClient()
	// fmt.Println(redisClient)

	startRouter()
}

func startRouter() {
	logger := logrus.New()
	router:=router.GetRouter()
	logger.Info("")
	router.Run(":8080")
}
