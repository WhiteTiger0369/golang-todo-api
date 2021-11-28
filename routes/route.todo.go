package routes

import (
	handlers "ex1/todo-api/handlers/todo"
	repositories "ex1/todo-api/repositories/todo"
	services "ex1/todo-api/services/todo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitTodoRoutes(db *gorm.DB, route *gin.Engine) {

	readTodoRepository := repositories.NewRepositoryRead(db)
	readTodoService := services.NewServiceRead(readTodoRepository)
	readTodoHandler := handlers.NewHandlerReadTodo(readTodoService)

	readsTodoRepository := repositories.NewRepositoryReads(db)
	readsTodoService := services.NewServiceReads(readsTodoRepository)
	readsTodoHandler := handlers.NewHandlerReadsTodo(readsTodoService)

	createTodoRepository := repositories.NewRepositoryCreate(db)
	createTodoService := services.NewServiceCreate(createTodoRepository)
	createTodoHandler := handlers.NewHandlerCreateTodo(createTodoService)

	groupRoute := route.Group("api/v1")
	groupRoute.GET("/todos/:id", readTodoHandler.ReadTodoHandler)
	groupRoute.GET("/todos", readsTodoHandler.ReadsTodoHandler)
	groupRoute.POST("/todos", createTodoHandler.CreateTodoHandler)

}
