package services

import (
	"ex1/todo-api/common"
	"ex1/todo-api/dtos"
	"ex1/todo-api/mappers"
	"ex1/todo-api/repositories"
)

type TodoService interface {
	FindAll() ([]dtos.TodoDTO, common.DatabaseError)
	FindByID(id uint) (dtos.TodoDTO, common.DatabaseError)
	Save(todoDTO dtos.TodoDTO) (dtos.TodoDTO, common.DatabaseError)
	Delete(id uint)
	FindByUserId(userId uint) ([]dtos.TodoDTO, common.DatabaseError)
}
type todoService struct {
	todoRepository repositories.RepositoryTodo
}

func ProvideTodoService(t repositories.RepositoryTodo) *todoService {
	return &todoService{todoRepository: t}
}

func (t *todoService) FindAll() ([]dtos.TodoDTO, common.DatabaseError) {
	res, err := t.todoRepository.FindAll()
	return mappers.ToTodoDTOs(res), err
}

func (t *todoService) FindByID(id uint) (dtos.TodoDTO, common.DatabaseError) {
	res, err := t.todoRepository.FindByID(id)
	return mappers.ToTodoDTO(res), err
}

func (t *todoService) Save(todoDTO dtos.TodoDTO) (dtos.TodoDTO, common.DatabaseError) {
	res, err := t.todoRepository.Save(mappers.ToTodo(todoDTO))
	return mappers.ToTodoDTO(res), err

}

func (t *todoService) Delete(id uint) {
	t.todoRepository.Delete(id)
}

func (t *todoService) FindByUserId(userId uint) ([]dtos.TodoDTO, common.DatabaseError) {
	res, err := t.todoRepository.FindByUserId(userId)
	return mappers.ToTodoDTOs(res), err
}
