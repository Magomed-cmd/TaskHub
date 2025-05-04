package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
	PORT    string
}

func (cfg *Config) GetDSN() string {

	return "user=" + cfg.DB_USER + " password=" + cfg.DB_PASS + " dbname=" + cfg.DB_NAME + " sslmode=disable"

}

func LoadConfig() (Config, error) {

	err := godotenv.Load()
	log.Println("Loading .env file")

	if err != nil {
		log.Println("Error loading .env file", err)
		return Config{}, err
	}

	log.Println("Config loaded")

	return Config{
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_NAME: os.Getenv("DB_NAME"),
		PORT:    os.Getenv("PORT"),
	}, nil
}
