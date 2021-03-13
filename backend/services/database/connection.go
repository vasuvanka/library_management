package database

import (
	"github.com/vasuvanka/library_management/backend/services/database/collections"
)

type Database struct {
	Db *collections.Mongo
}

func NewDatabase() *Database {
	return &Database{
		Db: collections.NewMongo(),	
	}
}


