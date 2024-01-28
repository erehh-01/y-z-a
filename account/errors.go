package account

import "errors"

var (
	ErrPermissionDenied = errors.New("0")
	ErrUserNotFound     = errors.New("1")
	ErrAlreadyExists    = errors.New("2")
	ErrUnkown           = errors.New("3")
)
