package product

import (
	"errors"
)

var (
	ErrEmptyName       = errors.New("name is required")
	ErrEmptyImageURL   = errors.New("image_url is required")
	ErrEmptyStock      = errors.New("stock is required")
	ErrEmptyPrice      = errors.New("price is required")
	ErrEmptyCategoryId = errors.New("category_id is required")
	ErrNotFound        = errors.New("product not found")
)

type Product struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Stock      int    `json:"stock" db:"stock"`
	Price      int    `json:"price" db:"price"`
	CategoryId int    `json:"category_id" db:"category_id"`
	ImageURL   string `json:"image_url" db:"image_url"`
	AuthEmail  string `db:"email_auth"`
}

func newFromRequest(req CreateProductRequest) (product Product, err error) {
	product = Product{
		Name:       req.Name,
		ImageURL:   req.ImageURL,
		Stock:      req.Stock,
		Price:      req.Price,
		CategoryId: req.CategoryId,
		AuthEmail:  req.Auth.Email,
	}

	err = product.validateRequestProduct()
	return
}

func (p Product) validateRequestProduct() (err error) {
	if p.Name == "" {
		return ErrEmptyName
	}

	if p.ImageURL == "" {
		return ErrEmptyImageURL
	}

	if p.Stock == 0 {
		return ErrEmptyStock
	}

	if p.Price == 0 {
		return ErrEmptyPrice
	}

	if p.CategoryId == 0 {
		return ErrEmptyCategoryId
	}

	if err != nil {
		return ErrNotFound
	}

	return
}

type ProductAuth struct {
	AuthId     int    `db:"auth_id"`
	ProductId  int    `db:"product_id"`
	Name       string `db:"name"`
	Stock      int    `db:"stock"`
	Price      int    `db:"price"`
	CategoryId int    `db:"category_id"`
	ImageURL   string `db:"image_url"`
	Email      string `db:"email"`
	Role       string `db:"role"`
}
