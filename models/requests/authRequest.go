package requests

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"email,required" form:"email"`
	Password string `json:"password" binding:"required" validate:"required,min=5" form:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required" validate:"email,required" form:"email"`
	Password string `json:"password" binding:"required" validate:"required,min=5" form:"password"`
}

func (l *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(l)
}

func (r *RegisterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
