package middleware

import (
	"context"
	"strings"

	errordomain "github.com/arthurhzna/go-clean-architecture/internal/domain/error"
	constants "github.com/arthurhzna/go-clean-architecture/internal/presentation/middleware/constant"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateWithToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(constants.Authorization)
		if token == "" {
			ctx.Error(response.MapError(errordomain.ErrInvalidCredential))
			return
		}

		err := validateBearerToken(ctx, token)
		if err != nil {
			ctx.Error(response.MapError(err))
			return
		}

		err = validateAPIKey(ctx)
		if err != nil {
			ctx.Error(response.MapError(err))
			return
		}

		ctx.Next()
	}
}

func AuthenticateWithoutToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validateAPIKey(ctx)
		if err != nil {
			ctx.Error(response.MapError(err))
			return
		}
		ctx.Next()
	}
}

func validateBearerToken(ctx *gin.Context, token string, jwtSecretKey string) error {
	if !strings.Contains(token, "Bearer") {
		return errordomain.ErrInvalidCredential
	}

	tokenString := extractBearerToken(token)
	if tokenString == "" {
		return errordomain.ErrInvalidCredential
	}

	claims := &dto.Claims{}
	tokenJwt, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errordomain.ErrInvalidCredential
		}

		jwtSecret := []byte(jwtSecretKey)
		return jwtSecret, nil
	})

	if err != nil || !tokenJwt.Valid {
		return errordomain.ErrInvalidCredential
	}

	userLogin := ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), constants.UserLogin, claims.User))
	ctx.Request = userLogin
	ctx.Set(constants.Token, token)
	return nil
}

func extractBearerToken(token string) string {
	arrayToken := strings.Split(token, " ")
	if len(arrayToken) == 2 {
		return arrayToken[1]
	}
	return ""
}

func validateAPIKey(ctx *gin.Context, apiKeyApp string) error {
	apiKeyClient := ctx.GetHeader(constants.XApiKey)

	if apiKeyClient != apiKeyApp {
		return errordomain.ErrInvalidCredential
	}
	return nil
}
