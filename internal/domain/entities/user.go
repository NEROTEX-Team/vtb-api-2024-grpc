package model

import (
	"time"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type CreateUser struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type CreateUserWithID struct {
	ID        string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type UserListParams struct {
	Limit  int64
	Offset int64
}

type UserList struct {
	Total int64
	Items []User
}

type UpdateUser struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}
