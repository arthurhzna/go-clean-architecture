package jwt

type JWTClaims struct {
	UserID int64
}

type JWTService interface {
	Sign(userID int64) (string, error)
	Parse(tokenString string) (*JWTClaims, error)
}
