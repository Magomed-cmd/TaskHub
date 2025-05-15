package routers

import (
	"TaskHub/internal/handler"
	"TaskHub/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, s *service.UserService) {
	user := r.Group("user")

	h := handler.NewUserHandler(s)

	user.POST("/register", h.CreateUser)

}
