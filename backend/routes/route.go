package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vasuvanka/library_management/backend/config"
	"github.com/vasuvanka/library_management/backend/controllers"
	"github.com/vasuvanka/library_management/backend/services"
	"github.com/vasuvanka/library_management/backend/services/database/collections"
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
	server.Service.Db = collections.NewMongo()
	return server.Service.Db.Connect(server.Config.DbURI)
}

func (server *Server) Routes() {
	UserRoutes(server)
	BookRoutes(server)
}

func (server *Server) Controllers() {
	controllers.NewController(server.Service.Db,server.Service.Db)
}

func (server *Server) Bootstrap() error {
	log.Println("Server listening on PORT : "+ server.Config.Port)
	return server.App.Listen(fmt.Sprintf(":%s",server.Config.Port))
}

func (server *Server) Shutdown() error {
	log.Println("graceful shutdown")
	return server.App.Shutdown()
}
