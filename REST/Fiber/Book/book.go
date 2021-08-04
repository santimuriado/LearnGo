package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/santimuriado/GoProjects/REST/Fiber/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

//curl http://localhost:3000//api/v1/book
func GetBooks(c *fiber.Ctx) {

	db := database.DBConn
	var books []Book

	//Saves every book.
	db.Find(&books)

	//Serializes books into a JSON string and returns them.
	c.JSON(books)
}

/* It's assumed the book exists.
curl http://localhost:3000//api/v1/book/1 or any id. */

func GetBook(c *fiber.Ctx) {

	id := c.Params("id")
	db := database.DBConn
	var book Book

	//Search the book with that id.
	db.Find(&book, id)
	c.JSON(book)
}

//curl -X POST -H "Content-Type: application/json" --data
// "{\"title\": \"Angels and Demons\", \"author\": \"Dan Brown\", \"rating\": 4}"
// http://localhost:3000/api/v1/book
func NewBook(c *fiber.Ctx) {

	db := database.DBConn

	/*
		This is a way to create a book from here which is easier than creating a database
		just for the example.

		var book Book
		book.Title = "Harry Potter Y la Orden del Fenix"
		book.Author = "Jk Rowling"
		book.Rating = 10000
	*/

	//Read the request body and fill the struct.
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	//Creates the book in the db.
	db.Create(&book)

	c.JSON(book)

}

//curl -X DELETE http://localhost:3000/api/v1/book/1
func DeleteBook(c *fiber.Ctx) {

	id := c.Params("id")
	db := database.DBConn

	var book Book

	//Searches the first book with that id.
	db.First(&book, id)

	if book.Title == "" {
		c.Status(500).Send("Can't find a book with that id")
		return
	}

	//Deletes the book.
	db.Delete(&book)
	c.Send("Deleted book succesfully")

}
