package config

import (
	"os"
	"strconv"
)

type redisConfig struct {
	Host string
	Port string
	Db   int
}

func initRedisConfig() redisConfig {
	var redis = redisConfig{}
	redis.Host = os.Getenv("REDIS_HOST")
	redis.Port = os.Getenv("REDIS_PORT")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redis.Db = db
	return redis
}
