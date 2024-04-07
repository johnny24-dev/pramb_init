package main

import (
	"api_gateway/pkg/auth"
	"api_gateway/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load the config file : ", err)
	}
	r := gin.Default()

	auth.RegisterRoutes(r, &cfg)
	// Start the server
	r.Run(cfg.Port)
}
