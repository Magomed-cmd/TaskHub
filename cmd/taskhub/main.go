package main

import (
	"TaskHub/internal/config"
	"TaskHub/internal/db/postgres"
	"TaskHub/internal/db/redis"
	"TaskHub/internal/repository/postgresql"
	"TaskHub/internal/routers"
	"TaskHub/internal/server"
	"TaskHub/internal/service"
	"log"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln("Config error: ", err)
	}

	pgConn := postgres.Connect(cfg)

	// TODO: не забыть вернуть значение redis на место
	_, err = redis.Connect(cfg)
	if err != nil {
		log.Fatalln("Redis connection error: ", err)
	}

	userRepo := postgresql.NewUserRepo(pgConn)

	taskService := service.NewTaskService(postgresql.NewTaskRepo(pgConn))
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
