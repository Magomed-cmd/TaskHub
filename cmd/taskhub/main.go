package main

import (
	"TaskHub/internal/config"
	"log"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil{
		log.Fatalln(err)	
	}

	log.Println(&cfg)

}
