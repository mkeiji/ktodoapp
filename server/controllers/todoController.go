package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"todoApp/server/repository/structs"
)
type Todo = structs.Todo

type TodoController struct {}

func (c *TodoController) SetEndpoints(router *gin.Engine, db *gorm.DB) {
	// getOne(/hello)
	router.GET(serverRoot("/hello"), func (context *gin.Context) {
		context.JSON(200, gin.H {
			"message": "Hello World",
		})
	})

	// post(/todoItem)
	router.POST(serverRoot("/todoItem"), func (context *gin.Context) {
		var todo Todo

		handle(context.BindJSON(&todo))

		db.Create(&todo)
		context.JSON(200, todo)
	})

	// getMany(/todos)
	router.GET(serverRoot("/todos"), func (context *gin.Context) {
		var todos []Todo
		if err := db.Find(&todos).Error; err != nil {
			context.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			context.JSON(200, todos)
		}
	})

	// getOne(/todoItem/:id)
	router.GET(serverRoot("/todoItem/:id"), func (context *gin.Context) {
		id := context.Params.ByName("id")
		var todo Todo
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			context.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			context.JSON(200, todo)
		}
	})

	// deleteOne(/todoItem/:id)
	router.DELETE(serverRoot("/todoItem/:id"), func (context *gin.Context) {
		id := context.Params.ByName("id")
		var todo Todo
		if err := db.Where("id = ?", id).Delete(&todo).Error; err != nil {
			context.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			context.JSON(200, todo)
		}
	})
	/* --END-- */
}

func handle(err error) {
	if err != nil { fmt.Println(err) }
}

func serverRoot(endPoint string) string {
	var serverRoot = "/server"
	return serverRoot + endPoint
}