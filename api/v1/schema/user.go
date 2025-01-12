package schema

type UserRequest struct {
	FullName string `json:"full_name" validate:"required,min=3" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"password123"`
}

type UserResponse struct {
	ID        uint   `json:"id" example:"1"`
	FullName  string `json:"full_name" example:"John Doe"`
	Email     string `json:"email" example:"john.doe@example.com"`
	CreatedAt string `json:"created_at" example:"2025-01-01"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" validate:"required,min=3" example:"John Updated"`
	Email    string `json:"email" validate:"required,email" example:"updated.email@example.com"`
	Password string `json:"password,omitempty" validate:"omitempty,min=8" example:"newpassword123"`
}

type ErrorMessage struct {
	Error string `json:"error" example:"Invalid input"`
}
