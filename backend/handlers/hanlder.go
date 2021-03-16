package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vasuvanka/library_management/backend/models"
)

func SendError(ctx *fiber.Ctx,status int,message string) error {
	return ctx.Status(status).JSON(models.GeneralResponse{
		Status:  status,
		Message: message,
	})
}

func SendOk(ctx *fiber.Ctx,data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(data)
}

func SendCreated(ctx *fiber.Ctx,data interface{}) error {
	return ctx.Status(http.StatusCreated).JSON(data)
}

