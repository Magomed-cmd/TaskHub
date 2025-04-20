package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

)


func RegisterTaskRoutes(rg *gin.RouterGroup, db *sqlx.DB){
	tasks := rg.Group("/tasks")
	tasks.GET("/", )
}