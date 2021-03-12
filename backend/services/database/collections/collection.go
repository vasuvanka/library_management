package collections

import (
	"github.com/globalsign/mgo"
	"github.com/vasuvanka/library_management/backend/constants"
)

//Mongo - mongo struct
type Mongo struct {
	Session mgo.Session
}

func (mongo *Mongo) Books() *mgo.Collection {
	return mongo.Session.DB(constants.Database).C(constants.CollectionBooks)
}

func (mongo *Mongo) Users() *mgo.Collection {
	return mongo.Session.DB(constants.Database).C(constants.CollectionUsers)
}

func (mongo *Mongo) Index() error  {
	idx := mgo.Index{
		Key:      []string{"email"},
		Unique:   true,
		DropDups: true,
		Name:     "user_email_index",
	}
	// user indexes
	if err := mongo.Users().EnsureIndex(idx); err != nil {
		return err
	}

	idx = mgo.Index{
		Key:  []string{"removed", "name"},
		Name: "user_name_index",
	}
	// user indexes
	if err := mongo.Users().EnsureIndex(idx); err != nil {
		return err
	}

	idx = mgo.Index{
		Key:  []string{"removed", "name","available"},
		Name: "books_name_available_index",
	}

	return mongo.Books().EnsureIndex(idx)
}