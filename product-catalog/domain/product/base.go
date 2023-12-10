package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterServiceProduct(router fiber.Router, dbSqlx *sqlx.DB) {
	// repo := NewPosrgresRepository(db)
	repo := NewPostgresSQLXRepository(dbSqlx)

	svc := NewService(repo)
	handler := NewHandler(svc)

	// router.Post("/v1/products", handler.CreateProduct)
	router.Get("/products", handler.GetProducts)
	// router.Get("/product/:id", handler.GetProductById)
	// router.Put("/product/:id", handler.UpdateProduct)
	// router.Delete("/product/:id", handler.DeleteProduct)

}
