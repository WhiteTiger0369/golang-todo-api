package db

import (
	"ex1/todo-api/todo"
	"ex1/todo-api/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func DatabaseConnect() *gorm.DB {
	dbDriver := "mysql"
	dbUser := "tiger"
	dbPass := "Tiger123!@#"
	dbName := "to-do-app"
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")

	if err != nil {
		defer logrus.Info("Connect into Database Failed")
		logrus.Fatal(err.Error())
	}

	db.AutoMigrate(
		&todo.Todo{},
		&user.User{},
	)

	return db
}
