package services

import (
	"ex1/todo-api/models"
	repositories "ex1/todo-api/repositories/todo"
	"ex1/todo-api/schemas"
)

type ServiceCreate interface {
	CreateTodoService(input *schemas.SchemaTodo) (*models.ModelTodo, schemas.SchemaDatabaseError)
}

type serviceCreate struct {
	repository repositories.RepositoryCreate
}

func NewServiceCreate(repository repositories.RepositoryCreate) *serviceCreate {
	return &serviceCreate{repository: repository}
}

func (s *serviceCreate) CreateTodoService(input *schemas.SchemaTodo) (*models.ModelTodo, schemas.SchemaDatabaseError) {

	var todo schemas.SchemaTodo
	todo.Title = input.Title
	todo.Content = input.Content

	res, err := s.repository.CreateTodoRepository(&todo)
	return res, err
}
