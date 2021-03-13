package dbmodels_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/models/dbmodels"
)

func TestUser(t *testing.T){
	var user dbmodels.User
	if user.ID.Hex() != "" {
		t.Errorf("Expected empty string but got %s",user.ID.Hex())
	}
	if len(user.Books) != 0 {
		t.Errorf("Expected nil but got %v",user.Books)
	}

	if user.FirstName != "" {
		t.Errorf("Expected '' but got %s",user.FirstName)
	}

	if user.LastName != "" {
		t.Errorf("Expected '' but got %s",user.LastName)
	}

	if user.Username != "" {
		t.Errorf("Expected '' but got %s",user.Username)
	}
}