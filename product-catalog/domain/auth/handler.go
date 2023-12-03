package auth

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	svc authService
}

func newHandler(svc authService) authHandler {
	return authHandler{
		svc: svc,
	}
}

func (a authHandler) register(c *fiber.Ctx) error {
	var req registerRequest

	err := c.BodyParser(&req)
	if err != nil {
		return WriteError(c, err)
	}

	model, err := NewAuth().FormRegister(req)
	if err != nil {
		log.Println(err)
		return WriteError(c, err)
	}

	err = a.svc.register(c.UserContext(), model)
	if err != nil {
		log.Println(err)
		return WriteError(c, err)
	}

	// resp := newRegisterResponse(item)

	return WriteSuccess(c, true, "registration success", http.StatusCreated, nil)
}

func (a authHandler) login(c *fiber.Ctx) error {
	var req login
	if err := c.BodyParser(&req); err != nil {
		return WriteError(c, err)
	}

	model, err := NewAuth().FormLogin(req)
	if err != nil {
		return WriteError(c, err)
	}

	item, err := a.svc.login(c.UserContext(), model)
	if err != nil {
		return WriteError(c, err)
	}

	return WriteSuccess(c, true, "login success", http.StatusOK, item)
}
