package collections_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/services/database/collections"
)

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
	var mongo collections.Mongo

	err := mongo.Index()
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}