package model

import (
	"database/sql"
	"time"
)

type User struct {
	Name      string
	Address   string
	Email     string
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
