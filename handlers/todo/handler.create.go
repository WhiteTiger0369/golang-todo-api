package handlers

import (
	"ex1/todo-api/helpers"
	"ex1/todo-api/schemas"
	services "ex1/todo-api/services/todo"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handlerCreate struct {
	service services.ServiceCreate
}

func NewHandlerCreateTodo(service services.ServiceCreate) *handlerCreate {
	return &handlerCreate{service: service}
}

func (h *handlerCreate) CreateTodoHandler(ctx *gin.Context) {

	var input schemas.SchemaTodo
	ctx.ShouldBindJSON(&input)

	_, err := h.service.CreateTodoService(&input)
	logrus.Debug(err)

	switch err.Type {
	case "error_02":
		helpers.APIResponse(ctx, "Create new todo account failed", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "Create new todo account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
