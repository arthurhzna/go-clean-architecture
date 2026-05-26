package response

import (
	errordomain "github.com/arthurhzna/go-clean-architecture/internal/domain/error"
)

func MapError(err error) error {
	switch err {
	case errordomain.ErrEmailAlreadyExist:
		return NewConflictError(err)

	case errordomain.ErrInvalidCredential:
		return NewUnauthorizedError(err)

	case errordomain.ErrInvalidRole:
		return NewBadRequestError(err)

	case errordomain.ErrUserNotFound,
		errordomain.ErrEmailNotFound,
		errordomain.ErrDeviceNotFound:
		return NewNotFoundError(err)
	}

	return err
}
