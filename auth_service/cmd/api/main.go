package main

import (
	"auth_service/pkg/config"
	"auth_service/pkg/di"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to load the Congiguration File: ", err)
		return
	}
	server, err := di.InitApi(cfg)
	if err != nil {
		log.Println("Error in initializing the api", err)
	}
	server.Start()
}
