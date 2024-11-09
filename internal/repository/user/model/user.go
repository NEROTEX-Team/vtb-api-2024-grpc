package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	FirstName string
	LastName  string
	Email     string
}
