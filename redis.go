package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/redis.v5"
)

func GetClient() *redis.Client {
	log.Debug("Getting a redis client connection")

	redis_host := os.Getenv("WCGSTATS_API_REDIS_HOST")
	if redis_host == "" {
		redis_host = "localhost"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redis_host + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
