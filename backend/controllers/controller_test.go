package controllers_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/config"
	"github.com/vasuvanka/library_management/backend/controllers"
	"github.com/vasuvanka/library_management/backend/routes"
)

func TestNewUserControllers(t *testing.T)  {
	conf := config.NewConfig()

	server := routes.NewServer(conf)
	server.Init()
	controllers.NewUserControllers(server.Service)
	if controllers.iUser == nil {
		t.Errorf("Expected non nil user interface")
	}
}