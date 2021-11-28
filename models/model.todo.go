package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ModelTodo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (model *ModelTodo) BeforeCreate(db *gorm.DB) error {
	model.CreatedAt = time.Now().Local()
	return nil
}

func (model *ModelTodo) BeforeUpdate(db *gorm.DB) error {
	model.UpdatedAt = time.Now().Local()
	return nil
}
