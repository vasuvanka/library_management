package dbmodels

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Username string `bson:"username"`
	ID bson.ObjectId `bson:"_id"`
	FirstName string `bson:"firstname"`
	LastName string `bson:"lastname"`
	Books []bson.ObjectId `bson:"books"`
}