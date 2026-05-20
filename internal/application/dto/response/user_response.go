package response

type UserResponse struct {
	UUID   string
	Name   string
	Email  string
	RoleID int64
}

type RegisterResponse struct {
	User UserResponse
}

type LoginResponse struct {
	User  UserResponse
	Token string
}
