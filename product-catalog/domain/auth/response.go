package auth

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Error     *string     `json:"error,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
}

func WriteError(c *fiber.Ctx, err error) error {
	switch {
	case err == ErrEmailEmpty:
		return write(c, "bad request", err.Error(), ErrCodeEmailEmpty, nil)
	case err == ErrInvalidEmail:
		return write(c, "bad request", err.Error(), ErrCodeInvalidEmail, nil)
	case err == ErrPasswordEmpty:
		return write(c, "bad request", err.Error(), ErrCodePassworEmpty, nil)
	case err == ErrInvalidPassword:
		return write(c, "bad request", err.Error(), ErrCodeInvalidPassword, nil)
	default:
		return write(c, "internal server error", "error repository", ErrCodeInternalServer, nil)
	}
}

func WriteSuccess(c *fiber.Ctx, success bool, message string, statusCode int, payload interface{}) error {
	resp := Response{
		Success: success,
		Message: message,
		// ErrorCode: errorCode,
		// Payload: payload,
	}
	c = c.Status(statusCode)
	return c.JSON(resp)
}

func write(c *fiber.Ctx, message string, err string, statusCode string, payload interface{}) error {
	var httpStatusCode int
	c = c.Status(httpStatusCode)
	isSuccess := httpStatusCode >= 200 && httpStatusCode < 300

	var resp Response
	if isSuccess {
		resp = Response{
			Success:   true,
			Message:   message,
			Error:     &err,
			ErrorCode: statusCode,
			// StatusCode: statusCode,
			// Payload:     payload,
		}
	} else {
		resp = Response{
			Success:   false,
			Message:   message,
			Error:     &err,
			ErrorCode: statusCode,
		}
	}

	return c.JSON(resp)
}
