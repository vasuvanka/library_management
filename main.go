package main

import (
	"log"

	"github.com/vasuvanka/library_management/backend/config"
)

func main(){
	conf := config.NewConfig()

	err := conf.Init()
	if err != nil {
		log.Fatal(err)
	}
	
}