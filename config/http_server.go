package config

import "os"

type httpServerConfig struct {
	Host string
	Port string
}

func initHttpServerConfig() httpServerConfig {
	var config = httpServerConfig{}
	config.Host = os.Getenv("WEB_HOST")
	config.Port = os.Getenv("WEB_PORT")
	return config
}
