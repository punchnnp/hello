package handler_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	"hello/handler"
	"hello/repository"
	"hello/service"

	"github.com/stretchr/testify/assert"
)

func TestIntegGetAll(t *testing.T) {
	tests := []struct {
		name         string
		description  string
		route        string
		expectedCode int
		expected     string
	}{
		{
			name:         "get all book",
			description:  "GET HTTP 200 status",
			route:        "/api/book",
			expectedCode: 200,
			expected:     `[{"BookID":1,"Name":"First Book","Description":"Tell something about this book"},{"BookID":2,"Name":"Second Book","Description":"This book is about how to cook"}]`,
		},
		{
			name:         "get all book wrong route",
			description:  "GET HTTP 404 status when route is not exists",
			route:        "/api/something",
			expectedCode: 404,
			expected:     "Cannot GET /api/something",
		},
	}

	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	app := fiber.New()
	app.Get("/api/book", bookHandler.GetAllBooks)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.route, nil)
			resp, _ := app.Test(req, 1)
			defer resp.Body.Close()

			if assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description) {
				body, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tt.expected, string(body))
			}
		})
	}

}

func TestIntegGetBookById(t *testing.T) {
	tests := []struct {
		name         string
		description  string
		route        string
		expectedCode int
		expected     string
	}{
		{
			name:         "get existing book id: 1",
			description:  "GET HTTP 200 status",
			route:        "/api/book/1",
			expectedCode: 200,
			expected:     `{"BookID":1,"Name":"First Book","Description":"Tell something about this book"}`,
		},
		{
			name:         "get existing book id: 2",
			description:  "GET HTTP 200 status",
			route:        "/api/book/2",
			expectedCode: 200,
			expected:     `{"BookID":2,"Name":"Second Book","Description":"This book is about how to cook"}`,
		},
		{
			name:         "book is not exists",
			description:  "GET HTTP 200 status",
			route:        "/api/book/3",
			expectedCode: 200,
			expected:     "this book id is not exist",
		},
		{
			name:         "wrong route",
			description:  "GET HTTP 404 status",
			route:        "/api/book",
			expectedCode: 404,
			expected:     "Cannot GET /api/book",
		},
	}

	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	app := fiber.New()
	app.Get("/api/book/:id", bookHandler.GetBookById)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.route, nil)
			resp, _ := app.Test(req, 1)
			defer resp.Body.Close()

			if assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description) {
				body, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tt.expected, string(body))
			}
		})
	}

}

func TestIntegAddNewBook(t *testing.T) {
	tests := []struct {
		name         string
		description  string
		route        string
		expectedCode int
		expected     string
	}{
		{
			name:         "add new book",
			description:  "GET HTTP 200 status",
			route:        "/api/book",
			expectedCode: 200,
			expected:     `{"BookID":3,"Name":"Add book","Description":"this is new book"}`,
		},
		{
			name:         "add new book wrong route",
			description:  "GET HTTP 404 status",
			route:        "/api/something",
			expectedCode: 404,
			expected:     "Cannot POST /api/something",
		},
	}

	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	app := fiber.New()
	app.Post("/api/book", bookHandler.AddNewBook)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", tt.route, nil)
			resp, _ := app.Test(req, 1)
			defer resp.Body.Close()

			if assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description) {
				body, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tt.expected, string(body))
			}
		})
	}
}

func TestIntegUpdateBook(t *testing.T) {
	tests := []struct {
		name         string
		description  string
		route        string
		expectedCode int
		expected     string
	}{
		{
			name:         "update existing book id: 1",
			description:  "GET HTTP 200 status",
			route:        "/api/book/1",
			expectedCode: 200,
			expected:     `{"BookID":1,"Name":"Change name","Description":"Change description"}`,
		},
		{
			name:         "update existing book id: 2",
			description:  "GET HTTP 200 status",
			route:        "/api/book/2",
			expectedCode: 200,
			expected:     `{"BookID":2,"Name":"Change name","Description":"Change description"}`,
		},
		{
			name:         "update not existing book",
			description:  "GET HTTTO 200 status",
			route:        "/api/book/3",
			expectedCode: 200,
			expected:     "this book id is not exist",
		},
		{
			name:         "update with wrong route",
			description:  "GET HTTP 404 status",
			route:        "/api/book",
			expectedCode: 404,
			expected:     "Cannot POST /api/book",
		},
	}

	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	app := fiber.New()
	app.Post("/api/book/:id", bookHandler.UpdateBook)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", tt.route, nil)
			resp, _ := app.Test(req, 1)
			defer resp.Body.Close()

			if assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description) {
				body, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tt.expected, string(body))
			}
		})
	}
}

func TestIntegDeleteBook(t *testing.T) {
	tests := []struct {
		name         string
		description  string
		route        string
		expectedCode int
		expected     string
	}{
		{
			name:         "delete existing book id: 1",
			description:  "GET HTTP 200 status",
			route:        "/api/book/1",
			expectedCode: 200,
			expected:     "this book ID is deleted",
		},
		{
			name:         "delete existing book id: 2",
			description:  "GET HTTP 200 status",
			route:        "/api/book/2",
			expectedCode: 200,
			expected:     "this book ID is deleted",
		},
		{
			name:         "delete not existing book",
			description:  "GET HTTP 200 status",
			route:        "/api/book/3",
			expectedCode: 200,
			expected:     "this book id is not exist",
		},
		{
			name:         "delete with wrong route",
			description:  "GET HTTP 404 status",
			route:        "/api/book",
			expectedCode: 404,
			expected:     "Cannot DELETE /api/book",
		},
	}

	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	app := fiber.New()
	app.Delete("/api/book/:id", bookHandler.DeleteBook)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", tt.route, nil)
			resp, _ := app.Test(req, 1)
			defer resp.Body.Close()

			if assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description) {
				body, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tt.expected, string(body))
			}
		})
	}
}
