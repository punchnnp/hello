package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type bookRepositoryDB struct {
	db *sql.DB
}

func NewBookRepositoryDB(db *sql.DB) bookRepositoryDB {
	return bookRepositoryDB{db: db}
}

func (r bookRepositoryDB) GetAll() ([]Book, error) {
	books := []Book{}
	query := "SELECT * FROM books"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var id int
			var name string
			var desc string
			err := rows.Scan(&id, &name, &desc)
			if err != nil {
				return nil, err
			} else {
				book := Book{id, name, desc}
				books = append(books, book)
			}
		}
	}

	return books, nil
}

func (r bookRepositoryDB) GetById(id int) (*Book, error) {
	book := Book{}
	query := "SELECT * FROM books WHERE book_id = ?"
	err := r.db.QueryRow(query, id).Scan(&book.BookID, &book.Name, &book.Description)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r bookRepositoryDB) AddBook() (*Book, error) {
	query := "INSERT INTO books (book_name, book_desc) VALUES (?, ?)"
	result, err := r.db.Exec(query, "What makes you", "To know more about yourself")
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.GetById(int(id))
}

func (r bookRepositoryDB) UpdateBook(id int) (*Book, error) {
	query := "UPDATE books SET book_name = ?, book_desc = ? where book_id = ?"
	_, err := r.db.Exec(query, "Name updated", "Description updated", id)
	if err != nil {
		return nil, err
	}
	return r.GetById(id)
}

func (r bookRepositoryDB) DeleteBook(id int) (string, error) {
	query := "DELETE FROM books WHERE book_id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return "", err
	}
	return "This book ID is deleted", nil
}
