package jwt

import (
	"errors"
	"time"

	jwtiface "github.com/arthurhzna/go-clean-architecture/internal/domain/interface/security/jwt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtUtil struct {
	secretKey     string
	issuer        string
	tokenDuration time.Duration
	allowedAlgs   []string
}

func NewJwtUtil(
	secretKey string,
	issuer string,
	tokenDuration time.Duration,
	allowedAlgs []string,
) *JwtUtil {
	return &JwtUtil{
		secretKey:     secretKey,
		issuer:        issuer,
		tokenDuration: tokenDuration,
		allowedAlgs:   allowedAlgs,
	}
}

func (h *JwtUtil) Sign(
	userID int64,
) (string, error) {

	currentTime := time.Now()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwtClaims{
			UserID: userID,
			RegisteredClaims: jwt.RegisteredClaims{
				ID:       uuid.NewString(),
				IssuedAt: jwt.NewNumericDate(currentTime),
				ExpiresAt: jwt.NewNumericDate(
					currentTime.Add(h.tokenDuration),
				),
				Issuer: h.issuer,
			},
		},
	)

	return token.SignedString(
		[]byte(h.secretKey),
	)
}

func (h *JwtUtil) Parse(
	tokenString string,
) (*jwtiface.JWTClaims, error) {

	parser := jwt.NewParser(
		jwt.WithValidMethods(h.allowedAlgs),
		jwt.WithIssuer(h.issuer),
		jwt.WithIssuedAt(),
	)

	token, err := parser.ParseWithClaims(
		tokenString,
		&jwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(h.secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token not valid")
	}

	return &jwtiface.JWTClaims{
		UserID: claims.UserID,
	}, nil
}
