package config

import "os"

type mysqlConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	Dbname string
}

func initMysqlConfig() mysqlConfig {
	var mysql = mysqlConfig{}
	mysql.Host = os.Getenv("MYSQL_HOST")
	mysql.Port = os.Getenv("MYSQL_PORT")
	mysql.User = os.Getenv("MYSQL_USER")
	mysql.Pass = os.Getenv("MYSQL_PASS")
	mysql.Dbname = os.Getenv("MYSQL_DB")
	return mysql
}
