package routes

import (
	handlers "ex1/todo-api/handlers/user"
	repositories "ex1/todo-api/repositories/user"
	services "ex1/todo-api/services/user"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {

	readsUserRepository := repositories.NewRepositoryReads(db)
	readsUserService := services.NewServiceReads(readsUserRepository)
	readsUserHandler := handlers.NewHandlerReadsUser(readsUserService)

	groupRoute := route.Group("api/v1")
	groupRoute.GET("/users", readsUserHandler.ReadsUserHandler)

}
