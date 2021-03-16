package database

import "github.com/vasuvanka/library_management/backend/models/dbmodels"

type IBook interface {
	GetBooks(query interface{},skip,limit int) ([]dbmodels.Book, error)
	GetBookByID(id string) (dbmodels.Book, error)
	CreateBook(book dbmodels.Book) (dbmodels.Book, error)
	UpdateBook(book dbmodels.Book) error
	DeleteBook(id string) error
	DeleteAllBooks() error
}