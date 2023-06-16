package repository

import "errors"

type bookRepositoryMock struct {
	books []Book
}

func NewBookRepositoryMock() bookRepositoryMock {
	books := []Book{
		{BookID: 1, Name: "First Book", Description: "Tell something about this book"},
		{BookID: 2, Name: "Second Book", Description: "This book is about how to cook"},
	}
	return bookRepositoryMock{books: books}
}

func (m bookRepositoryMock) GetAll() ([]Book, error) {
	return m.books, nil
}

func (m bookRepositoryMock) GetById(id int) (*Book, error) {
	for _, book := range m.books {
		if book.BookID == id {
			return &book, nil
		}
	}
	return nil, errors.New("this book is not exist")
}

func (m bookRepositoryMock) AddBook() (*Book, error) {
	newBook := Book{
		BookID:      3,
		Name:        "Add book",
		Description: "this is new book",
	}
	m.books = append(m.books, newBook)
	return &newBook, nil
}

func (m bookRepositoryMock) UpdateBook(id int) (*Book, error) {
	for i := 0; i < len(m.books); i++ {
		if m.books[i].BookID == id {
			m.books[i].Name = "Change name"
			m.books[i].Description = "Change description"
			return &m.books[i], nil
		}
	}
	return nil, errors.New("this book is not exist")
}

func (m bookRepositoryMock) DeleteBook(id int) (string, error) {
	for i := 0; i < len(m.books); i++ {
		if m.books[i].BookID == id {
			m.books = append(m.books[:i], m.books[i+1:]...)
			return "this book ID is deleted", nil
		}
	}
	return "", errors.New("this book is not exist")
}
