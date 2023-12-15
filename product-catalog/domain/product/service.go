package product

import (
	"context"
	"log"
)

type Repository interface {
	InsertProduct(ctx context.Context, model Product) (id int, err error)
	// FindAll(ctx context.Context, models []Product) ([]Product, error)
	// FindByID(ctx context.Context, model Product, id int) (Product, error)
	// Update(ctx context.Context, id int, model Product) error
	// Delete(ctx context.Context, model Product, id int) error
}

type SearchEngineInterface interface {
	SyncPartial(ctx context.Context, productList []Product) (statusId int, err error)
	// SyncAll(ctx context.Context, productList []Product) (statusId int, err error)
	// GetAll(ctx context.Context) (productList []Product, err error)
}

type Service struct {
	repo   Repository
	search SearchEngineInterface
}

func NewService(repo Repository, search SearchEngineInterface) Service {
	return Service{
		repo:   repo,
		search: search,
	}
}

func (s Service) createProduct(ctx context.Context, req CreateProductRequest) (err error) {
	product, err := newFromRequest(req)
	if err != nil {
		return
	}

	// if err = req.(); err != nil {
	// 	log.Println("erro when try to validate request with error")
	// 	return
	// }

	id, err := s.repo.InsertProduct(ctx, product)
	if err != nil {
		log.Println("error when try to Create to database with error :", err.Error(), product)
		return err
	}

	product.Id = id

	status, err := s.search.SyncPartial(ctx, []Product{product})
	if err != nil {
		return err
	}

	log.Println("status id", status)

	return
}

// func (s Service) GetProducts(ctx context.Context, models []Product) ([]Product, error) {
// 	products, err := s.repo.FindAll(ctx, models)
// 	if err != nil {
// 		return products, err
// 	}

// 	return products, nil

// }

// func (s Service) GetProductById(ctx context.Context, model Product, param int) (Product, error) {
// 	product, err := s.repo.FindByID(ctx, model, param)
// 	if err != nil {
// 		return product, err
// 	}

// 	return product, nil
// }

// func (s Service) UpdateProduct(ctx context.Context, req Product, param int) (err error) {
// 	if err = req.Validate(); err != nil {
// 		log.Println("erro when try to validate request with error")
// 		return
// 	}

// 	if err = s.repo.Update(ctx, param, req); err != nil {
// 		log.Println("error when try to Update to database with error :", err.Error())
// 		return
// 	}
// 	return

// }

// func (s Service) DeleteProduct(ctx context.Context, model Product, param int) (err error) {
// 	if err = s.repo.DeleteProduct(ctx, model, param); err != nil {
// 		log.Println("error when try to Delete to database with error :", err.Error())
// 		return
// 	}
// 	return

// }
