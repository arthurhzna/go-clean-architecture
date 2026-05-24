package response

import (
	errordomain "github.com/arthurhzna/go-clean-architecture/internal/domain/error"
)

func MapError(err error) error {
	switch err {
	case errordomain.ErrEmailAlreadyExist:
		return newConflictError(err)

	case errordomain.ErrInvalidCredential:
		return newUnauthorizedError(err)

	case errordomain.ErrUserNotFound,
		errordomain.ErrEmailNotFound,
		errordomain.ErrDeviceNotFound:
		return newNotFoundError(err)
	}

	return err
}
