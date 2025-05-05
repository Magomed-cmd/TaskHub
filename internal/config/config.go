package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DataBase struct {
		DBHost  string `mapstructure:"DB_HOST"`
		DBPort  string `mapstructure:"DB_PORT"`
		DBUser  string `mapstructure:"DB_USER"`
		DBPass  string `mapstructure:"DB_PASS"`
		DBName  string `mapstructure:"DB_NAME"`
		SSLMode string `mapstructure:"SSLMode"`
	} `mapstructure:"DataBase"`
	App struct {
		PORT           string   `mapstucture:"PORT"`
		TrustedProxies []string `mapstructure:"TrustedProxies"`
	} `mapstructure:"App"`
}

func (cfg *Config) GetDSN() string {

	return "host=" + cfg.DataBase.DBHost +
		" port=" + cfg.DataBase.DBPort +
		" user=" + cfg.DataBase.DBUser +
		" password=" + cfg.DataBase.DBPass +
		" dbname=" + cfg.DataBase.DBName +
		" sslmode=" + cfg.DataBase.SSLMode

}

func InitConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/config")

	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		log.Println("error to read config file: ", err)
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Println("error to Unmarshalling config structure: ", err)
		return nil, err
	}

	return &cfg, nil
}
