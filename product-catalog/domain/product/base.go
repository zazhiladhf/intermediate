package product

import (
	"product-catalog/domain/auth"
	"product-catalog/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
)

func RegisterServiceProduct(router fiber.Router, dbSqlx *sqlx.DB, client *meilisearch.Client) {
	meiliRepo := NewMeiliRepository(client)
	authRepo := auth.NewPostgreSqlxRepository(dbSqlx)
	authService := auth.NewService(authRepo)
	repo := NewPostgresSQLXRepository(dbSqlx)
	svc := NewService(repo, meiliRepo, authService)
	handler := NewHandler(svc)

	v1 := router.Group("v1")
	v1.Post("/products", middleware.AuthMiddleware(), handler.CreateProduct)
	// v1.Get("/products", middleware.AuthMiddleware(), handler.GetProducts)
	v1.Get("/products", middleware.AuthMiddleware(), handler.GetProductsByEmail)

	// router.Post("/v1/products", handler.CreateProduct)
	// router.Get("/products", handler.GetProducts)
	// router.Get("/product/:id", handler.GetProductById)
	// router.Put("/product/:id", handler.UpdateProduct)
	// router.Delete("/product/:id", handler.DeleteProduct)

}
