package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
