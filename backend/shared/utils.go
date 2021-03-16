package shared

import (
	"github.com/vasuvanka/library_management/backend/models/dbmodels"
)

func RemoveIndexFromBooks(s []dbmodels.Book, index int) []dbmodels.Book {
    return append(s[:index], s[index+1:]...)
}