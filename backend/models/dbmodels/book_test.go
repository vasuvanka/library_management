package dbmodels_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/models/dbmodels"
)

func TestBook(t *testing.T){
	var book dbmodels.Book

	if book.ID.Hex() != "" {
		t.Errorf("Expected id to be '' but got %s",book.ID.Hex())
	}

	if book.Name != "" {
		t.Errorf("Expected name to be '' but got %s", book.Name)
	}

	if book.Copies != 0 {
		t.Errorf("Expected copies to be 0 but got %d", book.Copies)
	}

	if book.Author != "" {
		t.Errorf("Expected author to be '' but got %s", book.Author)
	}
}