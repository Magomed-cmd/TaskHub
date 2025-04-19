package main

import (
	"TaskHub/internal/config"
	"TaskHub/internal/server"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	router := server.New()


	router.GET("/ping", handler)

	router.Run(":" + cfg.PORT, )

}


func handler(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
}