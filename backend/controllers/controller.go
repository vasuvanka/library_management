package controllers

import (
	"github.com/vasuvanka/library_management/backend/services/database/collections"
)

var iUser collections.IUser
func NewUserControllers(iuser collections.IUser)  {
	iUser = iuser
}