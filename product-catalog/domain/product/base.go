package product

import (
	"product-catalog/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
)

func RegisterServiceProduct(router fiber.Router, dbSqlx *sqlx.DB, client *meilisearch.Client) {
	meiliRepo := NewMeiliRepository(client)
	repo := NewPostgresSQLXRepository(dbSqlx)
	svc := NewService(repo, meiliRepo)
	handler := NewHandler(svc)

	v1 := router.Group("v1")
	v1.Post("/products", middleware.JWTProtected(), handler.CreateProduct)

	// router.Post("/v1/products", handler.CreateProduct)
	// router.Get("/products", handler.GetProducts)
	// router.Get("/product/:id", handler.GetProductById)
	// router.Put("/product/:id", handler.UpdateProduct)
	// router.Delete("/product/:id", handler.DeleteProduct)

}
