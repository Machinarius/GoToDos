package database

import (
	"database/sql"

	"gorm.io/gorm"
)

type TodoRecord struct {
	gorm.Model
	CompletedAt sql.NullTime
	Text        string
}
