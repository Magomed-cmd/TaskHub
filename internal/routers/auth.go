package routers

import (
	"TaskHub/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, s *service.TaskService) {
	auth := rg.Group("/auth")
	auth.GET("/")
}
