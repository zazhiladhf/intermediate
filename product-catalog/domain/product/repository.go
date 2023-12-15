package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	ErrCategoriesNotFound = errors.New("categories not found")
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
			name, image_url, stock, price, category_id, email_auth
		) VALUES (
			:name, :image_url, :stock, :price, :category_id, :email_auth
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

func (r PostgresSQLXRepository) FindAll(ctx context.Context) (list []Product, err error) {
	query := `
		SELECT 
			id, name, image_url, stock, price, category_id, email_auth
		FROM products
	`

	err = r.db.SelectContext(ctx, &list, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoriesNotFound
		}
		return
	}

	if len(list) == 0 {
		return nil, ErrCategoriesNotFound
	}

	return list, nil
}

func (p PostgresSQLXRepository) FindByEmail(ctx context.Context, queryParam string, email string) (list []Product, err error) {
	filter := mappingQueryFilter(queryParam)

	queryByEmail := `
		SELECT 
		id, name, image_url, stock, price, category_id, email_auth 
		FROM products 
		WHERE email_auth = $1
	`

	query := fmt.Sprintf("%s %s", queryByEmail, filter)

	err = p.db.SelectContext(ctx, &list, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return []Product{}, ErrCategoriesNotFound
		}
		return
	}

	if len(list) == 0 {
		return nil, ErrCategoriesNotFound
	}

	return list, nil
}

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

func mappingQueryFilter(queryParam string) string {
	filter := ""

	if queryParam != "" {
		filter = fmt.Sprintf("%s AND name ILIKE '%%%s%%'", filter, queryParam)
	}

	return filter
}
