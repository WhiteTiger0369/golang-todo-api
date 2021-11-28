package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
