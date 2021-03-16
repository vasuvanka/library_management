package collections_test

import (
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/vasuvanka/library_management/backend/models/dbmodels"
	"github.com/vasuvanka/library_management/backend/services/database/collections"
)

func TestGetBookByID(t *testing.T)  {
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	_, err = mongo.GetBookByID(bson.NewObjectId().Hex())
	if err != nil && err.Error() != "not found" {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}

func TestGetBooks(t *testing.T)  {
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	_, err = mongo.GetBooks(bson.M{},0,5)
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}

func TestCreateBook(t *testing.T)  {
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	var book dbmodels.Book
	book.Author = "test"
	book.Copies=2
	book.Name="Test Book"
	_, err = mongo.CreateBook(book)
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}

func TestUpdateBook(t *testing.T)  {
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	var book dbmodels.Book
	book.Author = "test"
	book.Copies=2
	book.Name="Test Book"
	dbBook, err := mongo.CreateBook(book)
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	dbBook.Copies = 3
	err = mongo.UpdateBook(dbBook)
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}

func TestDeleteBook(t *testing.T)  {
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	var book dbmodels.Book
	book.Author = "test"
	book.Copies=2
	book.Name="Test Book"
	dbBook, err := mongo.CreateBook(book)
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	err = mongo.DeleteBook(dbBook.ID.Hex())
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}

func TestDeleteAllBooks(t *testing.T)  {
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	var book dbmodels.Book
	book.Author = "test"
	book.Copies=2
	book.Name="Test Book"
	dbBook, err := mongo.CreateBook(book)
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	err = mongo.DeleteAllBooks()
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}

	_, err = mongo.GetBookByID(dbBook.ID.Hex())
	if err != nil && err.Error() != "not found" {
		t.Errorf("Expected not found but got %s", err.Error())
	}
}
