package database

import (
	"cake-store/src/config"
	"context"

	// "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func NewRedisConn(url string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// pool := goredis.NewPool(rdb)

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.WithField("Addr", config.RedisHost()).Fatal("Failed to connect:", err)
	}

	log.Info("Success connect redis")
	return rdb
}
