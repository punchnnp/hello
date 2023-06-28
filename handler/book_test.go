package handler_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"hello/handler"
	"hello/repository"
	"hello/service"
)

func TestGetAll(t *testing.T) {

	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "GET HTTP status 200",
			route:        "/api/book",
			expectedCode: 200,
		},
		{
			description:  "GET HTTP status 404, when route is not exists",
			route:        "/api/something",
			expectedCode: 404,
		},
	}

	app := fiber.New()
	app.Get("/api/book", bookHandler.GetAllBooks)

	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.route, nil)
		resp, _ := app.Test(req, 1)

		assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
	}

}

func TestGetBookById(t *testing.T) {

	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "GET HTTP status 200",
			route:        "/api/book/1",
			expectedCode: 200,
		},
		{
			description:  "GET HTTP status 404, when route/bookID is not exist",
			route:        "/api/book/1/book",
			expectedCode: 404,
		},
	}

	app := fiber.New()
	app.Get("/api/book/:id", bookHandler.GetBookById)

	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.route, nil)
		resp, _ := app.Test(req, 1)

		assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
	}

}

func TestAddNewBook(t *testing.T) {

	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "GET HTTP 200 status",
			route:        "/api/book",
			expectedCode: 200,
		},
		{
			description:  "GET HTTP 404 status when route is not exists",
			route:        "/api/something",
			expectedCode: 404,
		},
	}

	app := fiber.New()
	app.Post("/api/book", bookHandler.AddNewBook)

	for _, tt := range tests {
		req := httptest.NewRequest("POST", tt.route, nil)
		resp, _ := app.Test(req, 1)

		assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)

	}
}

func TestUpdateBook(t *testing.T) {

	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "GET HTTP 200 status",
			route:        "/api/book/2",
			expectedCode: 200,
		},
		{
			description:  "GET HTTP 404 status when route is not exists",
			route:        "/api/book",
			expectedCode: 404,
		},
	}

	app := fiber.New()
	app.Post("/api/book/:id", bookHandler.AddNewBook)

	for _, tt := range tests {
		req := httptest.NewRequest("POST", tt.route, nil)
		resp, _ := app.Test(req, 1)

		assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}

func TestDeleteBook(t *testing.T) {

	bookRepositoryMock := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepositoryMock)
	bookHandler := handler.NewBookHandler(bookService)

	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "GET HTTP 200 status",
			route:        "/api/book/2",
			expectedCode: 200,
		},
		{
			description:  "GET HTTP 404 status when route is not exists",
			route:        "/api/book",
			expectedCode: 404,
		},
	}

	app := fiber.New()
	app.Delete("/api/book/:id", bookHandler.DeleteBook)

	for _, tt := range tests {
		req := httptest.NewRequest("DELETE", tt.route, nil)
		resp, _ := app.Test(req, 1)

		assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
	}
}
