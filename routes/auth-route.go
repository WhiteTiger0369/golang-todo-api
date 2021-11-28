package routes

import (
	"ex1/todo-api/auth"
	"ex1/todo-api/user"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	repo := user.ProvideUserRepository(db)
	service := auth.ProvideAuthService(repo)
	api := auth.ProvideAuthAPI(service)

	groupRoute := route.Group("api/v1")
	groupRoute.POST("/login", api.Login)
}
