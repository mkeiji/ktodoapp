package application

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"os"
)

var r *gin.Engine
var routerErr error

type AppRouter struct {}
func (c AppRouter) SetRouter() *gin.Engine {
	r = gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./app/build", true)))
	return r
}

func (c AppRouter) Run() {
	routerErr = r.Run(":" + os.Getenv("PORT"))
	if routerErr != nil { fmt.Println(routerErr) }
}