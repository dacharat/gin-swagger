package requests

// CreateUserRequest create user request
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UpdateUserRequest update user request
type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
