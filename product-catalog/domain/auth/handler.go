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

func (a authHandler) signUp(c *fiber.Ctx) error {
	var req register

	if err := c.BodyParser(&req); err != nil {
		return WriteError(c, err)
	}

	model, err := NewAuth().FromRegister(req)
	if err != nil {
		log.Println(err)
		return WriteError(c, err)
	}

	item, err := a.svc.register(c.UserContext(), model)
	if err != nil {
		log.Println(err)
		return WriteError(c, err)
	}

	resp := newRegisterResponse(item)

	return WriteSuccess(c, resp, http.StatusCreated)
}

func (a authHandler) signIn(c *fiber.Ctx) error {
	var req login
	if err := c.BodyParser(&req); err != nil {
		return WriteError(c, err)
	}

	model, err := NewAuth().FromLogin(req)
	if err != nil {
		return WriteError(c, err)
	}

	item, err := a.svc.login(c.UserContext(), model)
	if err != nil {
		return WriteError(c, err)
	}

	return WriteSuccess(c, item, http.StatusOK)
}
