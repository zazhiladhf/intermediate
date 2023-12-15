package main

import (
	"log"
	"product-catalog/config"
	"product-catalog/domain/auth"
	"product-catalog/domain/category"
	"product-catalog/domain/product"
	"product-catalog/pkg/database"
	"product-catalog/pkg/middleware"
	"product-catalog/pkg/search"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	router := fiber.New(fiber.Config{
		AppName: config.Cfg.App.Name,
		// Prefork: true,
	})

	middleware.FiberMiddleware(router)

	dbSqlx, err := database.ConnectPostgresSqlx(config.Cfg.DB)
	if err != nil {
		log.Println("error when to try migration db with error :", err.Error())
		// panic(err)
	}

	client, err := search.ConnectMeilisearch(config.Cfg.Meili.Host, config.Cfg.Meili.ApiKey)
	if err != nil {
		log.Println("error connect meili", err)
	}

	log.Println("running db migration")
	err = database.Migrate(dbSqlx)
	if err != nil {
		log.Println("migration failed with error:", err)
		panic(err)
	}
	log.Println("migration done")

	auth.Run(router, dbSqlx)
	category.RegisterCategoriesRouter(router, dbSqlx)
	product.RegisterServiceProduct(router, dbSqlx, client)

	// redis(5 * time.Second)

	router.Listen(config.Cfg.App.Port)
}
