package error

import "errors"

var (
	ErrInvalidCredential = errors.New(
		"invalid credential",
	)

	ErrEmailAlreadyExist = errors.New(
		"email already exists",
	)
)
