package auth

import (
	"context"
	"database/sql"
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

type authService struct {
	repo repository
}

func newService(repo repository) authService {
	return authService{
		repo: repo,
	}
}

func (s authService) CreateAuth(ctx context.Context, req Auth) (err error) {
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

// func (a authService) login(ctx context.Context, req Auth) (item Auth, err error) {
// 	auth, err := a.repo.findByEmail(ctx, req.Email)
// 	if err != nil {
// 		return
// 	}

// 	ok, err := auth.ValidatePassword(req.Password)
// 	if err != nil {
// 		return req, err
// 	}

// 	if !ok {
// 		return req, ErrInvalidPassword
// 	}

// 	return auth, nil

// }

// func (a authService) isEmailAvailable(ctx context.Context, req registerRequest) (bool, error) {
// 	// var req registerRequest
// 	email := req.Email

// 	auth, err := a.repo.findByEmail(ctx, email)
// 	if err != nil {
// 		log.Println("auth:", auth)
// 		log.Println("error sql:", err)
// 		return true, err
// 	}

// 	if auth.Id == 0 {
// 		log.Println("error error id == 0:", err)

// 		return false, nil
// 	}

// 	// if auth.Email == req.Email {
// 	// 	log.Println("error email sama:", err)

// 	// 	return false, err
// 	// }

// 	return false, nil
// }

// func (a authService) CreateAuth() {
// 	// check to database
// 	isExists, err := a.repo.IsEmailAlreadyExists(req.Email)
// 	if err != nil {
// 		if err != sql.ErrNoRows {
// 			return ErrorRepository
// 		}
// 	}

// 	if isExists {
// 		return ErrorDuplicateEntry
// 	}

// 	// insert into auth
// }
