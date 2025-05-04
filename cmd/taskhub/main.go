package main

import (
	"TaskHub/internal/config"
	"TaskHub/internal/db"
	"TaskHub/internal/repository/postgres"
	"TaskHub/internal/routers"
	"TaskHub/internal/server"
	"TaskHub/internal/service"
	"log"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Config error: ", err)
	}
	conn := db.Connect(&cfg)

	taskRepo := postgres.NewTaskRepo(conn)
	taskService := service.NewTaskService(taskRepo)

	router := server.New()
	routers.RegisterRoutes(router, taskService)

	if err := router.Run(":" + cfg.PORT); err != nil {
		log.Fatal(err)
	}

}
