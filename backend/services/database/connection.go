package database

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/vasuvanka/library_management/backend/services/database/collections"
)

type Database struct {
	collections.Mongo
}

func (database *Database) Connect(URI string) error  {
	session, err := mgo.DialWithTimeout(URI,time.Second*10)
	if err != nil {
		return err
	}
	database.Mongo.Session = *session.Clone()
	return nil
}
