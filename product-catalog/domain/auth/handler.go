package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/omeid/pgerror"
)

type authHandler struct {
	svc authService
}

func newHandler(svc authService) authHandler {
	return authHandler{
		svc: svc,
	}
}

func (h authHandler) RegisterAuth(c *fiber.Ctx) (err error) {
	var req registerRequest

	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST body parser",
			"error":   err.Error(),
		})
	}

	model := Auth{
		Email:    req.Email,
		Password: req.Password,
		Role:     "merchant",
	}

	// model, err := NewAuth().ValidateFormRegister(req)
	// if err != nil {
	// 	log.Println(err)
	// 	return ResponseError(c, err)
	// }

	// isEmailAvailable, err := a.svc.isEmailAvailable(c.UserContext(), req)
	// if err != nil {
	// 	log.Println("err is available:", err)
	// 	log.Println("is available:", isEmailAvailable)
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "ERR BAD REQUEST",
	// 		"error":   err.Error(),
	// 	})
	// }

	// if !isEmailAvailable {
	// 	return c.Status(http.StatusConflict).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "duplicate email",
	// 		"error":   err.Error(),
	// 	})
	// }

	resp := Response{}

	err = h.svc.CreateAuth(c.UserContext(), model)
	if err != nil {
		switch {
		case err == ErrEmailEmpty:
			resp = Response{
				HttpCode:  http.StatusBadRequest,
				Success:   false,
				Message:   "bad request",
				Error:     err.Error(),
				ErrorCode: "40001",
			}
		case err == ErrInvalidEmail:
			resp = Response{
				HttpCode:  http.StatusBadRequest,
				Success:   false,
				Message:   "bad request",
				Error:     err.Error(),
				ErrorCode: "40002",
			}
		case err == ErrPasswordEmpty:
			resp = Response{
				HttpCode:  http.StatusBadRequest,
				Success:   false,
				Message:   "bad request",
				Error:     err.Error(),
				ErrorCode: "40003",
			}
		case err == ErrInvalidPassword:
			resp = Response{
				HttpCode:  http.StatusBadRequest,
				Success:   false,
				Message:   "bad request",
				Error:     err.Error(),
				ErrorCode: "40004",
			}
		case pgerror.UniqueViolation(err) != nil:
			resp = Response{
				HttpCode:  http.StatusConflict,
				Success:   false,
				Message:   "duplicate",
				Error:     err.Error(),
				ErrorCode: "40901",
			}
		default:
			resp = Response{
				HttpCode:  http.StatusInternalServerError,
				Success:   false,
				Message:   "internal server error",
				Error:     err.Error(),
				ErrorCode: "50001",
			}
		}
		resp.Error = err.Error()
		return c.Status(resp.HttpCode).JSON(resp)
	}
	resp = Response{
		HttpCode:  http.StatusCreated,
		Success:   true,
		Message:   "registration success",
		Error:     "",
		ErrorCode: "",
	}
	return c.Status(resp.HttpCode).JSON(resp)
}

// func (a authHandler) login(c *fiber.Ctx) error {
// 	var req login
// 	err := c.BodyParser(&req)
// 	if err != nil {
// 		return WriteError(c, err)
// 	}

// 	model, err := NewAuth().FormLogin(req)
// 	if err != nil {
// 		return WriteError(c, err)
// 	}

// 	item, err := a.svc.login(c.UserContext(), model)
// 	if err != nil {
// 		return WriteError(c, err)
// 	}

// 	return WriteSuccess(c, true, "login success", http.StatusOK, item)
// }

// func (a authHandler) CheckEmailAvailability(c *fiber.Ctx) (err error) {
// 	var req registerRequest

// 	err = c.BodyParser(&req)
// 	if err != nil {
// 		return WriteError(c, err)
// 	}

// 	isEmailAvailable, err := a.svc.isEmailAvailable(c.UserContext(), req)
// 	if err != nil {
// 		return WriteError(c, err)
// 	}

// 	data := gin.H{
// 		"is_available": isEmailAvailable,
// 	}

// 	metaMessage := "Email has been registered"

// 	if isEmailAvailable {
// 		metaMessage = "Email is available"
// 	}

//		response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
//		c.JSON(http.StatusUnprocessableEntity, response)
//	}
