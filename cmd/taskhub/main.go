package main

import (
	"TaskHub/internal/server"
	"log"
	"TaskHub/internal/db"
	"TaskHub/internal/config"
	"TaskHub/internal/routers"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil{
		log.Fatalln("Config error: ", err)
	}

	conn := db.Connect(&cfg)
	router := server.New()
	routers.RegisterRoutes(router, conn)

	router.Run(":" + cfg.PORT)
}
