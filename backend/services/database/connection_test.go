package database_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/services/database"
)
func TestConnect(t *testing.T)  {
	var database database.Database
	err := database.Connect("mongodb://localhost:27017/test")
	if err != nil {
		t.Errorf("Expected nil but got error %s", err.Error())
	}
}