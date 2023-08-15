package test

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"  validate:"required,min=2,max=32"`
	Email      string `json:"email" validate:"required,email,min=6,max=32"  gorm:"unique" format:"email" `
	Password   string `json:"password" validate:"required,min=6,max=32"`
}
