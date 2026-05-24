package usecase

import (
	"context"

	"github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	"github.com/arthurhzna/go-clean-architecture/internal/application/dto/response"
)

type UserUseCase interface {
	Register(
		ctx context.Context,
		req *request.RegisterUserRequest,
	) (*response.RegisterResponse, error)

	Login(
		ctx context.Context,
		req *request.LoginUserRequest,
	) (*response.LoginResponse, error)
}
