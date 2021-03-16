package collections

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/vasuvanka/library_management/backend/models/dbmodels"
)

func (db *Mongo) GetUserByID(id string) (dbmodels.User,error) {
	var user dbmodels.User
	err := db.Users().FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (db *Mongo) GetUserByUsername(username string) (dbmodels.User,error) {
	var user dbmodels.User
	err := db.Users().Find(bson.M{
		"username":username,
	}).One(&user)
	return user, err
}

func (db *Mongo) UserExist(username string) (bool,error) {
	count, err := db.Users().Find(bson.M{
		"username":username,
	}).Count()
	if err != nil {
		return false, err
	}
	return count>=1, err
}

func (db *Mongo) CreateUser(user dbmodels.User) (dbmodels.User,error) {
	user.ID = bson.NewObjectId()
	err := db.Users().Insert(user)
	return user, err
}

func (db *Mongo) UpdateUser(user dbmodels.User) error {
	return db.Users().UpdateId(user.ID,user)
}

func (db *Mongo) DeleteUser(id string) error {
	return db.Users().RemoveId(id)
}

func (db *Mongo) DeleteAllUsers() error {
	info, err := db.Users().RemoveAll(bson.M{})
	fmt.Println(info)
	return err
}