package main

import (
	"log"
	"product-catalog/config"
	"product-catalog/domain/auth"
	"product-catalog/domain/category"
	"product-catalog/domain/files"
	"product-catalog/domain/product"
	"product-catalog/pkg/database"
	"product-catalog/pkg/images"
	"product-catalog/pkg/middleware"
	"product-catalog/pkg/search"

	"github.com/gofiber/fiber/v2"
)

// type CloudSvc interface {
// 	Upload(ctx context.Context, file interface{}, pathDestination string, quality string) (uri string, err error)
// }

// // setup services
// type service struct {
// 	// disini kita akan menggunakan kontrak ke cloud providernya
// 	cloud CloudSvc
// }

// var path = "public/uploads"
// var svc = service{}

// func init() {
// 	err := config.LoadConfig("./config/config.yaml")
// 	if err != nil {
// 		log.Println("error when try to LoadConfig with error :", err.Error())
// 	}

// 	cloudName := config.Cfg.Cloudinary.Name
// 	apiKey := config.Cfg.Cloudinary.ApiKey
// 	apiSecret := config.Cfg.Cloudinary.ApiSecret

// 	cloudClient, err := images.NewCloudinary(cloudName, apiKey, apiSecret)
// 	if err != nil {
// 		panic(err)
// 	}

// 	svc = service{
// 		cloud: cloudClient,
// 	}
// }

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

	cloudName := config.Cfg.Cloudinary.Name
	apiKey := config.Cfg.Cloudinary.ApiKey
	apiSecret := config.Cfg.Cloudinary.ApiSecret

	cloudClient, err := images.NewCloudinary(cloudName, apiKey, apiSecret)
	if err != nil {
		panic(err)
	}
	files.RegisterRouteFiles(router, cloudClient, cloudName, apiKey, apiSecret)

	router.Listen(config.Cfg.App.Port)
}
