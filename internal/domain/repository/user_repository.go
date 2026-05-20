package repository

import (
	"context"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"
)

type UserRepository interface {
	Create(
		ctx context.Context,
		user *entity.User,
	) error

	FindByUUID(
		ctx context.Context,
		uuid string,
	) (*entity.User, error)

	FindByEmail(
		ctx context.Context,
		email string,
	) (*entity.User, error)

	Update(
		ctx context.Context,
		user *entity.User,
	) error

	DeleteByUUID(
		ctx context.Context,
		uuid string,
	) error
}
