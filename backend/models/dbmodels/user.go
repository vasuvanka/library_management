package dbmodels

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Username string `bson:"username" json:"username"`
	ID bson.ObjectId `bson:"_id" json:"id"`
	FirstName string `bson:"firstname" json:"firstname"`
	LastName string `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Books []Book `bson:"books,omitempty" json:"books,omitempty"`
}