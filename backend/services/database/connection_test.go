package database_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/services/database"
)

func TestNewDatabase(t *testing.T){
	db := database.NewDatabase()
	if db == nil {
		t.Error("Expected database object but got nil")
	}
}
