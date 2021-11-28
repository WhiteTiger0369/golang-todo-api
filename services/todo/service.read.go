package services

import (
	"ex1/todo-api/models"
	repositories "ex1/todo-api/repositories/todo"
	"ex1/todo-api/schemas"
)

type ServiceRead interface {
	ReadTodoService(input *schemas.SchemaTodo) (*models.ModelTodo, schemas.SchemaDatabaseError)
}

type serviceResult struct {
	repository repositories.RepositoryRead
}

func NewServiceRead(repository repositories.RepositoryRead) *serviceResult {
	return &serviceResult{repository: repository}
}

func (s *serviceResult) ReadTodoService(input *schemas.SchemaTodo) (*models.ModelTodo, schemas.SchemaDatabaseError) {
	var todo schemas.SchemaTodo
	todo.ID = input.ID

	result, err := s.repository.ReadTodoRepository(&todo)

	return result, err
}
