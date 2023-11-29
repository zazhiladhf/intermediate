package main

import (
	"log"
	"product-catalog/app/product"
	"product-catalog/config"
	"product-catalog/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	router := fiber.New(fiber.Config{
		AppName: config.Cfg.App.Name,
		Prefork: true,
	})

	// db, err := database.ConnectGORMPostgres(config.Cfg.DB)
	// if err != nil {
	// 	panic(err)
	// }

	dbSqlx, err := database.ConnectSqlxPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	product.RegisterServiceProduct(router, dbSqlx)

	router.Listen(config.Cfg.App.Port)
}
