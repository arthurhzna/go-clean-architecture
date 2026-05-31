package middleware

import (
	enumdomain "github.com/arthurhzna/go-clean-architecture/internal/domain/enum"
	errordomain "github.com/arthurhzna/go-clean-architecture/internal/domain/error"
	"github.com/arthurhzna/go-clean-architecture/internal/domain/security"
	constants "github.com/arthurhzna/go-clean-architecture/internal/presentation/middleware/constant"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/response"
	"github.com/gin-gonic/gin"
)

func CheckRole(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userClaims := ctx.Request.Context().Value(constants.UserLogin)
		if userClaims == nil {
			ctx.Error(response.MapError(errordomain.ErrInvalidCredential))
			ctx.Abort()
			return
		}

		claims, ok := userClaims.(*security.TokenClaims)
		if !ok {
			ctx.Error(response.MapError(errordomain.ErrInvalidCredential))
			ctx.Abort()
			return
		}

		roleName, ok := enumdomain.RoleIDToName[claims.RoleID]
		if !ok {
			ctx.Error(response.MapError(errordomain.ErrInvalidRole))
			ctx.Abort()
			return
		}

		if !contains(roles, roleName) {
			ctx.Error(response.MapError(errordomain.ErrForbidden))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func contains(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}

	return false
}
