package handlers

import (
	"ex1/todo-api/helpers"
	"ex1/todo-api/schemas"
	services "ex1/todo-api/services/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handlerRead struct {
	service services.ServiceRead
}

func NewHandlerReadTodo(service services.ServiceRead) *handlerRead {
	return &handlerRead{service: service}
}

func (h *handlerRead) ReadTodoHandler(ctx *gin.Context) {

	var input schemas.SchemaTodo
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	input.ID = id

	res, err := h.service.ReadTodoService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Todo data is not exist or deleted", err.Code, http.MethodGet, nil)
		return
	default:
		helpers.APIResponse(ctx, "Result Todo data successfully", http.StatusOK, http.MethodGet, res)
	}
}
