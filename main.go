package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"todoApp/server/structs"
)

type Todo = structs.Todo
var db *gorm.DB
var err error

func main() {
	// setup db
	db, err = gorm.Open("sqlite3", "./db/todoApp.db")
	handle(err)
	defer db.Close()

	db.AutoMigrate(&Todo{})

	router := gin.Default()

	// serve front end
	router.Use(static.Serve("/", static.LocalFile("./app/build", true)))

	/* --CONTROLLER FUNCTIONS-- */
	// getOne(/hello)
	router.GET(serverPath("/hello"), func (context *gin.Context) {
		context.JSON(200, gin.H {
			"message": "Hello World",
		})
	})

	// post(/todoItem)
	router.POST(serverPath("/todoItem"), func (context *gin.Context) {
		var todo Todo

		handle(context.BindJSON(&todo))

		db.Create(&todo)
		context.JSON(200, todo)
	})

	// getMany(/todos)
	router.GET(serverPath("/todos"), func (context *gin.Context) {
		var todos []Todo
		if err := db.Find(&todos).Error; err != nil {
			context.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			context.JSON(200, todos)
		}
	})

	// getOne(/todoItem/:id)
	router.GET(serverPath("/todoItem/:id"), func (context *gin.Context) {
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
	router.DELETE(serverPath("/todoItem/:id"), func (context *gin.Context) {
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

	router.Run(":8081")
}

func serverPath(endPointPath string) string {
	var serverRoot = "/server"
	return serverRoot + endPointPath
}

func handle(err error) {
	if err != nil { fmt.Println(err) }
}