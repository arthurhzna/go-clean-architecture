package security

type JWTService interface {
	Sign(userID int64) (string, error)
	Parse(tokenString string) (*JWTClaims, error)
}
