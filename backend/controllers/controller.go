package controllers

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/vasuvanka/library_management/backend/models"
	"github.com/vasuvanka/library_management/backend/models/dbmodels"
	"github.com/vasuvanka/library_management/backend/services/database"
	"github.com/vasuvanka/library_management/backend/shared"
)

var iUser database.IUser
var iBook database.IBook
func NewController(iuser database.IUser,ibook database.IBook)  {
	iUser = iuser
	iBook = ibook
}

func GetBooks(skip,limit int) ([]dbmodels.Book,error) {
	return iBook.GetBooks(bson.M{
		"copies": bson.M{
			"$gt":0,
		},
	},skip,limit)
}

func GetBookByID(id string) (dbmodels.Book, error){
	return iBook.GetBookByID(id)
}

func CreateBook(book dbmodels.Book) (dbmodels.Book,error) {
	return iBook.CreateBook(book)
}

func BorrowBook(bookID,userID string) error {
	dbUser, err := iUser.GetUserByID(userID)
	if err != nil {
		return err
	}
	if len(dbUser.Books) > 2 {
		return errors.New("Borrowed limit reached")
	}
	for _, userBook := range dbUser.Books {
		if userBook.ID.Hex() == bookID {
			return errors.New("This book is already borrowed by you")
		}
	}
	dbBook, err := iBook.GetBookByID(bookID)
	if err != nil {
		return err
	}
	if dbBook.Copies == 0 {
		return errors.New("No copies available")
	}
	dbBook.Copies -=1 
	err = iBook.UpdateBook(dbBook)
	if err != nil {
		return err
	}
	dbUser.Books = append(dbUser.Books, dbBook)
	err = iUser.UpdateUser(dbUser)
	if err != nil {
		dbBook.Copies+=1
		_ = iBook.UpdateBook(dbBook)
	}
	return err
}

func ReturnBook(bookID,userID string) error {
	dbUser, err := iUser.GetUserByID(userID)
	if err != nil {
		return err
	}
	found := false
	foundAt := -1
	for index, userBook := range dbUser.Books {
		if userBook.ID.Hex() == bookID {
			found= true
			foundAt = index
			break
		}
	}
	if !found {
		return errors.New("This book is not borrowed by you")
	}

	dbBook, err := iBook.GetBookByID(bookID)
	dbBook.Copies+=1
	if err != nil {
		return err
	}
	
	dbUser.Books = shared.RemoveIndexFromBooks(dbUser.Books,foundAt)
	err = iUser.UpdateUser(dbUser)
	if err != nil {
		dbBook.Copies-=1
		_ = iBook.UpdateBook(dbBook)
	}
	err = iBook.UpdateBook(dbBook)
	return err
}

func CreateUser(user models.CreateUser) (dbmodels.User,error)  {
	exist, err := iUser.UserExist(user.Username)
	if err != nil {
		return dbmodels.User{},err
	}
	if exist {
		return dbmodels.User{},errors.New("user already exists")
	}
	var dbUser dbmodels.User
	dbUser.FirstName = user.FirstName
	dbUser.LastName = user.LastName
	dbUser.Username = user.Username
	return iUser.CreateUser(dbUser)
}

func GetUserByID(id string) (dbmodels.User,error)  {
	return iUser.GetUserByID(id)
}

func RemoveAllUsers() error {
	return iUser.DeleteAllUsers()
}

func RemoveAllBooks() error {
	return iBook.DeleteAllBooks()
}
