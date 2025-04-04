package config

import (
	"github.com/joho/godotenv"
	"log"
)

type appConfig struct {
	Folder     folderConfig
	Log        logConfig
	Mysql      mysqlConfig
	Redis      redisConfig
	HttpServer httpServerConfig
}

var AppConfig *appConfig = nil

func Init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := appConfig{}
	config.Folder = initFolderConfig()
	config.Log = initLogConfig()
	config.Mysql = initMysqlConfig()
	config.Redis = initRedisConfig()
	config.HttpServer = initHttpServerConfig()

	AppConfig = &config
}
