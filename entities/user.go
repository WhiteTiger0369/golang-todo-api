package entities

import (
	"ex1/todo-api/pkg"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (model *User) BeforeCreate(db *gorm.DB) error {
	model.Password = pkg.HashPassword(model.Password)
	return nil
}

func (model *User) BeforeUpdate(db *gorm.DB) error {
	model.Password = pkg.HashPassword(model.Password)
	return nil
}
