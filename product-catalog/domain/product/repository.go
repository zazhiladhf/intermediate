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

func (r PostgresSQLXRepository) InsertProduct(ctx context.Context, model Product) (id int, err error) {
	query := `
		INSERT INTO products (
			name, image_url, stock, price, category_id
		) VALUES (
			:name, :image_url, :stock, :price, :category_id
		)
		RETURNING id
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.GetContext(ctx, &id, model)

	return
}

// func (r PostgresSQLXRepository) FindAll(ctx context.Context, models []Product) ([]Product, error) {
// 	query := `SELECT id, name, category, price, stock FROM products`

// 	err := r.db.Select(&models, query)
// 	if err != nil {
// 		return models, err
// 	}

// 	return models, nil
// }

// func (p PostgresSQLXRepository) FindByID(ctx context.Context, model Product, id int) (Product, error) {
// 	query := `SELECT id, name, category, price, stock FROM products WHERE id = $1`

// 	err := p.db.Get(&model, query, id)
// 	if err != nil {
// 		return model, err
// 	}

// 	return model, nil
// }

// func (p PostgresSQLXRepository) Update(ctx context.Context, id int, model Product) error {
// 	query := `UPDATE products SET name = :name, category = :category, price = :price, stock = :stock WHERE id = :id`

// 	stmt, err := p.db.PrepareNamed(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	model.Id = id

// 	_, err = stmt.Exec(model)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

// func (p PostgresSQLXRepository) Delete(ctx context.Context, model Product, id int) error {
// 	query := `DELETE FROM products WHERE id = $1`

// 	stmt, err := p.db.PrepareContext(ctx, query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }
