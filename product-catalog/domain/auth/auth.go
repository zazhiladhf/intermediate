package auth

import (
	"errors"

	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Id       ulid.ULID `json:"id" db:"id"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	Role     string    `json:"role" db:"role"`
}

var (
	ErrEmailEmpty      = errors.New("email required")
	ErrEmailNotFound   = errors.New("invalid email")
	ErrPasswordEmpty   = errors.New("password required")
	ErrInvalidPassword = errors.New("invalid password")
	ErrPasswordLength  = errors.New("password length must be equal to or greater than 6")
)

func NewAuth() Auth {
	return Auth{}
}

func (a *Auth) EncryptPassword() (err error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	a.Password = string(encrypted)
	return
}

func (a Auth) FromRegister(req register) (Auth, error) {
	if req.Email == "" {
		return a, ErrEmailEmpty
	}

	if req.Password == "" {
		return a, ErrPasswordEmpty
	}

	if len(req.Password) < 6 {
		return a, ErrPasswordLength
	}

	a.Email = req.Email
	a.Password = req.Password
	a.Id = ulid.Make()
	return a, nil
}

func (a Auth) FromLogin(req login) (Auth, error) {
	if req.Email == "" {
		return a, ErrEmailEmpty
	}

	if req.Password == "" {
		return a, ErrPasswordEmpty
	}

	if len(req.Password) <= 6 {
		return a, ErrPasswordLength
	}

	a.Email = req.Email
	a.Password = req.Password
	return a, nil
}

func (a Auth) ValidatePasswordFromPlainText(plainText string) (ok bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plainText))
	if err != nil {
		return ok, ErrInvalidPassword
	}
	ok = true
	return
}
