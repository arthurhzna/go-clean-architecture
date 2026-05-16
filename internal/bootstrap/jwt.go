package bootstrap

import (
	"time"

	"github.com/arthurhzna/go-clean-architecture/internal/config"
	jwtinfra "github.com/arthurhzna/go-clean-architecture/internal/infrastructure/security/jwt"
)

func NewJwtUtil(
	cfg *config.Config,
) *jwtinfra.JwtUtil {

	return jwtinfra.NewJwtUtil(
		cfg.Jwt.SecretKey,
		cfg.Jwt.Issuer,
		time.Duration(cfg.Jwt.TokenDuration)*time.Minute,
		cfg.Jwt.AllowedAlgs,
	)
}
