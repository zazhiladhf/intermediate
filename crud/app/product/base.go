package product

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func RegisterServiceProduct(router fiber.Router, db *gorm.DB, dbSqlx *sqlx.DB, dbNative *sql.DB) {
	repo := NewPosrgresRepository(db)
	// repo := NewPostgresSQLXRepository(dbSqlx)

	svc := NewService(repo)
	handler := NewHandler(svc)

	router.Post("/products", handler.CreateProduct)
	router.Get("/products", handler.GetProducts)
	router.Get("/product/:id", handler.GetProductById)
	router.Put("/product/:id", handler.UpdateProduct)
	router.Delete("/product/:id", handler.DeleteProduct)

}
