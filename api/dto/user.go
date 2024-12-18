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
