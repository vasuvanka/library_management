package collections

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/vasuvanka/library_management/backend/models/dbmodels"
)

func (mongo *Mongo) GetBookByID(id string) (dbmodels.Book, error)  {
	var book dbmodels.Book
	err := mongo.Books().FindId(bson.ObjectIdHex(id)).One(&book)
	return book,err
}

func (mongo *Mongo) GetBooks(query interface{},skip,limit int) ([]dbmodels.Book, error)  {
	var books []dbmodels.Book
	err := mongo.Books().Find(query).Skip(skip).Limit(limit).All(&books)
	return books, err
}

func (mongo *Mongo) CreateBook(book dbmodels.Book) (dbmodels.Book, error)  {
	book.ID = bson.NewObjectId()
	err := mongo.Books().Insert(book)
	return book, err
}

func (mongo *Mongo) UpdateBook(book dbmodels.Book) error  {
	return mongo.Books().UpdateId(book.ID,book)
}

func (mongo *Mongo) DeleteBook(id string) error  {
	return mongo.Books().RemoveId(bson.ObjectIdHex(id))
}

func (db *Mongo) DeleteAllBooks() error {
	info, err:= db.Books().RemoveAll(bson.M{})
	fmt.Println(info)
	return  err
}