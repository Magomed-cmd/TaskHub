package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

)


func RegisterAuthRoutes(rg *gin.RouterGroup, db *sqlx.DB){
	auth := rg.Group("/auth")
	auth.GET("/", )
}