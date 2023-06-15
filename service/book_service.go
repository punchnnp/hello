package service

import (
	"hello/repository"
	"log"
)

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) bookService {
	return bookService{bookRepo: bookRepo}
}

func (s bookService) GetAllBooks() ([]BookResponse, error) {
	books, err := s.bookRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	results := []BookResponse{}
	for _, book := range books {
		result := BookResponse{
			BookID:      book.BookID,
			Name:        book.Name,
			Description: book.Description,
		}
		results = append(results, result)
	}

	return results, nil
}

func (s bookService) GetBookById(id int) (*BookResponse, error) {
	book, err := s.bookRepo.GetById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := BookResponse{
		BookID:      book.BookID,
		Name:        book.Name,
		Description: book.Description,
	}

	return &result, nil
}

func (s bookService) AddNewBook() (*BookResponse, error) {
	book, err := s.bookRepo.AddBook()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result := BookResponse{
		BookID:      book.BookID,
		Name:        book.Name,
		Description: book.Description,
	}

	return &result, nil
}

func (s bookService) UpdateBook(id int) (*BookResponse, error) {
	book, err := s.bookRepo.UpdateBook(id)
	if err != nil {
		log.Println(err)
	}

	result := BookResponse{
		BookID:      book.BookID,
		Name:        book.Name,
		Description: book.Description,
	}

	return &result, nil
}

func (s bookService) DeleteBook(id int) (string, error) {
	book, err := s.bookRepo.DeleteBook(id)
	if err != nil {
		log.Println(err)
	}

	return book, nil
}
