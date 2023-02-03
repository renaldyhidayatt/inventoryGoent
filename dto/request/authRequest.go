package request

type AuthRequest struct {
	Email    string `json:"email" validate:"required,lowercase,email"`
	Password string `json:"password" validate:"required,lowercase"`
}
