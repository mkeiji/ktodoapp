package structs

type Todo struct {
	Id uint `gorm:"primary_key"`
	Title string	`gorm:"size:500"`
	Completed bool	`gorm:"default:'false'"`
}
