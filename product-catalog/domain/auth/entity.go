package auth

import (
	"errors"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailEmpty      = errors.New("email required")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrPasswordEmpty   = errors.New("password required")
	ErrInvalidPassword = errors.New("invalid password")
	ErrDuplicateEmail  = errors.New("email already used")
	ErrRepository      = errors.New("error repository")
	ErrInternalServer  = errors.New("unknown error")
)

type Auth struct {
	Id       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) ValidateFormRegister() (err error) {
	if a.Email == "" {
		return ErrEmailEmpty
	}

	if !valid(a.Email) {
		return ErrInvalidEmail
	}

	if a.Password == "" {
		return ErrPasswordEmpty
	}

	if len(a.Password) < 6 {
		return ErrInvalidPassword
	}

	// if err != nil && errors.Is(err, UniqueViolationErr) {
	// 	return ErrDuplicateEmail
	// }

	// if err.Error() == "email already used" {
	// 	return ErrDuplicateEmail
	// }

	// if a.Email == a.Email {
	// 	return ErrDuplicateEmail
	// }

	// if e := pgerror.UniqueViolation(err); e != nil {
	// 	// you can use e here to check the fields et al
	// 	return ErrDuplicateEmail
	// }

	// if err != nil && pgerror.UniqueViolation(err) != nil {
	// 	return ErrDuplicateEmail
	// }

	// var duplicateEntryError = &pgconn.PgError{Code: "23505"}
	// if errors.As(err, &duplicateEntryError) {
	// 	return ErrDuplicateEmail
	// }

	// a.Email = req.Email
	// a.Password = req.Password
	// a.Role = "merchant"
	return nil
}

func (a Auth) ValidateFormLogin(req loginRequest) (Auth, error) {
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
	return a, nil
}

func (a *Auth) EncryptPassword() (err error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	a.Password = string(encrypted)
	return
}

func (a Auth) ValidatePassword(plainText string) (ok bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plainText))
	if err != nil {
		return ok, ErrInvalidPassword
	}
	ok = true
	return
}

// func errorCodeBadRequest(n string) string {
// 	concatenated := fmt.Sprintf("%d%s", http.StatusBadRequest, n)
// 	return concatenated
// }

// func errorCodeConflict(n string) string {
// 	concatenated := fmt.Sprintf("%d%s", http.StatusConflict, n)
// 	return concatenated
// }

// func errorCodeInternalServer(n string) string {
// 	concatenated := fmt.Sprintf("%d%s", http.StatusInternalServerError, n)
// 	return concatenated
// }

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// func (a Auth) isUsed(email string) bool {
// 	var ctx context.Context
// 	var req registerRequest
// 	usedEmail, _ := newHandler(authService{}).svc.IsEmailAvailable(ctx, req)
// 	return usedEmail
// }
