package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint64    `gorm:"column:id;primarykey" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Email      string    `gorm:"column:email;uniqueIndex" json:"email"`
	Password   string    `gorm:"column:password" json:"password"`
	Created_at time.Time `gorm:"column:created_at" json:"created_at"`
}

type Tasks struct {
	ID          uint64    `gorm:"column:id;primarykey" json:"id"`
	UserID      uint64    `gorm:"column:user_id" json:"user_id"`
	Title       string    `gorm:"column:title;not null" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Status      string    `gorm:"column:status;default:pending" json:"status"`
	Priority    string    `gorm:"column:priority" json:"priority"`
	Created_at  time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at  time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted   bool      `gorm:"column:is_deleted;default:false" json:"is_deleted"`

	UserIDKey User `gorm:"foreignkey:UserID;references:ID"`
}

type Database struct {
	DB    *gorm.DB
	SqlDB *sql.DB
}
