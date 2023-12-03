package auth

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		db: db,
	}
}

func (p Repository) save(ctx context.Context, auth Auth) (err error) {
	query := `
		INSERT INTO auth (
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

func (p Repository) findByEmail(ctx context.Context, email string) (auth Auth, err error) {
	query := `SELECT id, email, password, role FROM auth WHERE email = $1`

	err = p.db.Get(&auth, query, email)
	if err != nil {
		return
	}

	return
}
