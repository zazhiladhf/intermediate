package main

import (
	"crud/app/product"
	"crud/config"
	"crud/pkg/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New(fiber.Config{
		AppName: "Product Services",
		Prefork: true,
	})

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	db, err := database.ConnectGORMPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	dbSqlx, err := database.ConnectSqlxPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	product.RegisterServiceProduct(router, db, dbSqlx, nil)

	router.Listen(config.Cfg.App.Port)
}
