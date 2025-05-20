package routers

import (
	"TaskHub/internal/handler"
	"TaskHub/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, s *service.Services) {
	v1 := r.Group("/api/v1")

	authHandler := handler.NewAuthHandler(s.AuthService)
	userHandler := handler.NewUserHandler(s.UserService)
	taskHandler := handler.NewTaskHandler(s.TaskService)

	// Auth routes
	v1.POST("/auth", authHandler.Login)

	// User routes
	v1.POST("/users/register", userHandler.CreateUser)

	// Task routes
	tasks := v1.Group("/tasks")
	{
		tasks.GET("/", taskHandler.GetTasks)
		tasks.POST("/", taskHandler.Create)
		tasks.GET("/:id", taskHandler.GetTaskByID)
		tasks.PATCH("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.Delete)
	}
}
