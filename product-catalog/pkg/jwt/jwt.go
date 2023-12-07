package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// type Service interface {
// 	GenerateToken(userID int) (string, error)
// 	ValidateToken(encodedToken string) (*jwt.Token, error)
// }

type JwtService struct {
}

// var Jwt *JwtService

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

func NewService() *JwtService {
	return &JwtService{}
}

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   jwt.NewNumericDate(time.Now().Add(10 * time.Minute)).Unix(),
	}
	// claim["user_id"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
