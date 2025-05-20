package redis

import (
	"TaskHub/internal/config"
	"github.com/redis/go-redis/v9"
	"log"
)

func Connect(cfg *config.Config) (*redis.Client, error) {

	opts, err := redis.ParseURL(cfg.GetRedisURL())
	if err != nil {
		log.Println("error to get connect with redis")
		return nil, err
	}
	return redis.NewClient(opts), nil
}
