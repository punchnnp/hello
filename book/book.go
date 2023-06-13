package book

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Books struct {
	Id          int
	Book_name   string
	Description string
}

var BookJson = `[
	{
		"id": 1,
		"book_name": "this is a book",
		"description": "something that you want to read"},
	{
		"id": 2,
		"book_name": "second book",
		"description": "second book is about the cook"
	}
]`

var Book []Books

var AllBook = json.Unmarshal([]byte(BookJson), &Book)

func GetBooks(c *fiber.Ctx) error {
	return c.JSON(Book)
}

func GetBook(c *fiber.Ctx) error {
	var id, _ = c.ParamsInt("id")
	for i := 0; i < len(Book); i++ {
		if Book[i].Id == id {
			return c.JSON(Book[i])
		}
	}
	return c.SendString("No book found")
}

func AddBook(c *fiber.Ctx) error {
	var tempBook []Books
	var BookJsonInput = `[
		{
			"id": 3,
			"book_name": "this is thrid book",
			"description": "something that you want to eat"
		}
	]`
	err := json.Unmarshal([]byte(BookJsonInput), &tempBook)
	if err != nil {
		fmt.Println(err.Error())
	}
	Book = append(Book, tempBook...)
	return c.JSON(Book)
}

func UpdateBook(c *fiber.Ctx) error {
	var id, _ = c.ParamsInt("id")
	for i := 0; i < len(Book); i++ {
		if Book[i].Id == id {
			Book[i].Book_name = "Change name"
			Book[i].Description = "New description"
			json.Marshal(Book)
			return c.JSON(Book)
		}
	}
	return c.SendString("This book ID is not exist")
}

func DeleteBook(c *fiber.Ctx) error {
	var id, _ = c.ParamsInt("Id")
	for i := 0; i < len(Book); i++ {
		if Book[i].Id == id {
			return c.SendString("Delete book")
		}
	}
	return c.SendString("This book ID is not exist")
}
