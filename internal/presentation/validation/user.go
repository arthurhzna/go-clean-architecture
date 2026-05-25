package validation

import (
	"github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	"github.com/arthurhzna/go-clean-architecture/internal/domain/policy"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/constant"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/rule"
)

func RegisterUserRules(
	req *request.RegisterUserRequest,
) []core.Rule {

	return []core.Rule{

		rule.RequiredString(
			constant.UserNameField,
			req.Name,
		),

		rule.MinLength(
			constant.UserNameField,
			req.Name,
			policy.MinNameLength,
		),

		rule.RequiredString(
			constant.UserEmailField,
			req.Email,
		),

		rule.Email(
			constant.UserEmailField,
			req.Email,
		),

		rule.RequiredString(
			constant.UserPasswordField,
			req.Password,
		),

		rule.MinLength(
			constant.UserPasswordField,
			req.Password,
			policy.MinPasswordLength,
		),

		rule.RequiredString(
			constant.UserConfirmPasswordField,
			req.ConfirmPassword,
		),

		rule.Equal(
			constant.UserConfirmPasswordField,
			req.ConfirmPassword,
			req.Password,
		),

		rule.RequiredInt64(
			constant.UserRoleField,
			req.RoleID,
		),
	}
}

func LoginUserRules(
	req *request.LoginUserRequest,
) []core.Rule {

	return []core.Rule{

		rule.RequiredString(
			constant.UserEmailField,
			req.Email,
		),

		rule.Email(
			constant.UserEmailField,
			req.Email,
		),

		rule.RequiredString(
			constant.UserPasswordField,
			req.Password,
		),
	}
}
