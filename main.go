package main

import (
	"log"

	"github.com/vasuvanka/library_management/backend/config"
	"github.com/vasuvanka/library_management/backend/routes"
)

func main(){
	conf := config.NewConfig()

	err := conf.Init()
	if err != nil {
		log.Fatal(err)
	}

	server := routes.NewServer(conf)
	if err := server.Init(); err != nil {
		log.Fatal(err)
	}

	server.Controllers()
	server.Routes()
	
	if err := server.Bootstrap(); err != nil {
		log.Fatal(err)
	}
}