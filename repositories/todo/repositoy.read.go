package repositories

import (
	"ex1/todo-api/models"
	"ex1/todo-api/schemas"
	"github.com/jinzhu/gorm"
	"net/http"
)

type RepositoryRead interface {
	ReadTodoRepository(input *schemas.SchemaTodo) (*models.ModelTodo, schemas.SchemaDatabaseError)
}

type repositoryRead struct {
	db *gorm.DB
}

func NewRepositoryRead(db *gorm.DB) *repositoryRead {
	return &repositoryRead{db: db}
}

func (r *repositoryRead) ReadTodoRepository(input *schemas.SchemaTodo) (*models.ModelTodo, schemas.SchemaDatabaseError) {
	var todo models.ModelTodo
	db := r.db.Model(&todo)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	result := db.Debug().First(&todo)

	if result.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &todo, <-errorCode
	}

	return &todo, <-errorCode
}
