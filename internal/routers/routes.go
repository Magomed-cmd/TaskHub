package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r *gin.Engine, conn *sqlx.DB){
	v1 := r.Group("api/v1")
	RegisterAuthRoutes(v1, conn)
	RegisterTaskRoutes(v1, conn)
}