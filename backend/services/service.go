package services

import "github.com/vasuvanka/library_management/backend/services/database"

type Service struct {
	database.Database
}

func NewService() *Service {
	return &Service{}
}