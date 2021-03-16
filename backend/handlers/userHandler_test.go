package handlers_test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/vasuvanka/library_management/backend/handlers"
)

func setupApp() *fiber.App {
	app := fiber.New()

	return app
}

func TestGetBooksValidSkipAndLimit(t *testing.T){
	query := make(fiber.Map)
	query["skip"]="0"
	query["limit"]="10"
	ctx := fiber.Ctx{
		
	}
	err := handlers.GetBooks(&ctx)
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
}

