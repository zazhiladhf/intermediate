package product

type CreateProductRequest struct {
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	CategoryId string `json:"category_id"`
}
