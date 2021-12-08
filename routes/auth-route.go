package routes

import (
	"ex1/todo-api/api"
	"ex1/todo-api/repositories"
	"ex1/todo-api/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	repo := repositories.ProvideUserRepository(db)
	service := services.ProvideAuthService(repo)
	api := api.ProvideAuthAPI(service)

	groupRoute := route.Group("api/v1")
	groupRoute.POST("/login", api.Login)
}
