package routes

import "github.com/vasuvanka/library_management/backend/handlers"


func UserRoutes(server *Server){
	api := server.App.Group("/api")
	api.Post("/users", handlers.CreateUser)
	api.Get("/users/:id", handlers.GetUserByID)
}