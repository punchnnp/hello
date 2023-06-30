package main

import (
	"fmt"
	"time"

	"hello/handler"
	"hello/repository"
	"hello/service"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	app := fiber.New()

	db := initDB()
	bookRepository := repository.NewBookRepositoryDB(db)
	_ = bookRepository
	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	app.Get("/api/book", bookHandler.GetAllBooks)
	app.Get("/api/book/:id", bookHandler.GetBookById)
	app.Post("/api/book", bookHandler.AddNewBook)
	app.Post("/api/book/:id", bookHandler.UpdateBook)
	app.Delete("/api/book/:id", bookHandler.DeleteBook)

	app.Listen(":3000")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDB() *sql.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.hostname"),
		viper.GetInt("db.port"),
		viper.GetString("db.dbname"),
	)

	db, err := sql.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db
}
