package application

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"todoApp/server/controllers"
)

var todoController controllers.TodoController

type ControllerManager struct {}
func (c ControllerManager) SetEndpoints(router *gin.Engine, db *gorm.DB) {
	todoController.SetEndpoints(router, db)
}
