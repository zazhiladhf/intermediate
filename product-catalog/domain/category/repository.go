package category

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrCategoriesNotFound = errors.New("categories not found")
)

type PostgreSqlxRepository struct {
	db *sqlx.DB
}

func NewPostgreSqlxRepository(db *sqlx.DB) PostgreSqlxRepository {
	return PostgreSqlxRepository{
		db: db,
	}
}

func (p PostgreSqlxRepository) FindAll(ctx context.Context) (items []category, err error) {
	query := `
    	SELECT 
			id, name 
    	FROM categories
    `

	err = p.db.SelectContext(ctx, &items, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoriesNotFound
		}
		return
	}

	if len(items) == 0 {
		return nil, ErrCategoriesNotFound
	}

	return
}
