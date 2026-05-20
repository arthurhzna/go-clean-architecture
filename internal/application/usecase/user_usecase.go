package usecase

import (
	"context"

	"github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	"github.com/arthurhzna/go-clean-architecture/internal/application/dto/response"

	"github.com/arthurhzna/go-clean-architecture/internal/application/validation/builder"
	"github.com/arthurhzna/go-clean-architecture/internal/application/validation/rule"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"

	repositoryiface "github.com/arthurhzna/go-clean-architecture/internal/domain/repository"

	securitydomain "github.com/arthurhzna/go-clean-architecture/internal/domain/security"

	servicedomain "github.com/arthurhzna/go-clean-architecture/internal/domain/service"

	errordomain "github.com/arthurhzna/go-clean-architecture/internal/domain/error"
)

type UserUseCase struct {
	uow repositoryiface.UnitOfWork

	passwordHasher securitydomain.PasswordHasher
	tokenService   securitydomain.TokenService
	uuidGenerator  servicedomain.UUIDGenerator
}

func NewUserUseCase(
	uow repositoryiface.UnitOfWork,
	passwordHasher securitydomain.PasswordHasher,
	tokenService securitydomain.TokenService,
	uuidGenerator servicedomain.UUIDGenerator,
) *UserUseCase {
	return &UserUseCase{
		uow:            uow,
		passwordHasher: passwordHasher,
		tokenService:   tokenService,
		uuidGenerator:  uuidGenerator,
	}
}

func (u *UserUseCase) Register(
	ctx context.Context,
	req request.RegisterUserRequest,
) (*response.RegisterResponse, error) {

	err := rule.Execute(
		builder.RegisterUserRules(req),
	)

	if err != nil {
		return nil, err
	}

	var registeredUser *entity.User

	err = u.uow.WithTransaction(
		ctx,
		func(txUow repositoryiface.UnitOfWork) error {

			existingUser, err := txUow.
				UserRepository().
				FindByEmail(
					ctx,
					req.Email,
				)

			if err != nil {
				return err
			}

			if existingUser != nil {
				return errordomain.ErrEmailAlreadyExist
			}

			hashedPassword, err := u.passwordHasher.Hash(
				req.Password,
			)

			if err != nil {
				return err
			}

			user := &entity.User{
				UUID:     u.uuidGenerator.New(),
				Name:     req.Name,
				Email:    req.Email,
				Password: hashedPassword,
				RoleID:   req.RoleID,
			}

			err = txUow.
				UserRepository().
				Create(
					ctx,
					user,
				)

			if err != nil {
				return err
			}

			registeredUser = user

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return &response.RegisterResponse{
		User: response.UserResponse{
			UUID:   registeredUser.UUID,
			Name:   registeredUser.Name,
			Email:  registeredUser.Email,
			RoleID: registeredUser.RoleID,
		},
	}, nil
}

func (u *UserUseCase) Login(
	ctx context.Context,
	req request.LoginUserRequest,
) (*response.LoginResponse, error) {

	errs := rule.Execute(
		builder.LoginUserRules(req),
	)

	if len(errs) > 0 {
		return nil, errs[0]
	}

	user, err := u.uow.
		UserRepository().
		FindByEmail(
			ctx,
			req.Email,
		)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errordomain.ErrInvalidCredential
	}

	isValid := u.passwordHasher.Check(
		req.Password,
		user.Password,
	)

	if !isValid {
		return nil, errordomain.ErrInvalidCredential
	}

	token, err := u.tokenService.Generate(
		user.ID,
		user.RoleID,
	)

	if err != nil {
		return nil, err
	}

	return &response.LoginResponse{
		User: response.UserResponse{
			UUID:   user.UUID,
			Name:   user.Name,
			Email:  user.Email,
			RoleID: user.RoleID,
		},
		Token: token,
	}, nil
}
