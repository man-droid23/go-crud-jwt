package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required" gorm:"unique;not null;type:varchar(100)" validate:"email,required" form:"email"`
	Password string `json:"-" binding:"required" gorm:"not null;column:password;->'<-'" validate:"required,min=5" form:"password"`
}
