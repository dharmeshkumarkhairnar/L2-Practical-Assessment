package main

import (
	"fmt"
	"log"
	"practical-assessment/model"
	"practical-assessment/router"
	"practical-assessment/utils/database"
	"practical-assessment/utils/redis"
	// "practical-assessment/utils/redis"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	db := database.GetDB()
	fmt.Println(db)

	if err := db.DB.AutoMigrate(&model.Users{}); err != nil {
		log.Fatal("Error in migrating Users")
	}

	if err := db.DB.AutoMigrate(&model.Expenses{}); err != nil {
		log.Fatal("Error in migrating Expenses")
	}

	log.Print("Database Migration Success")
	err = redis.InitRedis()
	if err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}

	redisClient := redis.GetRedisClient()
	fmt.Println(redisClient)

	startRouter()
}

func startRouter() {
	router := router.GetRouter()
	router.Run(":1010")
}
