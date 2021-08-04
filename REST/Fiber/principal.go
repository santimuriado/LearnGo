package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	book "github.com/santimuriado/GoProjects/REST/Fiber/Book"
	"github.com/santimuriado/GoProjects/REST/Fiber/database"
)

func setupRoutes(app *fiber.App) {

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book/", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)

}

func initDatabase() {

	var err error

	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Can't connect to database")
	}
	fmt.Println("Opened connection to database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Migrated database succesfully")
}

func main() {

	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
