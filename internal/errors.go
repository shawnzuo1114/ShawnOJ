package internal

import "errors"

var (
	ErrUserExists      = errors.New("user exists")
	ErrUserNotExists   = errors.New("user does not exist")
	ErrVerifyExists    = errors.New("verify code exists")
	ErrVerifyCodeWrong = errors.New("verify code was wrong")
)
