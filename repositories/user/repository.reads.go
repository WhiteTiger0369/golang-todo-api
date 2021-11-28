package repositories

import (
	"ex1/todo-api/models"
	"ex1/todo-api/schemas"
	"github.com/jinzhu/gorm"
	"net/http"
)

type RepositoryReads interface {
	ReadsUserRepository() (*[]models.ModelUser, schemas.SchemaDatabaseError)
}

type repositoryReads struct {
	db *gorm.DB
}

func NewRepositoryReads(db *gorm.DB) *repositoryReads {
	return &repositoryReads{db: db}
}

func (r *repositoryReads) ReadsUserRepository() (*[]models.ModelUser, schemas.SchemaDatabaseError) {

	var users []models.ModelUser
	db := r.db.Model(&users)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	resultTodos := db.Debug().Find(&users)

	if resultTodos.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &users, <-errorCode
	}

	return &users, <-errorCode
}
