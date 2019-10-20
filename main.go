package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"todoApp/server/controllers"
	"todoApp/server/resources"
)

var todoController controllers.TodoController
var appRouter resources.AppRouter
var dbContext resources.DbContext
var db *gorm.DB
var router *gin.Engine

func main() {
	db = dbContext.SetContext()
	router = appRouter.SetRouter()

	todoController.SetEndpoints(router, db)

	appRouter.Run()
	dbContext.EndContext()
}
