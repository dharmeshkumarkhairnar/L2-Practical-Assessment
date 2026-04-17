package database

import (
	"context"
	"fmt"
	"practical-assessment/model"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *model.Database
var once sync.Once

// InitDB opens a PostgreSQL connection via GORM using DATABASE_URL
// InitDB initializes the database only once
func InitDB() error {
	var initErr error

	once.Do(func() {
		// dsn := os.Getenv("DATABASE_URL")
		// if dsn == "" {
		// 	initErr = errors.New("DATABASE_URL is not set")
		// 	return
		// }

		dsn := "host=localhost port=5432 user=postgres password=6414 dbname=practical_assessment sslmode=disable TimeZone=Asia/Kolkata"
		gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			initErr = fmt.Errorf("open database: %w", err)
			return
		}

		sqlDB, err := gdb.DB()
		if err != nil {
			initErr = fmt.Errorf("get underlying sql db: %w", err)
			return
		}

		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := sqlDB.PingContext(ctx); err != nil {
			initErr = fmt.Errorf("ping database: %w", err)
			return
		}

		db = &model.Database{
			DB:    gdb,
			SqlDB: sqlDB,
		}
	})

	return initErr
}

// GetDB returns the shared model.Database. Call InitDB first.
func GetDB() *model.Database {
	return db
}
