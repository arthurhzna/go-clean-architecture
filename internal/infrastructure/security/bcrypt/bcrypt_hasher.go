package encryptutils

import (
	"github.com/arthurhzna/go-clean-architecture/internal/domain/security"
	"golang.org/x/crypto/bcrypt"
)

type bcryptEncryptor struct {
	cost int
}

func NewBcryptEncryptor(cost int) security.PasswordHasher {
	return &bcryptEncryptor{
		cost: cost,
	}
}

func (e *bcryptEncryptor) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), e.cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (e *bcryptEncryptor) Check(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
