package config

import "os"

type logConfig struct {
	logFilePath string
}

func initLogConfig() logConfig {
	var log = logConfig{}
	log.logFilePath = os.Getenv("LOG_FILE")
	return log
}
