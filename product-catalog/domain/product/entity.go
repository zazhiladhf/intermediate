package product

import "errors"

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
	// AuthId     int
}

func newFromRequest(req CreateProductRequest) (product Product, err error) {
	product = Product{
		Name:       req.Name,
		ImageURL:   req.ImageURL,
		Stock:      req.Stock,
		Price:      req.Price,
		CategoryId: req.CategoryId,
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
