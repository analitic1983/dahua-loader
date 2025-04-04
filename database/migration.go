package database

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"koshmin/dahua-loader/config"
	_ "koshmin/dahua-loader/database/migrations"
	"log"
)

func migrateCommand(command string) {
	dsn := config.AppConfig.Mysql.User + ":" + config.AppConfig.Mysql.Pass + "@tcp(" + config.AppConfig.Mysql.Host + ":" + config.AppConfig.Mysql.Port + ")/" + config.AppConfig.Mysql.Dbname + "?parseTime=true"
	migrationsDir := "./database/migrations"

	db, err := goose.OpenDBWithDriver("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run goose command
	ctx := context.Background()
	if err := goose.RunContext(ctx, command, db, migrationsDir); err != nil {
		log.Fatalf("Goose command failed: %v", err)
	}
}

func MigrationsUp() {
	migrateCommand("up")
}

func MigrationsDown() {
	migrateCommand("down")
}

func MigrationsStatus() {
	migrateCommand("status")
}
