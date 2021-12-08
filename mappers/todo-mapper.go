package mappers

import (
	"ex1/todo-api/dtos"
	"ex1/todo-api/entities"
)

func ToTodo(todoDTO dtos.TodoDTO) entities.Todo {
	return entities.Todo{Title: todoDTO.Title, Content: todoDTO.Content, UserID: todoDTO.UserID}
}

func ToTodoDTO(todo entities.Todo) dtos.TodoDTO {
	return dtos.TodoDTO{ID: todo.ID, Title: todo.Title, Content: todo.Content, UserID: todo.UserID, CreatedAt: todo.CreatedAt, UpdatedAt: todo.UpdatedAt}
}

func ToTodoDTOs(products []entities.Todo) []dtos.TodoDTO {
	TodoDTOs := make([]dtos.TodoDTO, len(products))

	for i, itm := range products {
		TodoDTOs[i] = ToTodoDTO(itm)
	}

	return TodoDTOs
}
