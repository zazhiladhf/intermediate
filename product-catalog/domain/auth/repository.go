package auth

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type PostgreSqlxRepository struct {
	db *sqlx.DB
}

func NewPostgreSqlxRepository(db *sqlx.DB) PostgreSqlxRepository {
	return PostgreSqlxRepository{
		db: db,
	}
}

func (p PostgreSqlxRepository) save(ctx context.Context, auth Auth) (err error) {
	query := `
		INSERT INTO auths (
			email, password, role
		) VALUES (
			:email, :password, :role
		)
	`

	stmt, err := p.db.PrepareNamed(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(auth)

	return
}

func (p PostgreSqlxRepository) findByEmail(ctx context.Context, email string) (auth Auth, err error) {
	query := `SELECT id, email, password, role FROM auths WHERE email = $1`

	err = p.db.Get(&auth, query, email)
	if err != nil {
		return
	}

	return
}
