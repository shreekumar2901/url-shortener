package dto

type UserRequestDTO struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDTO struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserLoginRequestDTO struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}
