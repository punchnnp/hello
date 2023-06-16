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

	tests := []struct {
		name     string
		expected []BookResponse
	}{
		{
			name: "get all book",
			expected: []BookResponse{
				{BookID: 1, Name: "First Book", Description: "Tell something about this book"},
				{BookID: 2, Name: "Second Book", Description: "This book is about how to cook"},
			},
		},
	}

	bookRepo := repository.NewBookRepositoryMock()
	bookService := NewBookService(bookRepo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			books, _ := bookService.GetAllBooks()
			assert.EqualValues(t, tt.expected, books)
		})
	}

}
