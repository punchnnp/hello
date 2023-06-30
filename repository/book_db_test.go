package repository

import (
	"testing"

	// "fmt"

	"github.com/DATA-DOG/go-sqlmock"
	// "github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s' when opening database connection", err)
	}
	defer db.Close()

	rows := mock.NewRows([]string{"book_id", "book_name", "book_desc"}).
		AddRow(1, "John", "Doe").
		AddRow(2, "name change", "desc change").
		AddRow(5, "name change", "desc change").
		AddRow(6, "name change", "desc change")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	bookRepository := NewBookRepositoryDB(db)
	_, err = bookRepository.GetAll()
	if err != nil {
		t.Errorf("error while get data from db %s", err)
	}

}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error %s when opening database connection", err)
	}
	defer db.Close()

	rows := mock.NewRows([]string{"book_id", "book_name", "book_desc"}).
		AddRow(1, "something", "description").
		AddRow(5, "blahh", "telling blah")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	bookRepository := NewBookRepositoryDB(db)
	_, err = bookRepository.GetById(2)
	if err != nil {
		t.Errorf("error while get data from db %s", err)
	}

}
