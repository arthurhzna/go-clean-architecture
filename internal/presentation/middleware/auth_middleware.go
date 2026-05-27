// internal/presentation/middleware/auth_middleware.go

package middleware

import (
	"context"
	"strings"

	errordomain "github.com/arthurhzna/go-clean-architecture/internal/domain/error"
	"github.com/arthurhzna/go-clean-architecture/internal/domain/security"
	constants "github.com/arthurhzna/go-clean-architecture/internal/presentation/middleware/constant"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/response"

	"github.com/gin-gonic/gin"
)

func AuthenticateWithToken(
	jwtUtils security.TokenService,
) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.GetHeader(
			constants.Authorization,
		)

		token, err := extractAndValidateBearerToken(
			token,
		)

		if err != nil {
			ctx.Error(response.MapError(err))
			return
		}

		claims, err := jwtUtils.Parse(token)
		if err != nil {
			ctx.Error(response.MapError(err))
			return
		}

		requestContext := context.WithValue(
			ctx.Request.Context(),
			constants.UserLogin,
			claims,
		)

		ctx.Request = ctx.Request.WithContext(
			requestContext,
		)

		ctx.Next()
	}
}

func AuthenticateWithApiKey(
	apiKeyApp string,
) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		apiKeyClient := ctx.GetHeader(
			constants.XApiKey,
		)

		if apiKeyClient != apiKeyApp {
			ctx.Error(response.MapError(
				errordomain.ErrInvalidCredential,
			))
			return
		}

		ctx.Next()
	}
}

func extractAndValidateBearerToken(
	token string,
) (string, error) {

	if token == "" {
		return "", errordomain.ErrInvalidCredential
	}

	if !strings.HasPrefix(
		token,
		constants.BearerSchema,
	) {
		return "", errordomain.ErrInvalidCredential
	}

	rawToken := strings.TrimPrefix(
		token,
		constants.BearerSchema,
	)

	if rawToken == "" {
		return "", errordomain.ErrInvalidCredential
	}

	return rawToken, nil
}
