package product

type CreateProductRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}

type GetProductByIdUri struct {
	Id int `uri:"id" binding:"required"`
}
