package book

import (
	"context"
	"database/sql"
	"fmt"
	"hello/database"

	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type Books struct {
	Id          int
	Book_name   string
	Description string
}

var Book []Books

var db, _ = sql.Open("mysql", database.Dns("book"))

func GetBooks(c *fiber.Ctx) error {
	var book []Books
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Print(err.Error())
	} else {
		for rows.Next() {
			var id int
			var name string
			var desc string
			err2 := rows.Scan(&id, &name, &desc)
			if err2 != nil {
				return err2
			} else {
				books := Books{id, name, desc}
				book = append(book, books)
			}
		}
	}
	return c.JSON(book)
}

func GetBook(c *fiber.Ctx) error {
	var book Books
	var id, _ = c.ParamsInt("id")
	err := db.QueryRow("SELECT * FROM books where book_id = ?", id).Scan(&book.Id, &book.Book_name, &book.Description)
	if err != nil {
		return c.SendString("This book ID is not exist")
	}
	return c.JSON(book)
}

func AddBook(c *fiber.Ctx) error {
	var book Books
	query := "INSERT INTO books(book_name, book_desc) VALUES (?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	result, err := db.ExecContext(ctx, query, "John", "Doe")
	if err != nil {
		return c.SendString("Unable to insert new book")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.SendString("Unable to get last insert ID")
	}
	fmt.Printf("ID: %d", id)

	err2 := db.QueryRow("SELECT * FROM books where book_id = ?", id).Scan(&book.Id, &book.Book_name, &book.Description)
	if err2 != nil {
		return c.SendString("This book ID is not exist")
	}

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	var book Books
	var id, _ = c.ParamsInt("id")
	_, err := db.Exec("UPDATE books SET book_name = ?, book_desc = ? where book_id = ?", "name change", "desc change", id)
	if err != nil {
		return err
	}

	err2 := db.QueryRow("SELECT * FROM books where book_id = ?", id).Scan(&book.Id, &book.Book_name, &book.Description)
	if err2 != nil {
		return c.SendString("This book ID is not exist")
	}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	var id, _ = c.ParamsInt("Id")
	_, err := db.Exec("DELETE FROM books where book_id = ?", id)
	if err != nil {
		return c.SendString("This book ID is not exist")
	}
	return c.SendString("This book ID is delete")
}
