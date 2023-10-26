package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterServiceProduct(router fiber.Router, db *gorm.DB) {
	repo := NewPosrgresRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	router.Post("/products", handler.CreateProduct)
	router.Get("/products", handler.GetProducts)
	router.Get("/product/:id", handler.GetProductById)
	router.Put("/product/:id", handler.UpdateProduct)
	router.Delete("/product/:id", handler.DeleteProduct)

}
