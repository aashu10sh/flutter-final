package auth

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gte=10"`
}
