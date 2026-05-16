package jwt

type JWTClaims struct {
	UserID int64
}

type IJwtUtil interface {
	Sign(userID int64) (string, error)
	Parse(tokenString string) (*JWTClaims, error)
}
