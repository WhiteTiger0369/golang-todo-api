package repositories

import (
	"ex1/todo-api/models"
	"ex1/todo-api/schemas"
	"github.com/jinzhu/gorm"
	"net/http"
)

type RepositoryCreate interface {
	CreateTodoRepository(input *schemas.SchemaTodo) (*models.ModelTodo, schemas.SchemaDatabaseError)
}

type repositoryCreate struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repositoryCreate {
	return &repositoryCreate{db: db}
}

func (r *repositoryCreate) CreateTodoRepository(input *schemas.SchemaTodo) (*models.ModelTodo, schemas.SchemaDatabaseError) {
	var todo models.ModelTodo
	db := r.db.Model(&todo)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	todo.Title = input.Title
	todo.Content = input.Content

	addNewTodo := db.Debug().Create(&todo).Commit()

	if addNewTodo.Error != nil {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &todo, <-errorCode
	}

	return &todo, <-errorCode

}
