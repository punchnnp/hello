package main

import (
	"fmt"
	// "hello/database"

	"hello/handler"
	"hello/repository"
	"hello/service"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	// "github.com/spf13/viper"
)

// func setupRoutes(app *fiber.App, h bookHandler) {
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, World!")
// 	})

// 	app.Get("/api/book", h.GetAll)
// 	app.Get("/api/book/:id", book.GetBook)
// 	app.Post("/api/book", book.AddBook)
// 	app.Post("/api/book/:id", book.UpdateBook)
// 	app.Delete("/api/book/:id", book.DeleteBook)
// }

// func initConfig() {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath(".")
// 	fmt.Println(viper.GetString("db.driver"))
// }

func main() {
	// initConfig()
	// dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
	// 	viper.GetString("db.username"),
	// 	viper.GetString("db.password"),
	// 	viper.GetString("db.hostname"),
	// 	viper.GetInt("db.port"),
	// 	viper.GetString("db.driver"),
	// )

	app := fiber.New()

	// fmt.Println(dsn)
	// database.ConnectDB()

	db, err := sql.Open("mysql", "root:1991932@tcp(127.0.0.1:3306)/book")
	if err != nil {
		panic(err)
	}

	bookRepository := repository.NewBookRepositoryDB(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	app.Get("/api/book", bookHandler.GetAllBooks)
	app.Get("/api/book/:id", bookHandler.GetBookById)
	app.Post("/api/book", bookHandler.AddNewBook)
	app.Post("/api/book/:id", bookHandler.UpdateBook)
	app.Delete("/api/book/:id", bookHandler.DeleteBook)

	// setupRoutes(app, bookHandler)

	books, err := bookService.GetBookById(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(books)

	app.Listen(":3000")
}
