package model

import (
	"time"
)

type Users struct {
	Id        int64     `gorm:"column:id;primarykey;autoincrement" json:"id"`
	Name      string    `gorm:"column:name;text;notnull" json:"name"`
	Email     string    `gorm:"column:email;text;notnull;unique" json:"email"`
	Password  string    `gorm:"column:password;text;notnull" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;time;null" json:"created_at"`
}

type Orders struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID      int64     `gorm:"column:user_id;not null" json:"user_id" binding:"required,gt=0"`
	User        Users     `gorm:"foreignKey:UserID;references:Id;oncascade:delete"`
	ProductName string    `gorm:"column:product_name;text;not null" json:"product_name" binding:"required"`
	Quantity    int       `gorm:"column:quantity;not null" json:"quantity" binding:"required,gt=0"`
	Price       float64   `gorm:"column:price;decimal;not null" json:"price" binding:"required,gt=0"`
	Status      string    `gorm:"column:status;type:text;default:'CREATED'" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
