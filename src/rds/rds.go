package rds

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

var mineConf = redis.Options{
	Addr:     "",
	Password: "",
	DB:       0,
}

func InitRedisClient() error {
	client := redis.NewClient(&mineConf)
	_, err := client.Ping().Result()
	if err != nil {
		return fmt.Errorf("ping to redis error: %v", err)
	} else {
		Client = client
		return nil
	}
}
