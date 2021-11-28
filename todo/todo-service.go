package todo

type TodoService struct {
	TodoRepository TodoRepository
}

func ProvideTodoService(t TodoRepository) TodoService {
	return TodoService{TodoRepository: t}
}

func (t *TodoService) FindAll() []Todo {
	return t.TodoRepository.FindAll()
}

func (t *TodoService) FindByID(id uint) Todo {
	return t.TodoRepository.FindByID(id)
}

func (t *TodoService) Save(todo Todo) Todo {
	t.TodoRepository.Save(todo)

	return todo
}

func (t *TodoService) Delete(Todo Todo) {
	t.TodoRepository.Delete(Todo)
}
