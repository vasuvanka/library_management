package database

import "github.com/vasuvanka/library_management/backend/models/dbmodels"

type IUser interface {
	GetUserByID(id string) (dbmodels.User,error)
	GetUserByUsername(username string) (dbmodels.User,error)
	CreateUser(user dbmodels.User) (dbmodels.User,error)
	UpdateUser(user dbmodels.User) error
	DeleteUser(id string) error
	UserExist(username string) (bool, error)
	DeleteAllUsers() error
}
