package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func WriteError(c *fiber.Ctx, err error) error {
	switch {
	case err == ErrEmailEmpty || err == ErrPasswordEmpty || err == ErrPasswordLength:
		return write(c, http.StatusBadRequest, err.Error(), nil)
	default:
		return write(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func WriteSuccess(c *fiber.Ctx, payload interface{}, statusCode int) error {
	resp := response{
		StatusCocde: statusCode,
		Payload:     payload,
	}
	c = c.Status(statusCode)
	return c.JSON(resp)
}

type response struct {
	StatusCocde int         `json:"status_code"`
	Error       *string     `json:"error,omitempty"`
	Payload     interface{} `json:"payload,omitempty"`
}

func write(c *fiber.Ctx, statusCode int, message string, payload interface{}) error {
	c = c.Status(statusCode)
	isSuccess := statusCode >= 200 && statusCode < 300

	resp := response{}
	if isSuccess {
		resp = response{
			StatusCocde: statusCode,
			Payload:     payload,
		}
	} else {
		resp = response{
			StatusCocde: statusCode,
			Error:       &message,
		}
	}

	return c.JSON(resp)
}
