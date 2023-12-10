package category

import (
	"context"
)

type repository interface {
	FindAll(ctx context.Context) (item []category, err error)
}

type categoryService struct {
	repo repository
}

func newService(repo repository) categoryService {
	return categoryService{
		repo: repo,
	}
}

func (s categoryService) getListCategories(ctx context.Context) (list []category, err error) {
	listCategories, err := s.repo.FindAll(ctx)
	if err != nil {
		if err == ErrCategoriesNotFound {
			return []category{}, nil
		}
		return nil, err
	}
	return listCategories, nil
}
