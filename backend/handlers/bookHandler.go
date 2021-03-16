package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/globalsign/mgo/bson"
	"github.com/gofiber/fiber/v2"
	"github.com/vasuvanka/library_management/backend/controllers"
	"github.com/vasuvanka/library_management/backend/models"
	"github.com/vasuvanka/library_management/backend/models/dbmodels"
)


func GetBooks(ctx *fiber.Ctx) error  {
	skip, err := strconv.Atoi(ctx.Query("skip","0"))
	if err != nil {
		return SendError(ctx,http.StatusBadRequest,err.Error())
	}
	limit, err := strconv.Atoi(ctx.Query("limit","10"))
	if err != nil {
		return SendError(ctx,http.StatusBadRequest,err.Error())
	}

	books, err := controllers.GetBooks(skip,limit)

	if err != nil {
		return SendError(ctx,http.StatusInternalServerError,err.Error())
	}
	
	return SendOk(ctx,books)
}

func CreateBook(ctx *fiber.Ctx) error {
	var book models.CreateBook
	if err := ctx.BodyParser(&book); err != nil {
		return SendError(ctx,http.StatusBadRequest,err.Error())
	}
	fmt.Println(book)
	dbBook, err := controllers.CreateBook(dbmodels.Book{
		Name:   book.Name,
		Author: book.Author,
		Copies: book.Copies,
	})
	if err != nil {
		return SendError(ctx,http.StatusInternalServerError,err.Error())
	}
	return SendCreated(ctx,dbBook)
}

func BorrowBook(ctx *fiber.Ctx) error  {
	id:=ctx.Query("id","")
	bookID := ctx.Params("id","")
	if id == "" {
		return SendError(ctx,http.StatusBadRequest,"missing user id")
	}
	if !bson.IsObjectIdHex(id) {
		return SendError(ctx,http.StatusBadRequest,"invalid user id")
	}
	if bookID == "" {
		return SendError(ctx,http.StatusBadRequest,"invalid book id")
	}
	if !bson.IsObjectIdHex(bookID) {
		return SendError(ctx,http.StatusBadRequest,"invalid book id")
	}
	err := controllers.BorrowBook(bookID,id)
	if err != nil {
		return SendError(ctx,http.StatusInternalServerError,err.Error())
	}
	return SendOk(ctx,models.GeneralResponse{
		Status:  200,
		Message: "success",
	})
}

func ReturnBook(ctx *fiber.Ctx) error  {
	id:=ctx.Query("id","")
	bookID := ctx.Params("id","")
	if id == "" {
		return SendError(ctx,http.StatusBadRequest,"missing user id")
	}
	if !bson.IsObjectIdHex(id) {
		return SendError(ctx,http.StatusBadRequest,"invalid user id")
	}
	if bookID == "" {
		return SendError(ctx,http.StatusBadRequest,"invalid book id")
	}
	if !bson.IsObjectIdHex(bookID) {
		return SendError(ctx,http.StatusBadRequest,"invalid book id")
	}
	err := controllers.ReturnBook(bookID,id)
	if err != nil {
		return SendError(ctx,http.StatusInternalServerError,err.Error())
	}
	return SendOk(ctx,models.GeneralResponse{
		Status:  200,
		Message: "success",
	})
}