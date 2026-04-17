package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int64     `gorm:"column:uid;primaryKey;autoIncrement" json:"userId"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email;unique" json:"email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
}

type Expenses struct {
	Id          int64     `gorm:"column:eid;primaryKey;autoIncrement"`
	UserId      int64     `gorm:"column:uid" json:"userId"`
	Category    string    `gorm:"column:category" json:"category"`
	Amount      float64   `gorm:"column:amount" json:"amount"`
	Description string    `gorm:"column:description" json:"description"`
	Date        time.Time `gorm:"column:date" json:"date"`
	CreatedAt   time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	Activity    string    `gorm:"column:activity" json:"activity"`

	user Users `gorm:"foreignKey:UserId;references:Id"`
}

type Database struct {
	DB *gorm.DB
}
