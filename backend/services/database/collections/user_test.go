package collections_test

import (
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/vasuvanka/library_management/backend/models/dbmodels"
	"github.com/vasuvanka/library_management/backend/services/database/collections"
)

func TestGetUserByID(t *testing.T){
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	_, err = mongo.GetUserByID(bson.NewObjectId().Hex())
	if err != nil && err.Error() != "not found" {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}

func TestGetUserByUsername(t *testing.T){
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	_, err = mongo.GetUserByUsername("test")
	if err != nil && err.Error() != "not found" {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}

func TestUserExist(t *testing.T){
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	_, err = mongo.UserExist("test")
	if err != nil && err.Error() != "not found" {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}

func TestCreateUser(t *testing.T){
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	var user dbmodels.User
	user.FirstName = "first"
	user.LastName = "last"
	user.Username = "test"
	_, err = mongo.CreateUser(user)
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}

func TestUpdateUser(t *testing.T){
	mongo := collections.NewMongo()
	err := mongo.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	var user dbmodels.User
	user.FirstName = "first"
	user.LastName = "last"
	user.Username = "test"
	dbuser, err := mongo.CreateUser(user)
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
	dbuser.FirstName = "first name"
	err = mongo.UpdateUser(dbuser)
	if err != nil {
		t.Errorf("Expected nil but got %s", err.Error())
	}
}
