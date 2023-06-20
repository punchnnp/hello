package service

import (
	"errors"
	"hello/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	clt := gomock.NewController(t)
	defer clt.Finish()

	// bookRepo := repository.NewBookRepositoryMock()
	mockGetAll := repository.NewMockBookRepository(clt)
	bookService := NewBookService(mockGetAll)

	tests := []struct {
		name        string
		ret         []repository.Book
		err         error
		expected    []BookResponse
		expectedErr string
	}{
		{
			name: "get all book",
			ret: []repository.Book{
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
			err: nil,
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
			expectedErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gomock.InOrder(
				mockGetAll.EXPECT().GetAll().Return(tt.ret, tt.err),
			)

			books, err := bookService.GetAllBooks()
			assert.EqualValues(t, tt.expected, books)
			if err != nil {
				assert.Equal(t, tt.expectedErr, err.Error())
			}
		})
	}

}

func TestGetBookById(t *testing.T) {
	clt := gomock.NewController(t)
	defer clt.Finish()

	mockGetById := repository.NewMockBookRepository(clt)
	// bookRepo := repository.NewBookRepositoryMock()
	bookService := NewBookService(mockGetById)

	tests := []struct {
		name        string
		input       int
		ret         *repository.Book
		err         error
		expected    *BookResponse
		expectedErr string
	}{
		{
			name:  "get existing book id: 1",
			input: 1,
			ret: &repository.Book{
				BookID:      1,
				Name:        "First Book",
				Description: "Tell something about this book",
			},
			err: nil,
			expected: &BookResponse{
				BookID:      1,
				Name:        "First Book",
				Description: "Tell something about this book",
			},
			expectedErr: "",
		},
		{
			name:  "get existing book id: 2",
			input: 2,
			ret: &repository.Book{
				BookID:      2,
				Name:        "Second Book",
				Description: "This book is about how to cook",
			},
			err: nil,
			expected: &BookResponse{
				BookID:      2,
				Name:        "Second Book",
				Description: "This book is about how to cook",
			},
			expectedErr: "",
		},
		{
			name:        "get not existing book",
			input:       3,
			ret:         nil,
			err:         errors.New("this book id is not exist"),
			expected:    nil,
			expectedErr: "this book id is not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gomock.InOrder(
				mockGetById.EXPECT().GetById(tt.input).Return(tt.ret, tt.err),
			)

			book, err := bookService.GetBookById(tt.input)
			assert.EqualValues(t, tt.expected, book)
			if err != nil {
				assert.Equal(t, tt.expectedErr, err.Error())
			}
		})
	}
}

func TestAddNewBook(t *testing.T) {
	clt := gomock.NewController(t)
	defer clt.Finish()

	mockAddBook := repository.NewMockBookRepository(clt)
	// bookRepo := repository.NewBookRepositoryMock()
	bookService := NewBookService(mockAddBook)

	tests := []struct {
		name        string
		ret         *repository.Book
		err         error
		expected    *BookResponse
		expectedErr string
	}{
		{
			name: "add new book",
			ret: &repository.Book{
				BookID:      3,
				Name:        "Add book",
				Description: "this is new book",
			},
			err: nil,
			expected: &BookResponse{
				BookID:      3,
				Name:        "Add book",
				Description: "this is new book",
			},
			expectedErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gomock.InOrder(
				mockAddBook.EXPECT().AddBook().Return(tt.ret, tt.err),
			)

			book, err := bookService.AddNewBook()
			assert.EqualValues(t, tt.expected, book)
			if err != nil {
				assert.Equal(t, tt.expectedErr, err.Error())
			}
		})
	}
}

func TestUpdateBook(t *testing.T) {
	clt := gomock.NewController(t)
	defer clt.Finish()

	mockUpdate := repository.NewMockBookRepository(clt)
	// bookRepo := repository.NewBookRepositoryMock()
	bookService := NewBookService(mockUpdate)

	tests := []struct {
		name        string
		input       int
		ret         *repository.Book
		err         error
		expected    *BookResponse
		expectedErr string
	}{
		{
			name:  "update existing book id: 1",
			input: 1,
			ret: &repository.Book{
				BookID:      1,
				Name:        "Change name",
				Description: "Change description",
			},
			err: nil,
			expected: &BookResponse{
				BookID:      1,
				Name:        "Change name",
				Description: "Change description",
			},
			expectedErr: "",
		},
		{
			name:  "update existing book id: 2",
			input: 2,
			ret: &repository.Book{
				BookID:      2,
				Name:        "Change name",
				Description: "Change description",
			},
			err: nil,
			expected: &BookResponse{
				BookID:      2,
				Name:        "Change name",
				Description: "Change description",
			},
			expectedErr: "",
		},
		{
			name:        "update not existing book",
			input:       3,
			ret:         nil,
			err:         errors.New("this book id is not exist"),
			expected:    nil,
			expectedErr: "this book id is not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gomock.InOrder(
				mockUpdate.EXPECT().UpdateBook(tt.input).Return(tt.ret, tt.err),
			)

			book, err := bookService.UpdateBook(tt.input)
			assert.EqualValues(t, tt.expected, book)
			if err != nil {
				assert.Equal(t, tt.expectedErr, err.Error())
			}
		})
	}
}

func TestDeleteBook(t *testing.T) {
	clt := gomock.NewController(t)
	defer clt.Finish()

	mockDelete := repository.NewMockBookRepository(clt)
	// bookRepo := repository.NewBookRepositoryMock()
	bookService := NewBookService(mockDelete)

	tests := []struct {
		name        string
		input       int
		ret         string
		err         error
		expected    string
		expectedErr string
	}{
		{
			name:        "delete existing book id: 1",
			input:       1,
			ret:         "this book ID is deleted",
			err:         nil,
			expected:    "this book ID is deleted",
			expectedErr: "",
		},
		{
			name:        "delete existing book id: 2",
			input:       2,
			ret:         "this book D is deleted",
			err:         nil,
			expected:    "this book ID is deleted",
			expectedErr: "",
		},
		{
			name:        "delete not existing book",
			input:       3,
			ret:         "this book id is not exist",
			err:         errors.New("this book id is not exist"),
			expected:    "this book id is not exist",
			expectedErr: "this book id is not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gomock.InOrder(
				mockDelete.EXPECT().DeleteBook(tt.input).Return(tt.ret, tt.err),
			)

			book, err := bookService.DeleteBook(tt.input)
			assert.EqualValues(t, tt.expected, book)
			if err != nil {
				assert.Equal(t, tt.expectedErr, err.Error())
			}
		})
	}
}
