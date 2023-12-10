package database

import (
	"fmt"
	"log"
	"product-catalog/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgresSqlx(cfg config.DB) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	if db == nil {
		log.Println("error when to try connect db with error :", err.Error())
		panic("db not connected")
	}

	log.Println("database connect success 🚀🚀🚀")
	log.Println("dsn :", dsn)

	return
}

func Migrate(db *sqlx.DB) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS auths (
			id SERIAL PRIMARY KEY,
			email varchar(100) NOT NULL,
			password varchar(100) NOT NULL,
			role varchar(100) NOT NULL,
			UNIQUE (email)
		);

		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name varchar(100) NOT NULL
		);
	`
	_, err = db.Exec(query)

	return
}
