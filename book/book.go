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
	// var book Books
	db.Query("SELECT * FROM books", &Book)
	// rows, err := db.Query("select * from books")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// } else {
	// 	var result []Books
	// 	for rows.Next() {
	// 		var id int
	// 		var name string
	// 		var desc string
	// 		err2 := rows.Scan(&id, &name, &desc)
	// 		if err2 != nil {
	// 			return err2
	// 		} else {
	// 			allBook := Books{id, name, desc}
	// 			result = append(result, allBook)
	// 		}
	// 	}
	// }
	// fmt.Println(book.Book_name)
	// fmt.Println(book.Description)
	return c.JSON(Book)
}

func GetBook(c *fiber.Ctx) error {
	var book Books
	var id, _ = c.ParamsInt("id")
	err := db.QueryRow("SELECT * FROM tags where id = ?", id).Scan(&book.Id, &book.Book_name, &book.Description)
	if err != nil {
		panic(err.Error())
	}
	return c.JSON(book)
}

func AddBook(c *fiber.Ctx) error {
	query := "INSERT INTO books(book_name, book_desc) VALUES (?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	insertResult, err := db.ExecContext(ctx, query, "John", "Doe")
	if err != nil {
		fmt.Printf("impossible insert: %s", err)
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		fmt.Printf("impossible to retrieve last inserted id: %s", err)
	}
	fmt.Printf("inserted id: %d", id)
	return c.JSON(insertResult)
}

func UpdateBook(c *fiber.Ctx) error {
	var id, _ = c.ParamsInt("id")
	result, err := db.Exec("UPDATE books SET book_name = ?, book_desc = ? where book_id = ?", "name change", "desc change", id)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return c.JSON(result)
}

func DeleteBook(c *fiber.Ctx) error {
	var id, _ = c.ParamsInt("Id")
	_, err := db.Exec("DELETE FROM books where book_id = ?", id)
	if err != nil {
		fmt.Print(err.Error())
	}
	return c.SendString("This book ID is delete")
}
