package entities

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}
