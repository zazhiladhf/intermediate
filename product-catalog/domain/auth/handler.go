package auth

import (
	"context"
	"log"
	"net/http"
	"product-catalog/config"
	"product-catalog/pkg/database"
	auth "product-catalog/pkg/jwt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/omeid/pgerror"
)

type AuthHandler struct {
	svc AuthService
}

func newHandler(svc AuthService) AuthHandler {
	return AuthHandler{
		svc: svc,
	}
}

func (h AuthHandler) RegisterAuth(c *fiber.Ctx) (err error) {
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
		HttpCode: http.StatusCreated,
		Success:  true,
		Message:  "registration success",
		// Error:     "",
		// ErrorCode: "",
		// Payload:   "",
	}
	return c.Status(resp.HttpCode).JSON(resp)
}

func (h AuthHandler) Login(c *fiber.Ctx) error {
	var req loginRequest
	var resp Response

	err := c.BodyParser(&req)
	if err != nil {
		resp = Response{
			HttpCode:  http.StatusBadRequest,
			Success:   false,
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: "400",
			// Payload:   "",
		}
		return c.Status(resp.HttpCode).JSON(resp)
	}

	// itemAuth, err := NewAuth().ValidateFormLogin(req)

	// resp = Response{
	// 	HttpCode:  http.StatusBadRequest,
	// 	Success:   false,
	// 	Message:   "bad request",
	// 	Error:     err.Error(),
	// 	ErrorCode: "400",
	// 	// Payload:   "",
	// }
	// return c.Status(resp.HttpCode).JSON(resp)

	// itemAuth = Auth{
	// 	Email:    itemAuth.Email,
	// 	Password: itemAuth.Password,
	// 	Role:     itemAuth.Role,
	// }
	// log.Println(itemAuth)

	itemAuth, err := h.svc.Login(c.UserContext(), req)
	if err != nil {
		log.Println("error when trying to login with error:", err.Error(), req)
		switch err {
		case ErrEmailEmpty:
			resp = Response{
				HttpCode:  http.StatusBadRequest,
				Success:   false,
				Message:   "bad request",
				Error:     err.Error(),
				ErrorCode: "40001",
			}
		case ErrInvalidEmail:
			resp = Response{
				HttpCode:  http.StatusBadRequest,
				Success:   false,
				Message:   "bad request",
				Error:     err.Error(),
				ErrorCode: "40002",
			}
		case ErrPasswordEmpty:
			resp = Response{
				HttpCode:  http.StatusBadRequest,
				Success:   false,
				Message:   "bad request",
				Error:     err.Error(),
				ErrorCode: "40003",
			}
		case ErrInvalidPassword:
			resp = Response{
				HttpCode:  http.StatusBadRequest,
				Success:   false,
				Message:   "bad request",
				Error:     err.Error(),
				ErrorCode: "40004",
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
		// resp.Error = err.Error()
		return c.Status(resp.HttpCode).JSON(resp)
	}

	token, err := auth.GenerateNewAccessToken()
	if err != nil {
		log.Println("error when trying to generate toker with error:", err.Error())
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(24*time.Hour))
	defer cancel()
	client, err := database.ConnectRedis(config.Cfg.Redis)
	if err != nil {
		log.Println("error when to try migration db with error :", err.Error())
		resp = Response{
			HttpCode:  http.StatusInternalServerError,
			Success:   false,
			Message:   "internal server error",
			Error:     err.Error(),
			ErrorCode: "50001",
		}
		return c.Status(resp.HttpCode).JSON(resp)
		// panic(err)
	}

	if client == nil {
		log.Println("db not connected with unknown error")
		resp = Response{
			HttpCode:  http.StatusInternalServerError,
			Success:   false,
			Message:   "internal server error",
			Error:     err.Error(),
			ErrorCode: "999999",
		}
		return c.Status(resp.HttpCode).JSON(resp)
	}

	err = client.Set(ctx, itemAuth.Email, token, 24*time.Hour).Err()
	if err != nil {
		log.Println("error when try to set data to redis with message :", err.Error())
		resp = Response{
			HttpCode:  http.StatusInternalServerError,
			Success:   false,
			Message:   "internal server error",
			Error:     err.Error(),
			ErrorCode: "50001",
		}
		return c.Status(resp.HttpCode).JSON(resp)
	}

	// resp := Response{}

	resp = Response{
		HttpCode: http.StatusOK,
		Success:  true,
		Message:  "login success",
		// Error:     "",
		// ErrorCode: "",
		Payload: Payload{
			AccessToken: token,
			Role:        itemAuth.Role,
		},
	}
	return c.Status(resp.HttpCode).JSON(resp)
}

// func (a AuthHandler) CheckEmailAvailability(c *fiber.Ctx) (err error) {
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
