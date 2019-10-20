package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"todoApp/server/resources/application"
)

var controllerManager application.ControllerManager
var appRouter application.AppRouter
var dbContext application.DbContext
var db *gorm.DB
var router *gin.Engine

func main() {
	db = dbContext.SetContext()
	router = appRouter.SetRouter()

	controllerManager.SetEndpoints(router, db)

	appRouter.Run()
	dbContext.EndContext()
}
