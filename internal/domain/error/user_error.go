package error

import "errors"

var (
	ErrEmailAlreadyExist = errors.New(
		"email already exists",
	)
	ErrUserNotFound = errors.New(
		"user not found",
	)
	ErrEmailNotFound = errors.New(
		"email not found",
	)
)
