package routes

import (
	"ex1/todo-api/todo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitTodoRoutes(db *gorm.DB, route *gin.Engine) {

	repo := todo.ProvideTodoRepository(db)
	service := todo.ProvideTodoService(repo)
	api := todo.ProvideTodoAPI(service)

	groupRoute := route.Group("api/v1")
	groupRoute.GET("/todos/:id", api.FindByID)
	groupRoute.GET("/todos", api.FindAll)
	groupRoute.POST("/todos", api.Create)
	groupRoute.PUT("/todos/:id", api.Update)
	groupRoute.DELETE("/todos/:id", api.Delete)

}