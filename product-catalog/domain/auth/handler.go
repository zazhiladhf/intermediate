package auth

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type AuthHandler struct {
	svc AuthService
}

func NewHandler(svc AuthService) AuthHandler {
	return AuthHandler{
		svc: svc,
	}
}

func (h AuthHandler) Register(c *fiber.Ctx) (err error) {
	var req registerRequest

	err = c.BodyParser(&req)
	if err != nil {
		log.Println("error when try to parsing body request with error", err)
		return ResponseError(c, err)
	}

	model, err := NewAuth().ValidateFormRegister(req)
	if err != nil {
		log.Println("error when try to validate form register with error", err)
		return ResponseError(c, err)
	}

	err = h.svc.RegisterAuth(c.UserContext(), model)
	if err != nil {
		log.Println("error when try to register with error", err)
		pqErr, ok := err.(*pq.Error)
		if ok {
			switch pqErr.Code {
			case "23505":
				return ResponseError(c, ErrDuplicateEmail)
			default:
				return ResponseError(c, ErrRepository)
			}
		} else {
			log.Println("unknown error with error:", err)
		}

		return ResponseError(c, err)
	}

	return ResponseSuccess(c, true, "registration success", http.StatusCreated, nil)
}

func (h AuthHandler) Login(c *fiber.Ctx) error {
	var req loginRequest

	err := c.BodyParser(&req)
	if err != nil {
		log.Println("error when try to parsing body request with error", err)
		return ResponseError(c, err)
	}

	model, err := NewAuth().ValidateFormLogin(req)
	if err != nil {
		log.Println("error when try to validate form login with error", err)
		return ResponseError(c, err)
	}

	itemAuth, token, err := h.svc.Login(c.UserContext(), model)
	if err != nil {
		log.Println("error when try to login with error", err)
		return ResponseError(c, err)
	}

	return ResponseSuccess(c, true, "login success", http.StatusOK, Payload{
		AccessToken: token,
		Role:        itemAuth.Role,
	})
}
