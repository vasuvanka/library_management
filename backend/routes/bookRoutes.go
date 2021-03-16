package routes

import "github.com/vasuvanka/library_management/backend/handlers"

func BookRoutes(server *Server){
	api := server.App.Group("/api")
	api.Get("books", handlers.GetBooks)
	api.Post("books", handlers.CreateBook)
	api.Post("books/:id/borrow", handlers.BorrowBook)
	api.Put("books/:id/return", handlers.ReturnBook)
}