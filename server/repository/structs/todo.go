package structs

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title string	`gorm:"size:500"`
	Completed bool	`gorm:"default:'false'"`
}
