package auth

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

var (
	ErrEmailEmpty      = errors.New("email required")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrPasswordEmpty   = errors.New("password required")
	ErrInvalidPassword = errors.New("invalid password")
	ErrDuplicateEmail  = errors.New("email already used")
	ErrRepository      = errors.New("error repository")
	ErrInternalServer  = errors.New("unknown error")

	ErrCodeEmailEmpty      = errorCodeBadRequest("01")
	ErrCodeInvalidEmail    = errorCodeBadRequest("02")
	ErrCodePassworEmpty    = errorCodeBadRequest("03")
	ErrCodeInvalidPassword = errorCodeBadRequest("04")
	ErrCodeDuplicateEmail  = errorCodeConflict("01")
	ErrCodeInternalServer  = errorCodeInternalServer("01")
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

func (a Auth) FormRegister(req registerRequest) (Auth, error) {
	if req.Email == "" {
		return a, ErrEmailEmpty
	}

	if !valid(req.Email) {
		return a, ErrInvalidEmail
	}

	if req.Password == "" {
		return a, ErrPasswordEmpty
	}

	if len(req.Password) < 6 {
		return a, ErrInvalidPassword
	}

	a.Email = req.Email
	a.Password = req.Password
	a.Role = "merchant"
	return a, nil
}

func (a Auth) FormLogin(req login) (Auth, error) {
	if req.Email == "" {
		return a, ErrEmailEmpty
	}

	if !valid(req.Email) {
		return a, ErrInvalidEmail
	}

	if req.Password == "" {
		return a, ErrPasswordEmpty
	}

	if len(req.Password) <= 6 {
		return a, ErrInvalidPassword
	}

	a.Email = req.Email
	a.Password = req.Password
	return a, nil
}

func (a Auth) ValidatePassword(plainText string) (ok bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plainText))
	if err != nil {
		return ok, ErrInvalidPassword
	}
	ok = true
	return
}

func errorCodeBadRequest(n string) string {
	concatenated := fmt.Sprintf("%d%s", http.StatusBadRequest, n)
	return concatenated
}

func errorCodeConflict(n string) string {
	concatenated := fmt.Sprintf("%d%s", http.StatusConflict, n)
	return concatenated
}

func errorCodeInternalServer(n string) string {
	concatenated := fmt.Sprintf("%d%s", http.StatusInternalServerError, n)
	return concatenated
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
