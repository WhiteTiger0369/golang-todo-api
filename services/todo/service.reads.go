package services

import (
	"ex1/todo-api/models"
	repositories "ex1/todo-api/repositories/todo"
	"ex1/todo-api/schemas"
)

type ServiceReads interface {
	ReadsTodoService() (*[]models.ModelTodo, schemas.SchemaDatabaseError)
}

type serviceReads struct {
	repository repositories.RepositoryReads
}

func NewServiceReads(repository repositories.RepositoryReads) *serviceReads {
	return &serviceReads{repository: repository}
}

func (s *serviceReads) ReadsTodoService() (*[]models.ModelTodo, schemas.SchemaDatabaseError) {

	res, err := s.repository.ReadsTodoRepository()
	return res, err
}
