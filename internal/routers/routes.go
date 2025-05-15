package routers

import (
	"TaskHub/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, s *service.Services) {

	v1 := r.Group("api/v1")

	RegisterAuthRoutes(v1, s.AuthService)
	RegisterTaskRoutes(v1, s.TaskService)
	RegisterUserRoutes(v1, s.UserService)
}
