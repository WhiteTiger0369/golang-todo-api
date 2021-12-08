package todo

import (
	"ex1/todo-api/common"
	"github.com/jinzhu/gorm"
	"net/http"
)

type RepositoryTodo interface {
	FindAll() ([]Todo, common.DatabaseError)
	FindByID(id uint) (Todo, common.DatabaseError)
	Save(todo Todo) (Todo, common.DatabaseError)
	Delete(id uint)
	FindByUserId(userId uint) ([]Todo, common.DatabaseError)
}

type todoRepository struct {
	DB *gorm.DB
}

func ProvideTodoRepository(DB *gorm.DB) *todoRepository {
	return &todoRepository{DB: DB}
}

func (t *todoRepository) FindAll() ([]Todo, common.DatabaseError) {
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

func (t *todoRepository) FindByID(id uint) (Todo, common.DatabaseError) {
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

func (t *todoRepository) Save(todo Todo) (Todo, common.DatabaseError) {

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

func (t *todoRepository) Delete(id uint) {
	t.DB.Delete(Todo{}, "id = ?", id)
}

func (t *todoRepository) FindByUserId(userId uint) ([]Todo, common.DatabaseError) {
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
