package model

import (
	"database/sql"
	"time"
)

type Gorm struct {
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index"`
}
