package collections

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/vasuvanka/library_management/backend/constants"
)

//Mongo - mongo struct
type Mongo struct {
	Session mgo.Session
}

func NewMongo() *Mongo {
	return &Mongo{}
}

func (mongo *Mongo) Connect(URI string) error {
	session, err := mgo.DialWithTimeout(URI,time.Second*10)
	if err != nil {
		return err
	}
	mongo.Session = *session.Clone()
	return nil
}

func (mongo *Mongo) Books() *mgo.Collection {
	return mongo.Session.DB(constants.Database).C(constants.CollectionBooks)
}

func (mongo *Mongo) Users() *mgo.Collection {
	return mongo.Session.DB(constants.Database).C(constants.CollectionUsers)
}

func (mongo *Mongo) Index() error  {
	idx := mgo.Index{
		Key:      []string{"username"},
		Unique:   true,
		DropDups: true,
		Name:     "username_index",
	}
	// user indexes
	if err := mongo.Users().EnsureIndex(idx); err != nil {
		return err
	}

	idx = mgo.Index{
		Key:  []string{"books"},
		Name: "user_books_index",
	}
	// user indexes
	if err := mongo.Users().EnsureIndex(idx); err != nil {
		return err
	}

	idx = mgo.Index{
		Key:  []string{"name","author","count"},
		Name: "books_name_count_index",
	}

	return mongo.Books().EnsureIndex(idx)
}