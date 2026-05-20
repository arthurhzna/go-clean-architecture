package request

type RegisterUserRequest struct {
	Name            string
	Password        string
	ConfirmPassword string
	Email           string
	RoleID          int64
}

type LoginUserRequest struct {
	Email    string
	Password string
}
