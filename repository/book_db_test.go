package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		name         string
		mockFunc     func(sqlmock.Sqlmock)
		expectResult bool
		expectError  bool
	}{
		{
			name: "success",
			mockFunc: func(s sqlmock.Sqlmock) {
				rows := s.NewRows([]string{"book_id", "book_name", "book_desc"}).
					AddRow(1, "John", "Doe").
					AddRow(2, "name change", "desc change").
					AddRow(5, "name change", "desc change").
					AddRow(6, "name change", "desc change")
				s.ExpectQuery("SELECT").WillReturnRows(rows)
			},
			expectResult: true,
			expectError:  false,
		},
		{
			name: "fail",
			mockFunc: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT").WillReturnError(errors.New("Fail to get data from database"))
			},
			expectResult: false,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error '%s' when openning database connection", err)
			}
			defer db.Close()
			tt.mockFunc(mock)

			bookRepository := NewBookRepositoryDB(db)
			b, err := bookRepository.GetAll()
			haveResult := b != nil
			haveError := err != nil
			if tt.expectResult != haveResult {
				t.Errorf("expect result %v but %v", tt.expectResult, haveResult)
			}
			if tt.expectError != haveError {
				t.Errorf("expect error %v but %v", tt.expectResult, haveResult)
			}
		})
	}
}

func TestGetById(t *testing.T) {
	tests := []struct {
		name         string
		mockFunc     func(sqlmock.Sqlmock)
		id           int
		expectResult bool
		expectError  bool
	}{
		{
			name: "found book",
			mockFunc: func(s sqlmock.Sqlmock) {
				rows := s.NewRows([]string{"book_id", "book_name", "book_desc"}).
					AddRow(2, "new book", "new description")
				s.ExpectQuery("SELECT").WillReturnRows(rows)
			},
			id:           2,
			expectResult: true,
			expectError:  false,
		},
		{
			name: "not found book",
			mockFunc: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT").WillReturnError(errors.New("Fail to get data"))
			},
			id:           2,
			expectResult: false,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error '%s' when openning database connection", err)
			}
			defer db.Close()
			tt.mockFunc(mock)

			bookRepository := NewBookRepositoryDB(db)
			b, err := bookRepository.GetById(tt.id)
			haveResult := b != nil
			haveError := err != nil
			if tt.expectResult != haveResult {
				t.Errorf("expect result %v but %v", tt.expectResult, haveResult)
			}
			if tt.expectError != haveError {
				t.Errorf("expect error %v but %v", tt.expectError, haveError)
			}
		})
	}
}

func TestAddBook(t *testing.T) {
	tests := []struct {
		name         string
		mockFunc     func(sqlmock.Sqlmock)
		expectResult bool
		expectError  bool
	}{
		{
			name: "success",
			mockFunc: func(s sqlmock.Sqlmock) {
				rows := s.NewRows([]string{"book_id", "book_name", "book_desc"}).
					AddRow(2, "something book", "describe about you")
				s.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(2, 1))
				s.ExpectQuery("SELECT").WillReturnRows(rows)
			},
			expectResult: true,
			expectError:  false,
		},
		{
			name: "fail",
			mockFunc: func(s sqlmock.Sqlmock) {
				s.ExpectExec("INSERT").WillReturnError(errors.New("cannot add new book"))
				s.ExpectQuery("SELECT").WillReturnError(errors.New("cannot get data"))
			},
			expectResult: false,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error %s when openning database connection", err)
			}
			defer db.Close()
			tt.mockFunc(mock)

			bookRepository := NewBookRepositoryDB(db)
			b, err := bookRepository.AddBook()
			haveResult := b != nil
			haveError := err != nil
			if tt.expectResult != haveResult {
				t.Errorf("expect result %v but %v details %v", tt.expectResult, haveResult, b)
			}
			if tt.expectError != haveError {
				t.Errorf("expect error %v but %v details %v", tt.expectError, haveError, err)
			}
		})
	}
}

func TestUpdateBook(t *testing.T) {
	tests := []struct {
		name         string
		mockFunc     func(sqlmock.Sqlmock)
		id           int
		expectResult bool
		expectError  bool
	}{
		{
			name: "success",
			mockFunc: func(s sqlmock.Sqlmock) {
				rows := s.NewRows([]string{"book_id", "book_name", "book_desc"}).
					AddRow(2, "something", "description")
				s.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(2, 1))
				s.ExpectQuery("SELECT").WillReturnRows(rows)
			},
			id:           2,
			expectResult: true,
			expectError:  false,
		},
		{
			name: "fail",
			mockFunc: func(s sqlmock.Sqlmock) {
				s.ExpectExec("UPDATE").WillReturnError(errors.New("cannot update database"))
				s.ExpectQuery("SELECT").WillReturnError(errors.New("cannot get data"))
			},
			id:           2,
			expectResult: false,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error %s when openning database connection", err)
			}
			defer db.Close()
			tt.mockFunc(mock)

			bookRepository := NewBookRepositoryDB(db)
			b, err := bookRepository.UpdateBook(tt.id)
			haveResult := b != nil
			haveError := err != nil
			if tt.expectResult != haveResult {
				t.Errorf("expect result %v but %v details: %v", tt.expectResult, haveResult, b)
			}
			if tt.expectError != haveError {
				t.Errorf("expect error %v but %v datails: %v", tt.expectError, haveError, err)
			}
		})
	}
}

func TestDeleteBook(t *testing.T) {
	tests := []struct {
		name         string
		mockFunc     func(sqlmock.Sqlmock)
		id           int
		expectResult bool
		expectError  bool
	}{
		{
			name: "success",
			mockFunc: func(s sqlmock.Sqlmock) {
				rows := s.NewRows([]string{"book_id", "book_name", "book_desc"}).
					AddRow(2, "something", "description")
				s.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(2, 1))
				s.ExpectQuery("SELECT").WillReturnRows(rows)
			},
			id:           2,
			expectResult: true,
			expectError:  false,
		},
		{
			name: "fail",
			mockFunc: func(s sqlmock.Sqlmock) {
				s.ExpectExec("DELETE").WillReturnError(errors.New("error when delete data from database"))
				s.ExpectQuery("SELECT").WillReturnError(errors.New("error when get data"))
			},
			id:           2,
			expectResult: false,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error %s when openning database connection", err)
			}
			defer db.Close()
			tt.mockFunc(mock)

			bookRepository := NewBookRepositoryDB(db)
			b, err := bookRepository.DeleteBook(tt.id)
			haveResult := b != ""
			haveError := err != nil
			if tt.expectResult != haveResult {
				t.Errorf("expect result %v but %v details: %v", tt.expectResult, haveResult, b)
			}
			if tt.expectError != haveError {
				t.Errorf("expect error %v but %v details: %v", tt.expectError, haveError, err)
			}
		})
	}
}
