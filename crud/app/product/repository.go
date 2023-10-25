package product

import (
	"context"

	"gorm.io/gorm"
)

type PostgresGormRepository struct {
	db *gorm.DB
}

func NewPosrgresRepository(db *gorm.DB) PostgresGormRepository {
	return PostgresGormRepository{
		db: db,
	}
}

func (p PostgresGormRepository) Create(ctx context.Context, model Product) (err error) {
	return p.db.Create(&model).Error
}

func (p PostgresGormRepository) FindAll(ctx context.Context, models []Product) ([]Product, error) {
	err := p.db.Find(&models).Error
	if err != nil {
		return models, err
	}

	return models, nil
}

func (p PostgresGormRepository) FindByID(ctx context.Context, model Product, id int) (Product, error) {
	findById := p.db.Where("id = ?", id).First(&gorm.ErrModelValueRequired)
	return model, findById.Error
}
