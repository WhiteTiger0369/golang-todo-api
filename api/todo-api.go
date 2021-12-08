package api

import (
	"ex1/todo-api/dtos"
	"ex1/todo-api/helpers"
	"ex1/todo-api/pkg"
	"ex1/todo-api/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type todoAPI struct {
	todoService services.TodoService
}

func NewTodoAPI(p services.TodoService) *todoAPI {
	return &todoAPI{todoService: p}
}

func (t *todoAPI) FindAll(c *gin.Context) {

	res, err := t.todoService.FindAll()

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "Todos data is not exists", err.Code, http.MethodGet, nil)
	default:
		helpers.APIResponse(c, "Results Todos data successfully", http.StatusOK, http.MethodGet, res)
	}
}

func (t *todoAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := t.todoService.FindByID(uint(id))

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "Todos data is not exists", err.Code, http.MethodGet, nil)
	default:
		helpers.APIResponse(c, "Results Todos data successfully", http.StatusOK, http.MethodGet, res)
	}
}

func (t *todoAPI) Create(c *gin.Context) {
	var todoDTO dtos.TodoDTO
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

func (t *todoAPI) Update(c *gin.Context) {
	var todoDTO dtos.TodoDTO
	err := c.BindJSON(&todoDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	t.todoService.Update(uint(id), todoDTO)

	c.Status(http.StatusOK)
}

func (t *todoAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	t.todoService.Delete(uint(id))

	c.Status(http.StatusOK)
}

func (t *todoAPI) FindByUserId(c *gin.Context) {
	token, _ := pkg.VerifyTokenHeader(c, "JWT_SECRET")
	accessToken := pkg.DecodeToken(token)
	userId := accessToken.Claims.ID

	res, err := t.todoService.FindByUserId(userId)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "Todos data is not exists", err.Code, http.MethodGet, nil)
	default:
		helpers.APIResponse(c, "Results Todos data successfully", http.StatusOK, http.MethodGet, res)
	}
}
