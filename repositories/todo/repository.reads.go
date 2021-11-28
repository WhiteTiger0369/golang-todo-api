package repositories

import (
	"ex1/todo-api/models"
	"ex1/todo-api/schemas"
	"github.com/jinzhu/gorm"
	"net/http"
)

type RepositoryReads interface {
	ReadsTodoRepository() (*[]models.ModelTodo, schemas.SchemaDatabaseError)
}

type repositoryReads struct {
	db *gorm.DB
}

func NewRepositoryReads(db *gorm.DB) *repositoryReads {
	return &repositoryReads{db: db}
}

func (r *repositoryReads) ReadsTodoRepository() (*[]models.ModelTodo, schemas.SchemaDatabaseError) {

	var todos []models.ModelTodo
	db := r.db.Model(&todos)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	resultTodos := db.Debug().Find(&todos)

	if resultTodos.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &todos, <-errorCode
	}

	return &todos, <-errorCode
}
