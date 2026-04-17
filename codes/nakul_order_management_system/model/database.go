package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Database struct {
	DB    *gorm.DB
	SqlDB *sql.DB
}
