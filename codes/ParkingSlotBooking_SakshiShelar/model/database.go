package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Database struct {
	DB    *gorm.DB
	SqlDB *sql.DB
}

type Users struct {
	Id        int       `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email;uniqueIndex" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type Bookings struct {
	Id            int       `gorm:"column:id;primaryKey" json:"id"`
	UserId        int       `gorm:"column:user_id" json:"user_id"`
	SlotNumber    int       `gorm:"column:slot_number" json:"slot_number"`
	VehicleNumber string    `gorm:"column:vehicle_number" json:"vehicle_number"`
	Status        string    `gorm:"column:status" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`

	Users Users `gorm:"foreignKey:UserId;references:Id"`
}

type Slots struct {
	Id         int    `gorm:"column:id;primaryKey" json:"id"`
	SlotNumber int    `gorm:"column:slot_number;uniqueIndex" json:"slot_number"`
	Status     string `gorm:"column:status" json:"status"`
}
