package shared_test

import (
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/vasuvanka/library_management/backend/models/dbmodels"
	"github.com/vasuvanka/library_management/backend/shared"
)

func TestRemoveIndex(t *testing.T){
	id := bson.NewObjectId()
	list := []dbmodels.Book{
		dbmodels.Book{
			Name:   "",
			Author: "",
			Copies: 0,
			ID:     id,
		},
		dbmodels.Book{
			Name:   "",
			Author: "",
			Copies: 0,
			ID:     bson.NewObjectId(),
		},
	}
	newList := shared.RemoveIndexFromBooks(list,0)
	found := false
	for _, item := range newList {
		if id.Hex() == item.ID.Hex() {
			found = true
		}
	}
	if found {
		t.Errorf("%s should be removed from list",id.Hex())
	}
}