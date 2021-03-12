package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vasuvanka/library_management/backend/config"
	"github.com/vasuvanka/library_management/backend/services"
)

type Server struct {
	Config *config.Config
	App	*fiber.App
	Service *services.Service
}

func NewServer(conf *config.Config) *Server {
	return &Server{
		Config: conf,
		App:    fiber.New(),
		Service: nil,
	}
}

func (server *Server) Init() error {
	server.Service = services.NewService()
	return server.Service.Database.Connect(server.Config.DbURI)
}

func (server *Server) Bootstrap() error {
	return server.App.Listen(server.Config.Port)
}