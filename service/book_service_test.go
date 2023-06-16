package service

import (
	"hello/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		name     string
		expected []BookResponse
	}{
		{
			name: "get all book",
			expected: []BookResponse{
				{
					BookID:      1,
					Name:        "First Book",
					Description: "Tell something about this book",
				},
				{
					BookID:      2,
					Name:        "Second Book",
					Description: "This book is about how to cook",
				},
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

func TestGetBookById(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected *BookResponse
	}{
		{
			name:  "get existing book id: 1",
			input: 1,
			expected: &BookResponse{
				BookID:      1,
				Name:        "First Book",
				Description: "Tell something about this book",
			},
		},
		{
			name:  "get existing book id: 2",
			input: 2,
			expected: &BookResponse{
				BookID:      2,
				Name:        "Second Book",
				Description: "This book is about how to cook",
			},
		},
		{
			name:     "get not existing book",
			input:    3,
			expected: nil,
		},
	}

	bookRepo := repository.NewBookRepositoryMock()
	bookService := NewBookService(bookRepo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			book, _ := bookService.GetBookById(tt.input)
			assert.EqualValues(t, tt.expected, book)

		})
	}
}

func TestAddNewBook(t *testing.T) {
	tests := []struct {
		name     string
		expected *BookResponse
	}{
		{
			name: "add new book",
			expected: &BookResponse{
				BookID:      3,
				Name:        "Add book",
				Description: "this is new book",
			},
		},
	}

	bookRepo := repository.NewBookRepositoryMock()
	bookService := NewBookService(bookRepo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			book, _ := bookService.AddNewBook()
			assert.EqualValues(t, tt.expected, book)
		})
	}
}

func TestUpdateBook(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected *BookResponse
	}{
		{
			name:  "update existing book id: 1",
			input: 1,
			expected: &BookResponse{
				BookID:      1,
				Name:        "Change name",
				Description: "Change description",
			},
		},
		{
			name:  "update existing book id: 2",
			input: 2,
			expected: &BookResponse{
				BookID:      2,
				Name:        "Change name",
				Description: "Change description",
			},
		},
		{
			name:     "update not existing book",
			input:    3,
			expected: nil,
		},
	}

	bookRepo := repository.NewBookRepositoryMock()
	bookService := NewBookService(bookRepo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			book, _ := bookService.UpdateBook(tt.input)
			assert.EqualValues(t, tt.expected, book)
		})
	}
}

func TestDeleteBook(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "delete existing book id: 1",
			input:    1,
			expected: "this book ID is deleted",
		},
		{
			name:     "delete existing book id: 2",
			input:    2,
			expected: "this book ID is deleted",
		},
		{
			name:     "delete not existing book",
			input:    3,
			expected: "this book id is not exist",
		},
	}

	bookRepo := repository.NewBookRepositoryMock()
	bookService := NewBookService(bookRepo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			book, _ := bookService.DeleteBook(tt.input)
			assert.EqualValues(t, tt.expected, book)
		})
	}
}
