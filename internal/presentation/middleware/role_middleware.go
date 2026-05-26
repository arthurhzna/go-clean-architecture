package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckRole(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims := c.Request.Context().Value(constants.UserLogin)
		if userClaims == nil {
			responseUnauthorized(c, errConstant.ErrUnauthorized.Error())
			return
		}

		user, ok := userClaims.(*dto.UserResponse)
		if !ok {
			responseUnauthorized(c, errConstant.ErrUnauthorized.Error())
			return
		}

		if !contains(roles, user.Role) {
			responseUnauthorized(c, errConstant.ErrUnauthorized.Error())
			return
		}
		c.Next()
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
