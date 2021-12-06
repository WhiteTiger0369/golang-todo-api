package routes

import (
	middleware "ex1/todo-api/middlewares"
	"ex1/todo-api/todo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitTodoRoutes(db *gorm.DB, route *gin.Engine) {

	repo := todo.ProvideTodoRepository(db)
	service := todo.ProvideTodoService(repo)
	api := todo.ProvideTodoAPI(service)

	groupRoute := route.Group("api/v1").Use(middleware.Auth())
	groupRoute.GET("/todos/:id", api.FindByID)
	groupRoute.GET("/todos", api.FindAll)
	groupRoute.POST("/todos", api.Create)
	groupRoute.PUT("/todos/:id", api.Update)
	groupRoute.DELETE("/todos/:id", api.Delete)
	groupRoute.GET("/todos/get-by-current-user", api.FindByUserId)

}
