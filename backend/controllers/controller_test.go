package controllers_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/config"
	"github.com/vasuvanka/library_management/backend/controllers"
	"github.com/vasuvanka/library_management/backend/models"
	"github.com/vasuvanka/library_management/backend/models/dbmodels"
	"github.com/vasuvanka/library_management/backend/routes"
)

func initServer() (*routes.Server,error){
	conf := config.NewConfig()
	conf.Init()
	server := routes.NewServer(conf)
	err := server.Init()
	if err != nil {
		return nil,err
	}
	return server,err
}
func TestGetBooks(t *testing.T)  {
	server,err := initServer()
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	controllers.NewController(server.Service.Db,server.Service.Db)

	_, err = controllers.GetBooks(0,5)
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
}


func TestCreateBook(t *testing.T)  {
	server,err := initServer()
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	controllers.NewController(server.Service.Db,server.Service.Db)
	var book dbmodels.Book
	book.Author="test"
	book.Copies=2
	book.Name="test name"
	_, err = controllers.CreateBook(book)
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
}

func TestGetBookByID(t *testing.T)  {
	server,err := initServer()
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	controllers.NewController(server.Service.Db,server.Service.Db)
	var book dbmodels.Book
	book.Author="test"
	book.Copies=2
	book.Name="test name"
	dbBook, err := controllers.CreateBook(book)
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	_, err = controllers.GetBookByID(dbBook.ID.Hex())
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
}

func TestBarrowBook(t *testing.T)  {
	server,err := initServer()
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	controllers.NewController(server.Service.Db,server.Service.Db)
	var book dbmodels.Book
	book.Author="test"
	book.Copies=2
	book.Name="test name"
	dbBook, err := controllers.CreateBook(book)
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	_ = controllers.RemoveAllUsers()
	var user models.CreateUser
	user.Username ="testuser"
	user.FirstName="test firstname"
	user.LastName = "test lastname"
	dbUser, err := controllers.CreateUser(user)
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	err = controllers.BorrowBook(dbBook.ID.Hex(),dbUser.ID.Hex())
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
}

func TestReturnBook(t *testing.T)  {
	server,err := initServer()
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	controllers.NewController(server.Service.Db,server.Service.Db)
	_ = controllers.RemoveAllBooks()
	var book dbmodels.Book
	book.Author="test"
	book.Copies=2
	book.Name="test name"
	dbBook, err := controllers.CreateBook(book)
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	_ = controllers.RemoveAllUsers()
	var user models.CreateUser
	user.Username ="testuser"
	user.FirstName="test firstname"
	user.LastName = "test lastname"
	dbUser, err := controllers.CreateUser(user)
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	err = controllers.BorrowBook(dbBook.ID.Hex(),dbUser.ID.Hex())
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	err = controllers.ReturnBook(dbBook.ID.Hex(),dbUser.ID.Hex())
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}

	updatedBook, err := controllers.GetBookByID(dbBook.ID.Hex())
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}

	if updatedBook.Copies != 2 {
		t.Errorf("Expected 2 but got %d",updatedBook.Copies)
	}
}