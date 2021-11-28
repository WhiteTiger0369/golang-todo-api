package todo

import (
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
	Todos := t.todoService.FindAll()

	c.JSON(http.StatusOK, gin.H{"Todos": ToTodoDTOs(Todos)})
}

func (t *TodoAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Todo := t.todoService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"Todo": ToTodoDTO(Todo)})
}

func (t *TodoAPI) Create(c *gin.Context) {
	var TodoDTO TodoDTO
	err := c.BindJSON(&TodoDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	createdTodo := t.todoService.Save(ToTodo(TodoDTO))

	c.JSON(http.StatusOK, gin.H{"Todo": ToTodoDTO(createdTodo)})
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
	todo := t.todoService.FindByID(uint(id))
	if todo == (Todo{}) {
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
	todo := t.todoService.FindByID(uint(id))
	if todo == (Todo{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	t.todoService.Delete(todo)

	c.Status(http.StatusOK)
}
