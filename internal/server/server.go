package server

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r
}
