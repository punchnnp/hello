package main

import (
	"github.com/gofiber/fiber/v2"

	"hello/book"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/api/book", book.AddBook)
	app.Post("/api/book/:id", book.UpdateBook)
	app.Delete("/api/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":3000")
}
