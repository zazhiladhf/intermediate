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

func (r PostgreSqlxRepository) FindAll(ctx context.Context) (categories []Category, err error) {
	query := `
    	SELECT 
			id, name 
    	FROM categories
    `

	err = r.db.SelectContext(ctx, &categories, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoriesNotFound
		}
		return
	}

	if len(categories) == 0 {
		return nil, ErrCategoriesNotFound
	}

	return
}
