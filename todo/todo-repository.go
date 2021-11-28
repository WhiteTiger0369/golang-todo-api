package todo

import "github.com/jinzhu/gorm"

type TodoRepository struct {
	DB *gorm.DB
}

func ProvideTodoRepository(DB *gorm.DB) TodoRepository {
	return TodoRepository{DB: DB}
}

func (t *TodoRepository) FindAll() []Todo {
	var todo []Todo
	t.DB.Find(&todo)

	return todo
}

func (t *TodoRepository) FindByID(id uint) Todo {
	var todo Todo
	t.DB.First(&todo, id)

	return todo
}

func (t *TodoRepository) Save(todo Todo) Todo {
	t.DB.Save(&todo)

	return todo
}

func (t *TodoRepository) Delete(todo Todo) {
	t.DB.Delete(&todo)
}
