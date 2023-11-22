package product

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type PostgresSQLXRepository struct {
	db *sqlx.DB
}

func NewPostgresSQLXRepository(db *sqlx.DB) PostgresSQLXRepository {
	return PostgresSQLXRepository{
		db: db,
	}
}

func (p PostgresSQLXRepository) Create(ctx context.Context, model Product) (err error) {
	query := `
		INSERT INTO products (
			name, category, price, stock
		) VALUES (
			:name, :category, :price, :stock
		)
	`

	stmt, err := p.db.PrepareNamed(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(model)

	return
}

func (p PostgresSQLXRepository) FindAll(ctx context.Context, models []Product) ([]Product, error) {
	query := `SELECT id, name, category, price, stock FROM products`

	err := p.db.Select(&models, query)
	if err != nil {
		return models, err
	}

	return models, nil
}

func (p PostgresSQLXRepository) FindByID(ctx context.Context, model Product, id int) (Product, error) {
	query := `SELECT id, name, category, price, stock FROM products WHERE id = $1`

	err := p.db.Get(&model, query, id)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (p PostgresSQLXRepository) Update(ctx context.Context, id int, model Product) error {
	query := `UPDATE products SET name = :name, category = :category, price = :price, stock = :stock WHERE id = :id`

	stmt, err := p.db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	model.Id = id

	_, err = stmt.Exec(model)
	if err != nil {
		return err
	}

	return err
}

func (p PostgresSQLXRepository) Delete(ctx context.Context, model Product, id int) error {
	query := `DELETE FROM products WHERE id = $1`

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return err
}
