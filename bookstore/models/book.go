package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title" validate:"required,title"`
	Author string `json:"author" validate:"required,author"`
}
