package collections_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/services/database/collections"
)

func TestGetMongo(t *testing.T){
	mongo := collections.NewMongo()
	if mongo == nil {
		t.Error("Expected mongo object but got nil")
	}
}

func TestBooksCollection(t *testing.T){
	var mongo collections.Mongo

	coll := mongo.Books()
	if coll == nil {
		t.Error("Expected nil books collection")
	}
}

func TestUsersCollection(t *testing.T){
	var mongo collections.Mongo

	coll := mongo.Users()
	if coll == nil {
		t.Error("Expected nil users collection")
	}
}

func TestIndexCollections(t *testing.T){
	mongo := collections.NewMongo()

	err := mongo.Index()
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}


func TestConnect(t *testing.T)  {
	db := collections.NewMongo()
	err := db.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got error %s", err.Error())
	}
}