package product

import (
	"context"
	"log"
)

type Repository interface {
	Create(ctx context.Context, req Product) (err error)
	FindAll(ctx context.Context, models []Product) ([]Product, error)
	FindByID(ctx context.Context, model Product, id int) (Product, error)
	Update(ctx context.Context, model Product) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateProduct(ctx context.Context, req Product) (err error) {
	if err = req.Validate(); err != nil {
		log.Println("erro when try to validate request with error")
		return
	}

	if err = s.repo.Create(ctx, req); err != nil {
		log.Println("error when try to Create to database with error :", err.Error())
		return
	}
	return
}

func (s Service) GetProducts(ctx context.Context, models []Product) ([]Product, error) {

	products, err := s.repo.FindAll(ctx, models)
	if err != nil {
		return products, err
	}

	return products, nil

}

func (s Service) GetProductById(ctx context.Context, model Product, param int) (Product, error) {
	product, err := s.repo.FindByID(ctx, model, param)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s Service) UpdateProduct(ctx context.Context, req Product, param int) (err error) {
	if err = req.Validate(); err != nil {
		log.Println("erro when try to validate request with error")
		return
	}

	// product, err := s.repo.FindByID(ctx, req, param)
	// if err != nil {
	// 	return product, err
	// }

	// if product.UserID != inputData.User.ID {
	// 	return product, errors.New("not an owner of the campaign")
	// }

	// campaign.Name = inputData.Name
	// campaign.ShortDescription = inputData.ShortDescription
	// campaign.Description = inputData.Description
	// campaign.Perks = inputData.Perks
	// campaign.GoalAmount = inputData.GoalAmount

	if err = s.repo.Update(ctx, req); err != nil {
		log.Println("error when try to Update to database with error :", err.Error())
		return
	}
	return
	// return s.repo.Update(ctx, req)

}
