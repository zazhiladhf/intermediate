package product

import "errors"

var (
	ErrEmptyName     = errors.New("name tidak boleh kosong")
	ErrEmptyImageURL = errors.New("image_url tidak boleh kosong")
	ErrEmptyPrice    = errors.New("price tidak boleh kosong")
	ErrEmptyStock    = errors.New("stock tidak boleh kosong")
	ErrNotFound      = errors.New("data tidak ditemukan")
)

type Product struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	CategoryId string `json:"category_id"`
}

func (p Product) Validate() (err error) {
	if p.Name == "" {
		return ErrEmptyName
	}

	if p.ImageURL == "" {
		return ErrEmptyImageURL
	}

	if p.Price == 0 {
		return ErrEmptyPrice
	}

	if p.Stock == 0 {
		return ErrEmptyStock
	}

	if err != nil {
		return ErrNotFound
	}

	return
}
