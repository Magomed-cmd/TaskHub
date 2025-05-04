package routers

import (
	"TaskHub/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, s *service.TaskService) {

	v1 := r.Group("api/v1")

	RegisterAuthRoutes(v1, s)
	RegisterTaskRoutes(v1, s)
}
