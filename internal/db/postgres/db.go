package postgres

import (
	"TaskHub/internal/config"
	model2 "TaskHub/internal/pkg/model"
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

	conn, err := gorm.Open(postgres.Open(cfg.GetPostgresDSN()), gormConfig)

	if err != nil {
		log.Fatalln("DB connection error: ", err)
	}

	err = conn.AutoMigrate(&model2.User{}, &model2.Task{})
	if err != nil {
		log.Fatalf("failed to migrate DB: %v", err)
	}

	return conn
}
