package main

import (
	"fmt"
	"log"
	"practical-assessment/model"
	"practical-assessment/router"
	"practical-assessment/utils/database"
	redisPackage "practical-assessment/utils/redis"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	db := database.GetDB()
	fmt.Println(db)

	err = redisPackage.InitRedis()
	if err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}

	redisClient := redisPackage.GetRedisClient()
	fmt.Println(redisClient)

	err = db.DB.AutoMigrate(
		&model.Users{},
		&model.Orders{},
	)
	if err != nil {
		// fmt.Println("automigration error:", err)
		log.Fatalf("automigration failed: %v", err)
	}
	fmt.Println("")

	startRouter(db.DB, redisClient)
}

func startRouter(db *gorm.DB, redisClient *redis.Client) {
	logger := logrus.New()
	router := router.GetRouter(db, redisClient)
	logger.Info("running server on port: 8080")
	router.Run(fmt.Sprintf(":%d", 8080))
}
