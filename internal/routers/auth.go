package routers

import (
	"TaskHub/internal/handler"
	"TaskHub/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, s *service.AuthService) {
	auth := rg.Group("/auth")

	h := handler.NewAuthHandler(*s)

	auth.POST("/", h.Login)
}
