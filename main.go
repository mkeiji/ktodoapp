package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID uint `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

func main() {
	//router := gin.Default()
	//
	//// GET(/)
	//router.GET("/", func(context *gin.Context) {
	//	context.JSON(200, gin.H {
	//		"message": "Hello World",
	//	})
	//})
	//
	//router.Run()

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./app/build", true)))
	router.GET("/hello", GetHelloWorld)
	router.Run(":8081")
}

func GetHelloWorld(context *gin.Context) {
	context.JSON(200, gin.H {
		"message": "Hello World",
	})
}