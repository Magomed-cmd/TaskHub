package db

import (
	"TaskHub/internal/config"
	"TaskHub/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func Connect(cfg *config.Config) *gorm.DB {

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	conn, err := gorm.Open(postgres.Open(cfg.GetDSN()), gormConfig)

	if err != nil {
		log.Fatalln("DB connection error: ", err)
	}

	err = conn.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Fatalf("failed to migrate DB: %v", err)
	}

	return conn
}
