package requests

import "github.com/go-playground/validator/v10"

type UserRequest struct {
	Email    string `json:"email" binding:"required" gorm:"unique;not null;type:varchar(100)" validate:"email,required" form:"email"`
	Password string `json:"password" binding:"required" gorm:"not null;->;<-" validate:"required,min=5" form:"password"`
}

func (u *UserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
