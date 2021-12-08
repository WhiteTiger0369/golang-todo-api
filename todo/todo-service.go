package todo

import "ex1/todo-api/common"

type ServiceTodo interface {
	FindAll() ([]TodoDTO, common.DatabaseError)
	FindByID(id uint) (TodoDTO, common.DatabaseError)
	Save(todoDTO TodoDTO) (TodoDTO, common.DatabaseError)
	Delete(id uint)
	FindByUserId(userId uint) ([]TodoDTO, common.DatabaseError)
}
type todoService struct {
	todoRepository RepositoryTodo
}

func ProvideTodoService(t RepositoryTodo) *todoService {
	return &todoService{todoRepository: t}
}

func (t *todoService) FindAll() ([]TodoDTO, common.DatabaseError) {
	res, err := t.todoRepository.FindAll()
	return ToTodoDTOs(res), err
}

func (t *todoService) FindByID(id uint) (TodoDTO, common.DatabaseError) {
	res, err := t.todoRepository.FindByID(id)
	return ToTodoDTO(res), err
}

func (t *todoService) Save(todoDTO TodoDTO) (TodoDTO, common.DatabaseError) {
	res, err := t.todoRepository.Save(ToTodo(todoDTO))
	return ToTodoDTO(res), err

}

func (t *todoService) Delete(id uint) {
	t.todoRepository.Delete(id)
}

func (t *todoService) FindByUserId(userId uint) ([]TodoDTO, common.DatabaseError) {
	res, err := t.todoRepository.FindByUserId(userId)
	return ToTodoDTOs(res), err
}
