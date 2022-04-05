package rds

import (
	"github.com/go-redis/redis"
	"log"
	"slink/src/conf"
)

var Client *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RdsAddr,
		Password: conf.RdsPswd,
		DB:       conf.RdsDb,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("ping to redis error: %v", err)
	} else {
		Client = client
	}
}
