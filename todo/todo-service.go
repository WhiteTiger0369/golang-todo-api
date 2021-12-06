package todo

import "ex1/todo-api/common"

type TodoService struct {
	TodoRepository TodoRepository
}

func ProvideTodoService(t TodoRepository) TodoService {
	return TodoService{TodoRepository: t}
}

func (t *TodoService) FindAll() ([]TodoDTO, common.DatabaseError) {
	res, err := t.TodoRepository.FindAll()
	return ToTodoDTOs(res), err
}

func (t *TodoService) FindByID(id uint) (TodoDTO, common.DatabaseError) {
	res, err := t.TodoRepository.FindByID(id)
	return ToTodoDTO(res), err
}

func (t *TodoService) Save(todoDTO TodoDTO) (TodoDTO, common.DatabaseError) {
	res, err := t.TodoRepository.Save(ToTodo(todoDTO))
	return ToTodoDTO(res), err

}

func (t *TodoService) Delete(id uint) {
	t.TodoRepository.Delete(id)
}

func (t *TodoService) FindByUserId(userId uint) ([]TodoDTO, common.DatabaseError) {
	res, err := t.TodoRepository.FindByUserId(userId)
	return ToTodoDTOs(res), err
}
