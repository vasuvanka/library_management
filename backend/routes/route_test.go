package routes_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/config"
	"github.com/vasuvanka/library_management/backend/routes"
)

func TestNewServer(t *testing.T)  {
	conf := config.NewConfig()
	server := routes.NewServer(conf)

	if server == nil {
		t.Error("Expected Server object but got nil")
	}

	if server.App == nil {
		t.Errorf("Expected app object but got %v",server.App)
	}

	if server.Service != nil {
		t.Errorf("Expected Service object to be nil but got %v",server.Service)
	}
}

func TestServerInit(t *testing.T){
	conf := config.NewConfig()
	server := routes.NewServer(conf)

	err := server.Init()

	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
}

func TestServerBootstrap(t *testing.T){
	conf := config.NewConfig()
	err := conf.Init()
	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	server := routes.NewServer(conf)
	err = server.Init()

	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
	err = server.Bootstrap()

	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}

	err = server.Shutdown()

	if err != nil {
		t.Errorf("Expected nil but got %s",err.Error())
	}
}
