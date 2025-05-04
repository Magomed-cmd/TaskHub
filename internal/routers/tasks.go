package routers

import (
	"TaskHub/internal/handler"
	"TaskHub/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(rg *gin.RouterGroup, s *service.TaskService) {
	task := rg.Group("/task")

	h := handler.NewTaskHandler(s)

	task.GET("/", h.GetTasks)
	task.POST("/", h.Create)
	task.DELETE("/:id/", h.Delete)
	task.PATCH("/:id", h.UpdateTask)
}
