package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int64
	UUID      uuid.UUID
	Name      string
	Email     string
	Password  string
	RoleID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
