package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"koshmin/dahua-loader/config"
	"log"
)

var DB *sql.DB
var GormDB *gorm.DB

// InitSQL db connection
func InitSQL() error {
	var err error
	dsn := config.AppConfig.Mysql.User + ":" + config.AppConfig.Mysql.Pass + "@tcp(" + config.AppConfig.Mysql.Host + ":" + config.AppConfig.Mysql.Port + ")/" + config.AppConfig.Mysql.Dbname + "?parseTime=true"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed database connect: ", err)
		return err
	}

	// Check connection
	if err := DB.Ping(); err != nil {
		log.Fatal("Db ping failed: ", err)
		return err
	}

	// InitSQL gorm
	GormDB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: DB,
	}), &gorm.Config{})

	return nil
}
