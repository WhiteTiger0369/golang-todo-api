package todo

func ToTodo(todoDTO TodoDTO) Todo {
	return Todo{Title: todoDTO.Title, Content: todoDTO.Content, UserID: todoDTO.UserID}
}

func ToTodoDTO(todo Todo) TodoDTO {
	return TodoDTO{ID: todo.ID, Title: todo.Title, Content: todo.Content, UserID: todo.UserID, CreatedAt: todo.CreatedAt, UpdatedAt: todo.UpdatedAt}
}

func ToTodoDTOs(products []Todo) []TodoDTO {
	TodoDTOs := make([]TodoDTO, len(products))

	for i, itm := range products {
		TodoDTOs[i] = ToTodoDTO(itm)
	}

	return TodoDTOs
}
