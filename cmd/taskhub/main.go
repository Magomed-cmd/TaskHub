package main

import (
	"TaskHub/internal/config"
	"TaskHub/internal/db"
	"TaskHub/internal/repository/postgres"
	"TaskHub/internal/routers"
	"TaskHub/internal/server"
	"TaskHub/internal/service"
	"fmt"
	"log"
)

func main() {
	fmt.Println("=== Start main.go ===")
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln("Config error: ", err)
	}
	conn := db.Connect(cfg)

	userRepo := postgres.NewUserRepo(conn)
	taskService := service.NewTaskService(postgres.NewTaskRepo(conn))
	authService := service.NewAuthService(userRepo, cfg.App.JWTSecret)
	userService := service.NewUserService(userRepo)
	services := service.Services{
		TaskService: taskService,
		UserService: userService,
		AuthService: authService,
	}

	router := server.New()
	routers.RegisterRoutes(router, &services)
	if err := router.Run(":" + cfg.App.PORT); err != nil {
		log.Fatal(err)
	}

}
