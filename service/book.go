package service

type BookResponse struct {
	BookID      int    `db:"book_id"`
	Name        string `db:"book_name"`
	Description string `db:"book_desc"`
}

type BookService interface {
	GetAllBooks() ([]BookResponse, error)
	GetBookById(int) (*BookResponse, error)
	AddNewBook() (*BookResponse, error)
	UpdateBook(int) (*BookResponse, error)
	DeleteBook(int) (string, error)
}
