package services_test

import (
	"testing"

	"github.com/vasuvanka/library_management/backend/services"
)

func TestNewService(t *testing.T)  {
	service := services.NewService()

	if service == nil {
		t.Error("Expected service object but got nil")
	}
}