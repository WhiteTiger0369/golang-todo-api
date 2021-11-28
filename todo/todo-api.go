package todo

import (
	"ex1/todo-api/helpers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type TodoAPI struct {
	todoService TodoService
}

func ProvideTodoAPI(p TodoService) TodoAPI {
	return TodoAPI{todoService: p}
}

func (t *TodoAPI) FindAll(c *gin.Context) {

	res, err := t.todoService.FindAll()

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "Todos data is not exists", err.Code, http.MethodGet, nil)
	default:
		helpers.APIResponse(c, "Results Todos data successfully", http.StatusOK, http.MethodGet, res)
	}
}

func (t *TodoAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := t.todoService.FindByID(uint(id))

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "Todos data is not exists", err.Code, http.MethodGet, nil)
	default:
		helpers.APIResponse(c, "Results Todos data successfully", http.StatusOK, http.MethodGet, res)
	}
}

func (t *TodoAPI) Create(c *gin.Context) {
	var todoDTO TodoDTO
	errIn := c.BindJSON(&todoDTO)
	if errIn != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(errIn)
		return
	}

	res, err := t.todoService.Save(todoDTO)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "Create new todo account failed", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(c, "Create new todo account successfully", http.StatusCreated, http.MethodPost, res)
	}
}

func (t *TodoAPI) Update(c *gin.Context) {
	var todoDTO TodoDTO
	err := c.BindJSON(&todoDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	todo, _ := t.todoService.FindByID(uint(id))
	if todo == (TodoDTO{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	todo.Title = todoDTO.Title
	todo.Content = todoDTO.Content
	t.todoService.Save(todo)

	c.Status(http.StatusOK)
}

func (t *TodoAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	t.todoService.Delete(uint(id))

	c.Status(http.StatusOK)
}
