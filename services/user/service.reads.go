package services

import (
	"ex1/todo-api/models"
	repositories "ex1/todo-api/repositories/user"
	"ex1/todo-api/schemas"
)

type ServiceReads interface {
	ReadsUserService() (*[]models.ModelUser, schemas.SchemaDatabaseError)
}

type serviceReads struct {
	repository repositories.RepositoryReads
}

func NewServiceReads(repository repositories.RepositoryReads) *serviceReads {
	return &serviceReads{repository: repository}
}

func (s *serviceReads) ReadsUserService() (*[]models.ModelUser, schemas.SchemaDatabaseError) {

	res, err := s.repository.ReadsUserRepository()
	return res, err
}
