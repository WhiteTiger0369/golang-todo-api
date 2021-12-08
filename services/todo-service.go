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
	Update(id uint, todoDTO dtos.TodoDTO) (dtos.TodoDTO, common.DatabaseError)
}

type todoService struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(t repositories.TodoRepository) *todoService {
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

func (t *todoService) Update(id uint, todoDTO dtos.TodoDTO) (dtos.TodoDTO, common.DatabaseError) {
	existsTodo, err := t.todoRepository.FindByID(id)
	if err.Type == "error_01" {
		return todoDTO, err
	}
	existsTodo.Title = todoDTO.Title
	existsTodo.Content = todoDTO.Content
	res, err := t.todoRepository.Save(existsTodo)
	return mappers.ToTodoDTO(res), err
}
