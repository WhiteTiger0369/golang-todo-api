package main

import (
	"ex1/todo-api/db"
	"ex1/todo-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	app := SetupRouter()
	logrus.Fatal(app.Run(": 8080"))

}

func SetupRouter() *gin.Engine {
	database := db.DatabaseConnect()
	app := gin.Default()

	routes.InitTodoRoutes(database, app)
	routes.InitUserRoutes(database, app)
	routes.InitAuthRoutes(database, app)

	return app
}
