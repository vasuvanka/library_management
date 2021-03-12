package constants_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/constants"
)

func TestConstants(t *testing.T)  {
	if constants.Database != "library" {
		t.Errorf("Expected library but got %s",constants.Database)
	}

	if constants.CollectionBooks != "books" {
		t.Errorf("Expected books but got %s",constants.CollectionBooks)
	}

	if constants.CollectionUsers != "users" {
		t.Errorf("Expected users but got %s",constants.CollectionUsers)
	}
}