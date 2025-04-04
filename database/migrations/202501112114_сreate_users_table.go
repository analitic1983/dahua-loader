package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up_20250111, Down_20250111)
}

func Up_20250111(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
CREATE TABLE users (
  uuid char(36) NOT NULL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  pass_hash varchar(520) DEFAULT NULL,
  created_at datetime NOT NULL,
  status enum('active', 'inactive') NOT NULL,
  PRIMARY KEY (uuid),
  UNIQUE KEY email (email)
) ENGINE=InnoDB;
	`)
	return err
}

func Down_20250111(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS users`)
	return err
}
