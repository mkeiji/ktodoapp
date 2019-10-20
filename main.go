package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"todoApp/server/controllers"
	"todoApp/server/structs"
)
type Todo = structs.Todo

var todoController controllers.TodoController
var db *gorm.DB
var router *gin.Engine
var err error

func main() {
	initializeDb()
	initializeRouter()

	todoController.SetEndpoints(router, db)

	handle(router.Run(":8081"))
	handle(db.Close())
}

func initializeDb() {
	db, err = gorm.Open("sqlite3", "./db/todoApp.db")
	handle(err)
	db.AutoMigrate(&Todo{})
}

func initializeRouter() {
	router = gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./app/build", true)))
}

func handle(err error) {
	if err != nil { fmt.Println(err) }
}