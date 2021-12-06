package todo

import (
	"ex1/todo-api/common"
	"github.com/jinzhu/gorm"
	"net/http"
)

type TodoRepository struct {
	DB *gorm.DB
}

func ProvideTodoRepository(DB *gorm.DB) TodoRepository {
	return TodoRepository{DB: DB}
}

func (t *TodoRepository) FindAll() ([]Todo, common.DatabaseError) {
	var todos []Todo
	errorCode := common.DatabaseError{}
	results := t.DB.Debug().Find(&todos)

	if results.RowsAffected < 1 {
		errorCode = common.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return todos, errorCode
	}

	return todos, errorCode
}

func (t *TodoRepository) FindByID(id uint) (Todo, common.DatabaseError) {
	var todo Todo
	errorCode := common.DatabaseError{}
	res := t.DB.First(&todo, id)

	if res.RowsAffected < 1 {
		errorCode = common.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return todo, errorCode
	}

	return todo, errorCode
}

func (t *TodoRepository) Save(todo Todo) (Todo, common.DatabaseError) {

	errorCode := common.DatabaseError{}
	addUser := t.DB.Debug().Save(&todo)
	if addUser.RowsAffected < 1 {
		errorCode = common.DatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return todo, errorCode
	}

	return todo, errorCode
}

func (t *TodoRepository) Delete(id uint) {
	t.DB.Delete(Todo{}, "id = ?", id)
}

func (t *TodoRepository) FindByUserId(userId uint) ([]Todo, common.DatabaseError) {
	var todos []Todo
	errorCode := common.DatabaseError{}
	results := t.DB.Debug().Find(&todos, "user_id = ?", userId)

	if results.RowsAffected < 1 {
		errorCode = common.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return todos, errorCode
	}

	return todos, errorCode
}
