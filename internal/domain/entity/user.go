package entity

import "github.com/google/uuid"

type User struct {
	ID        string
	UUID      uuid.UUID
	Name      string
	Email     string
	Password  string
	RoleId    int64
	CreatedAt int64
	UpdatedAt int64
}
