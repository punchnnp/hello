package handler

import (
	"hello/service"

	"log"

	"github.com/gofiber/fiber/v2"
)

type bookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) bookHandler {
	return bookHandler{bookService: bookService}
}

func (h bookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		log.Println(err)
		return c.SendString("This book ID is not exist")
	}
	return c.JSON(books)
}

func (h bookHandler) GetBookById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	book, err := h.bookService.GetBookById(id)
	if err != nil {
		log.Println(err)
		return c.SendString("This book ID is not exist")
	}
	return c.JSON(book)
}

func (h bookHandler) AddNewBook(c *fiber.Ctx) error {
	book, err := h.bookService.AddNewBook()
	if err != nil {
		log.Println(err)
		return c.SendString("This book ID is not exist")
	}
	return c.JSON(book)
}

func (h bookHandler) UpdateBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	book, err := h.bookService.UpdateBook(id)
	if err != nil {
		log.Println(err)
		return c.SendString("This book ID is not exist")
	}
	return c.JSON(book)
}

func (h bookHandler) DeleteBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	book, err := h.bookService.DeleteBook(int(id))
	if err != nil {
		log.Println(err)
		return c.SendString("This book ID is not exist")
	}
	return c.JSON(book)
}
