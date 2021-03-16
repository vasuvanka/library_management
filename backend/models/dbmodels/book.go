package dbmodels

import "github.com/globalsign/mgo/bson"

type Book struct {
	Name string `bson:"name" json:"name"`
	Author string `bson:"author" json:"author"`
	Copies int `bson:"copies" json:"copies"`
	ID bson.ObjectId `bson:"_id" json:"id"`
}