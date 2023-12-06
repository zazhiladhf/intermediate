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

func (p PostgreSqlxRepository) Save(ctx context.Context, auth Auth) (err error) {
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

func (p PostgreSqlxRepository) FindByEmail(ctx context.Context, email string) (auth Auth, err error) {
	query := `SELECT id, email, password, role FROM auths WHERE email = $1`

	err = p.db.Get(&auth, query, email)
	if err != nil {
		return
	}

	return
}

func (p PostgreSqlxRepository) IsEmailAlreadyExists(ctx context.Context, email string) (bool, error) {
	var auth Auth
	query := `SELECT id, email, password, role FROM auths WHERE email = $1`

	err := p.db.Get(&auth, query, email)
	if err != nil {
		return false, err
	}

	return true, err
}
