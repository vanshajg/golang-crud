package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vanshajg/golang-crud/book"
	"github.com/vanshajg/golang-crud/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//SetupRoutes creates all routes
func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

// InitDB initialise sqlite db
func InitDB() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database successfully connected")
	database.DBConn.AutoMigrate(&book.Book{})
}

func main() {
	app := fiber.New()
	InitDB()
	SetupRoutes(app)
	app.Listen(":3005")
}
