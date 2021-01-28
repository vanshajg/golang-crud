package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanshajg/golang-crud/database"
	"gorm.io/gorm"
)

//Book type
type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}

//GetBooks get all books
func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

//GetBook get single book based on id
func GetBook(c *fiber.Ctx) error {
	db := database.DBConn
	var book Book
	id := c.Params("id")
	db.First(&book, id)
	return c.JSON(book)
}

//NewBook creates new book
func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	var err error
	err = c.BodyParser(book)
	if err != nil {
		c.Status(503).SendString(err.Error())
		return nil
	}
	db.Create(&book)
	return c.JSON(book)
}

//DeleteBook deletes book for id
func DeleteBook(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).SendString("No book with this id")
		return nil
	}
	db.Delete(&book)
	return c.SendString("book deleted successfully")
}
