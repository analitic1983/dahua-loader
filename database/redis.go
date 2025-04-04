package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"koshmin/dahua-loader/config"
	"log"
)

var RedisDB *redis.Client

func InitRedis() error {
	db := redis.NewClient(&redis.Options{
		Addr: config.AppConfig.Redis.Host + ":" + config.AppConfig.Redis.Port,
		DB:   config.AppConfig.Redis.Db,
	})

	ctx := context.Background()
	if err := db.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to redis server: ", err)
		return err
	}

	RedisDB = db

	return nil
}
