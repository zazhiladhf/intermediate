package product

import (
	"context"
	"log"
	"product-catalog/domain/auth"
)

type Repository interface {
	InsertProduct(ctx context.Context, model Product) (id int, err error)
	FindAll(ctx context.Context) (list []Product, err error)
	FindByEmail(ctx context.Context, queryParam string, email string) (list []Product, err error)
	// Update(ctx context.Context, id int, model Product) error
	// Delete(ctx context.Context, model Product, id int) error
}

type SearchEngineInterface interface {
	SyncPartial(ctx context.Context, productList []Product) (statusId int, err error)
	// SyncAll(ctx context.Context, productList []Product) (statusId int, err error)
	// GetAll(ctx context.Context) (productList []Product, err error)
}

type Service struct {
	repo        Repository
	search      SearchEngineInterface
	authService auth.AuthService
}

func NewService(repo Repository, search SearchEngineInterface, authService auth.AuthService) Service {
	return Service{
		repo:        repo,
		search:      search,
		authService: authService,
	}
}

func (s Service) createProduct(ctx context.Context, req CreateProductRequest, token string) (err error) {
	product, err := newFromRequest(req)
	if err != nil {
		return
	}
	// product.AuthEmail = token

	// if err = req.(); err != nil {
	// 	log.Println("erro when try to validate request with error")
	// 	return
	// }

	auth, err := s.authService.GetAuthByEmail(ctx, token)
	if err != nil {
		log.Println("error when try to get auth by email with error :", err.Error(), token)
		return err
	}

	model := Product{
		Name:       product.Name,
		Stock:      product.Stock,
		Price:      product.Price,
		CategoryId: product.CategoryId,
		ImageURL:   product.ImageURL,
		AuthEmail:  auth.Email,
	}

	id, err := s.repo.InsertProduct(ctx, model)
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

func (s Service) GetProducts(ctx context.Context) (list []Product, err error) {
	listProducts, err := s.repo.FindAll(ctx)
	if err != nil {
		if err == ErrCategoriesNotFound {
			return []Product{}, nil
		}
		return nil, err
	}

	return listProducts, nil

}

func (s Service) GetProductsByEmail(ctx context.Context, queryParam string, email string) (list []Product, err error) {
	listProducts, err := s.repo.FindByEmail(ctx, queryParam, email)
	if err != nil {
		if err == ErrCategoriesNotFound {
			return []Product{}, nil
		}
		return nil, err
	}

	return listProducts, nil
}

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
