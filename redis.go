package main

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/redis.v5"
)

func GetClient() *redis.Client {
	log.Debug("Getting a redis client connection")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
