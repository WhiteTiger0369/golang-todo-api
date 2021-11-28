package handlers

import (
	"ex1/todo-api/helpers"
	services "ex1/todo-api/services/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handlerReads struct {
	service services.ServiceReads
}

func NewHandlerReadsUser(service services.ServiceReads) *handlerReads {
	return &handlerReads{service: service}
}

func (h *handlerReads) ReadsUserHandler(ctx *gin.Context) {

	logrus.Debug("Get all todo")
	res, err := h.service.ReadsUserService()

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Users data is not exists", err.Code, http.MethodGet, nil)
		return
	default:
		helpers.APIResponse(ctx, "Results Users data successfully", http.StatusOK, http.MethodGet, res)
	}
}
