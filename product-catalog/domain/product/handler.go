package product

import (
	"net/http"
	"product-catalog/pkg/jwt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	svc Service
}

func NewHandler(svc Service) productHandler {
	return productHandler{
		svc: svc,
	}
}

func (h productHandler) CreateProduct(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greater than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	var req = CreateProductRequest{}

	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	// currentAuth := c.GetReqHeaders().(auth.Auth)
	// req.Auth = currentAuth

	// model := Product{
	// 	Name:  req.Name,
	// 	Price: req.Price,
	// 	Stock: req.Stock,
	// }

	// client, err := database.ConnectRedis(config.Cfg.Redis)
	// if err != nil {
	// 	log.Println("error when to try migration db with error :", err.Error())
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "ERR BAD REQUEST",
	// 		"error":   err.Error(),
	// 	})
	// }

	// if client == nil {
	// 	log.Println("db not connected with unknown error")
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "ERR BAD REQUEST",
	// 		"error":   err.Error(),
	// 	})
	// }

	// var itemAuth auth.Auth
	// token := client.Get(c.UserContext(), itemAuth.Email)
	// if err != nil {
	// 	log.Println("error when try to get data to redis with message :", err.Error())
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "ERR BAD REQUEST",
	// 		"error":   err.Error(),
	// 	})
	// }

	// claim, err := jwt.ValidateToken(token.String())
	// if err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }

	err = h.svc.createProduct(c.UserContext(), req)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		switch err {
		case ErrEmptyName, ErrEmptyImageURL, ErrEmptyPrice, ErrEmptyStock:
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}
			httpCode = http.StatusBadRequest
		default:
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   err.Error(),
			}
			httpCode = http.StatusInternalServerError
		}
		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "CREATE SUCCESS",
	})
}

// func (h Handler) GetProducts(c *fiber.Ctx) error {
// 	var models []Product
// 	products, err := h.svc.GetProducts(c.UserContext(), models)

// 	if err != nil {
// 		var payload fiber.Map
// 		httpCode := 404

// 		switch err {
// 		case ErrNotFound:
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR NOT FOUND",
// 				"error":   err.Error(),
// 			}
// 			httpCode = http.StatusNotFound
// 		default:
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR INTERNAL",
// 				"error":   err.Error(),
// 			}
// 			httpCode = http.StatusInternalServerError
// 		}
// 		return c.Status(httpCode).JSON(payload)
// 	}

// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"message": "GET ALL SUCCESS",
// 		"payload": products,
// 	})
// }

// func (h Handler) GetProductById(c *fiber.Ctx) error {
// 	model := Product{}
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   "invalid id",
// 		})
// 	}

// 	product, err := h.svc.GetProductById(c.UserContext(), model, id)
// 	if err != nil {
// 		var payload fiber.Map
// 		httpCode := 400

// 		if model.Id != id {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   err.Error(),
// 			}
// 			httpCode = http.StatusNotFound
// 		} else if err == ErrEmptyName || err == ErrEmptyImageURL || err == ErrEmptyPrice || err == ErrEmptyStock {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   "invaid id",
// 			}
// 			httpCode = http.StatusBadRequest
// 		} else {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR INTERNAL",
// 				"error":   "ada masalah pada server",
// 			}
// 			httpCode = http.StatusInternalServerError
// 		}

// 		return c.Status(httpCode).JSON(payload)
// 	}
// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"message": "GET DATA SUCCESS",
// 		"payload": product,
// 	})

// }

// func (h Handler) UpdateProduct(c *fiber.Ctx) error {
// 	var model Product
// 	var req = CreateProductRequest{}
// 	err := c.BodyParser(&req)

// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   err.Error(),
// 		})
// 	}

// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   "invalid id",
// 		})
// 	}

// 	product, err := h.svc.GetProductById(c.UserContext(), model, id)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   err.Error(),
// 		})
// 	}

// 	product.Name = req.Name
// 	product.Price = req.Price
// 	product.Stock = req.Stock

// 	err = h.svc.UpdateProduct(c.UserContext(), product, id)
// 	if err != nil {
// 		var payload fiber.Map
// 		httpCode := 400

// 		if model.Id != id {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   err.Error(),
// 			}
// 			httpCode = http.StatusNotFound
// 		} else if err == ErrEmptyName || err == ErrEmptyImageURL || err == ErrEmptyPrice || err == ErrEmptyStock {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   "invaid id",
// 			}
// 			httpCode = http.StatusBadRequest
// 		} else {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR INTERNAL",
// 				"error":   "ada masalah pada server",
// 			}
// 			httpCode = http.StatusInternalServerError
// 		}

// 		return c.Status(httpCode).JSON(payload)
// 	}
// 	return c.Status(http.StatusCreated).JSON(fiber.Map{
// 		"success": true,
// 		"message": "UPDATE SUCCESS",
// 	})
// }

// func (h Handler) DeleteProduct(c *fiber.Ctx) error {
// 	model := Product{}
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   "invalid id",
// 		})
// 	}

// 	product, err := h.svc.GetProductById(c.UserContext(), model, id)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"message": "ERR BAD REQUEST",
// 			"error":   err.Error(),
// 		})
// 	}

// 	err = h.svc.DeleteProduct(c.UserContext(), product, id)
// 	if err != nil {
// 		var payload fiber.Map
// 		httpCode := 400

// 		if model.Id != id {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   err.Error(),
// 			}
// 			httpCode = http.StatusNotFound
// 		} else if err == ErrEmptyName || err == ErrEmptyImageURL || err == ErrEmptyPrice || err == ErrEmptyStock {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR BAD REQUEST",
// 				"error":   "invaid id",
// 			}
// 			httpCode = http.StatusBadRequest
// 		} else {
// 			payload = fiber.Map{
// 				"success": false,
// 				"message": "ERR INTERNAL",
// 				"error":   "ada masalah pada server",
// 			}
// 			httpCode = http.StatusInternalServerError
// 		}

// 		return c.Status(httpCode).JSON(payload)
// 	}
// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"message": "DELETE DATA SUCCESS",
// 	})

// }
