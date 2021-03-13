package dbmodels

import "github.com/globalsign/mgo/bson"

type Book struct {
	Name string `bson:"name"`
	Author string `bson:"author"`
	Copies int `bson:"copies"`
	ID bson.ObjectId `bson:"_id"`
}