package config

import (
	"fmt"
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
		JWTSecret      string   `mapstructure:"JWTSecret"`
	} `mapstructure:"App"`
	Client struct {
		User     string `mapstructure:"User"`
		Pass     string `mapstructure:"Pass"`
		Host     string `mapstructure:"Host"`
		Port     string `mapstructure:"Port"`
		Protocol string `mapstructure:"Protocol"`
		DB       string `mapstructure:"DB"`
	}
}

func (cfg *Config) GetPostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DataBase.DBHost,
		cfg.DataBase.DBPort,
		cfg.DataBase.DBUser,
		cfg.DataBase.DBPass,
		cfg.DataBase.DBName,
		cfg.DataBase.SSLMode,
	)
}

func (cfg *Config) GetRedisURL() string {
	return fmt.Sprintf(
		"redis://%s:%s@%s:%s/%s?protocol=%s",
		cfg.Client.User,
		cfg.Client.Pass,
		cfg.Client.Host,
		cfg.Client.Port,
		cfg.Client.DB,
		cfg.Client.Protocol,
	)
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
