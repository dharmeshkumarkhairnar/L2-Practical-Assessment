package database

import (
	"errors"
	"practical-assessment/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *model.Database

// InitDB opens a PostgreSQL connection via GORM using DATABASE_URL
// InitDB initializes the database only once
func InitDB() error {

	dsn := "host=localhost user=postgres password=sanjupost dbname=postgres sslmode=disable TimeZone=Asia/Kolkata"

	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("DB Connection Failed")
	}

	setDB(gdb)

	return nil
}

func setDB(data *gorm.DB) {
	db = &model.Database{DB: data}
}

// GetDB returns the shared model.Database. Call InitDB first.
func GetDB() *model.Database {
	return db
}
