package model

import "github.com/pkg/errors"

var (
	ErrorUserNotFound      = errors.New("user not found")
	ErrorUserAlreadyExists = errors.New("user already exists")
)
