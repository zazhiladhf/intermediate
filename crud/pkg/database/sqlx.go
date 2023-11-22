package database

import (
	"crud/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectSqlxPostgres(cfg config.DB) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	return
}
