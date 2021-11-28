package routes

import (
	"ex1/todo-api/user"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {

	repo := user.ProvideUserRepository(db)
	service := user.ProvideUserService(repo)
	api := user.ProvideUserAPI(service)

	groupRoute := route.Group("api/v1")
	groupRoute.GET("/users/:id", api.FindByID)
	groupRoute.GET("/users", api.FindAll)
	groupRoute.POST("/users", api.Create)
	groupRoute.PUT("/users/:id", api.Update)
	groupRoute.DELETE("/users/:id", api.Delete)

}
