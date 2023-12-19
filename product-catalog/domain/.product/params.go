package product

import "product-catalog/domain/auth"

type CreateProductRequest struct {
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	CategoryId int    `json:"category_id"`
	Auth       auth.Auth
}
