package internal

import "errors"

var (
	ErrUserExists    = errors.New("user exists")
	ErrUserNotExists = errors.New("user does not exist")
)
