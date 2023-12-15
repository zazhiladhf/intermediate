package category

import (
	"context"
)

type repository interface {
	FindAll(ctx context.Context) (item []Category, err error)
}

type categoryService struct {
	repo repository
}

func newService(repo repository) categoryService {
	return categoryService{
		repo: repo,
	}
}

func (s categoryService) getListCategories(ctx context.Context) (list []Category, err error) {
	listCategories, err := s.repo.FindAll(ctx)
	if err != nil {
		if err == ErrCategoriesNotFound {
			return []Category{}, nil
		}
		return nil, err
	}
	return listCategories, nil
}
