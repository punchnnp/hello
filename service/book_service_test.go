package service

import (
	"hello/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	type BookResponse struct {
		BookID      int    `db:"book_id"`
		Name        string `db:"book_name"`
		Description string `db:"book_desc"`
	}

	bookRepo := repository.NewBookRepositoryMock()
	bookService := NewBookService(bookRepo)

	books, _ := bookService.GetAllBooks()
	expected := []BookResponse{
		{BookID: 1, Name: "First Book", Description: "Tell something about this book"},
		{BookID: 2, Name: "Second Book", Description: "This book is about how to cook"},
	}

	assert.Equal(t, expected, books)
}
