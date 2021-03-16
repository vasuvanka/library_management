package handlers

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gofiber/fiber/v2"
	"github.com/vasuvanka/library_management/backend/controllers"
	"github.com/vasuvanka/library_management/backend/models"
)

func CreateUser(ctx *fiber.Ctx) error  {
	var user models.CreateUser
	if err := ctx.BodyParser(&user); err != nil {
		return SendError(ctx,http.StatusBadRequest,err.Error())
	}

	dbUser, err := controllers.CreateUser(user)
	if err != nil {
		return SendError(ctx,http.StatusInternalServerError,err.Error())
	}
	return SendCreated(ctx,dbUser)
}

func GetUserByID(ctx *fiber.Ctx) error {
	id:=ctx.Params("id","")
	if id == "" {
		return SendError(ctx,http.StatusBadRequest,"missing user id")
	}
	if !bson.IsObjectIdHex(id) {
		return SendError(ctx,http.StatusBadRequest,"invalid user id")
	}

	user, err := controllers.GetUserByID(id)
	if err != nil {
		return SendError(ctx,http.StatusInternalServerError,err.Error())
	}
	return SendCreated(ctx,user)

}
