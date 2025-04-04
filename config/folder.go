package config

import "os"

type folderConfig struct {
	Download  string
	Converted string
}

func initFolderConfig() folderConfig {
	var config = folderConfig{}
	config.Download = os.Getenv("FOLDER_DOWNLOAD")
	config.Converted = os.Getenv("FOLDER_CONVERTED")
	return config
}
