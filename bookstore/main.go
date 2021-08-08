package main

import (
	"github.com/gin-gonic/gin"

	"main.go/controllers"
	"main.go/models"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title" validate:"required,title"`
	Author string `json:"author" validate:"required,author"`
}

func main() {
	r := gin.New()
	// r.GET("/books/", controllers.FindBooks)
	// r.POST("/books/", controllers.CreateBook)
	//r.GET("/books/:id", controllers.WrapHandler(controllers.Handler))

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Run the server
	r.Run()
}
