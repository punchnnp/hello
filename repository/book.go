package repository

type Book struct {
	BookID      int    `json:"book_id"`
	Name        string `json:"book_name"`
	Description string `json:"book_desc"`
}

//go:generate mockgen -package=repository -destination book_gomock.go . BookRepository

type BookRepository interface {
	GetAll() ([]Book, error)
	GetById(int) (*Book, error)
	AddBook() (*Book, error)
	UpdateBook(int) (*Book, error)
	DeleteBook(int) (string, error)
}
