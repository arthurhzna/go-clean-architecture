package error

import "errors"

var (
	ErrEmailAlreadyExist = errors.New(
		"email already exists",
	)
)
