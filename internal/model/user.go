package model

import (
	"time"
)

type User struct {
	ID        string
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserInfo struct {
	FirstName string
	LastName  string
	Email     string
}
