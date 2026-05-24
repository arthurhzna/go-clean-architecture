package validation

import (
	"github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	"github.com/arthurhzna/go-clean-architecture/internal/domain/policy"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/field"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/rule"
)

func RegisterUserRules(
	req *request.RegisterUserRequest,
) []core.Rule {

	return []core.Rule{

		rule.RequiredString(
			field.UserNameField,
			req.Name,
		),

		rule.MinLength(
			field.UserNameField,
			req.Name,
			policy.MinNameLength,
		),

		rule.RequiredString(
			field.UserEmailField,
			req.Email,
		),

		rule.Email(
			field.UserEmailField,
			req.Email,
		),

		rule.RequiredString(
			field.UserPasswordField,
			req.Password,
		),

		rule.MinLength(
			field.UserPasswordField,
			req.Password,
			policy.MinPasswordLength,
		),

		rule.RequiredString(
			field.UserConfirmPasswordField,
			req.ConfirmPassword,
		),

		rule.Equal(
			field.UserConfirmPasswordField,
			req.ConfirmPassword,
			req.Password,
		),

		rule.RequiredInt64(
			field.UserRoleField,
			req.RoleID,
		),
	}
}

func LoginUserRules(
	req *request.LoginUserRequest,
) []core.Rule {

	return []core.Rule{

		rule.RequiredString(
			field.UserEmailField,
			req.Email,
		),

		rule.Email(
			field.UserEmailField,
			req.Email,
		),

		rule.RequiredString(
			field.UserPasswordField,
			req.Password,
		),
	}
}
