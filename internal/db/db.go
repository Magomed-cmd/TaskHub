package db

import (
	"TaskHub/internal/config"
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

)



func Connect(cfg *config.Config) *sqlx.DB{

	conn, err := sqlx.Connect("postgres", cfg.GetDSN())
	if err != nil{
		log.Fatalln("DB connection error: ", err)
	}
	return conn
}