package resources

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"todoApp/server/repository/structs"
)
type Todo = structs.Todo

var db *gorm.DB
var dbErr error

type DbContext struct {}
func (c *DbContext) SetContext() *gorm.DB {
	db, dbErr = gorm.Open("sqlite3", "./db/todoApp.db")
	db.AutoMigrate(&Todo{})
	return db
}

func (c *DbContext) EndContext() {
	dbErr = db.Close()
	if dbErr != nil { fmt.Println(dbErr) }
}