package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/omeid/pgerror"
)

type repository interface {
	Save(ctx context.Context, auth Auth) (err error)
	FindByEmail(ctx context.Context, email string) (auth Auth, err error)
	IsEmailAlreadyExists(ctx context.Context, email string) (bool, error)
	// readRepository
}

// type writeRepository interface {
// 	save(ctx context.Context, item Auth) (err error)
// }

// type readRepository interface {
// 	findByEmail(ctx context.Context, email string) (item Auth, err error)
// }

type AuthService struct {
	repo repository
}

func NewService(repo repository) AuthService {
	return AuthService{
		repo: repo,
	}
}

func (s AuthService) CreateAuth(ctx context.Context, req Auth) (err error) {
	err = req.ValidateFormRegister()
	if err != nil {
		log.Println("error when try to validate request with error")
		return
	}

	err = req.EncryptPassword()
	if err != nil {
		log.Println("error when try to encrypt password with error")
		return
	}

	email := req.Email

	_, err = s.repo.IsEmailAlreadyExists(ctx, email)
	if err != nil {
		// log.Println("auth:", auth)
		// log.Println("error sql:", err)
		if err == sql.ErrNoRows {
			err = s.repo.Save(ctx, req)
			if err != nil {
				return
			}
			return
		}
		if pgerror.UniqueViolation(err) != nil {
			return ErrDuplicateEmail
		}
		return
	}

	err = s.repo.Save(ctx, req)
	if err != nil {
		return
	}

	// if isExist {
	// 	log.Println("email already used with error:", err)
	// 	return err
	// }

	// model = req
	return
}

func (s AuthService) Login(ctx context.Context, req loginRequest) (item Auth, err error) {
	email := req.Email
	password := req.Password

	itemAuth, err := item.ValidateFormLogin(req)
	if err != nil {
		log.Println("error when try to validate request with error", err.Error(), itemAuth)
		return
	}

	itemAuth, err = s.repo.FindByEmail(ctx, email)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("error when try to findbyemail with error", err.Error(), itemAuth)
			return itemAuth, ErrRepository
		}
	}

	if itemAuth.Email == "" {
		log.Println("error when try to check email with error", err.Error(), itemAuth)
		return itemAuth, ErrInvalidEmail
	}

	ok, err := itemAuth.ValidatePassword(password)
	if err != nil {
		log.Println("error when try to validate password with error", err.Error(), itemAuth)
		return itemAuth, ErrInvalidPassword
	}

	if !ok {
		log.Println("error when try to !ok with error", err.Error(), itemAuth)
		return itemAuth, ErrInternalServer
	}

	return

}

func (s AuthService) GetAuthByEmail(ctx context.Context, email string) (auth Auth, err error) {
	auth, err = s.repo.FindByEmail(ctx, email)
	if err != nil {
		return auth, err
	}

	if auth.Id == 0 {
		return auth, errors.New("no auth found on with that email")
	}

	return auth, nil
}
