package product

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return Handler{
		svc: svc,
	}
}

func (h Handler) CreateProduct(c *fiber.Ctx) error {
	var req = CreateProductRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	model := Product{
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	err = h.svc.CreateProduct(c.UserContext(), model)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		switch err {
		case ErrEmptyName, ErrEmptyCategory, ErrEmptyPrice, ErrEmptyStock:
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

func (h Handler) GetProducts(c *fiber.Ctx) error {
	var models []Product
	products, err := h.svc.GetProducts(c.UserContext(), models)

	if err != nil {
		var payload fiber.Map
		httpCode := 404

		switch err {
		case ErrNotFound:
			payload = fiber.Map{
				"success": false,
				"message": "ERR NOT FOUND",
				"error":   err.Error(),
			}
			httpCode = http.StatusNotFound
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

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GET ALL SUCCESS",
		"payload": products,
	})
}

func (h Handler) GetProductById(c *fiber.Ctx) error {
	model := Product{}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   "invalid id",
		})
	}

	product, err := h.svc.GetProductById(c.UserContext(), model, id)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		if model.Id != id {
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}
			httpCode = http.StatusNotFound
		} else if err == ErrEmptyName || err == ErrEmptyCategory || err == ErrEmptyPrice || err == ErrEmptyStock {
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   "invaid id",
			}
			httpCode = http.StatusBadRequest
		} else {
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   "ada masalah pada server",
			}
			httpCode = http.StatusInternalServerError
		}

		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "GET DATA SUCCESS",
		"payload": product,
	})

}

func (h Handler) UpdateProduct(c *fiber.Ctx) error {
	var model Product
	var req = CreateProductRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   "invalid id",
		})
	}

	product, err := h.svc.GetProductById(c.UserContext(), model, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	product.Name = req.Name
	product.Category = req.Category
	product.Price = req.Price
	product.Stock = req.Stock

	err = h.svc.UpdateProduct(c.UserContext(), product, id)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		if model.Id != id {
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}
			httpCode = http.StatusNotFound
		} else if err == ErrEmptyName || err == ErrEmptyCategory || err == ErrEmptyPrice || err == ErrEmptyStock {
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   "invaid id",
			}
			httpCode = http.StatusBadRequest
		} else {
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   "ada masalah pada server",
			}
			httpCode = http.StatusInternalServerError
		}

		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "UPDATE SUCCESS",
	})
}

func (h Handler) DeleteProduct(c *fiber.Ctx) error {
	model := Product{}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   "invalid id",
		})
	}

	product, err := h.svc.GetProductById(c.UserContext(), model, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	err = h.svc.DeleteProduct(c.UserContext(), product, id)
	if err != nil {
		var payload fiber.Map
		httpCode := 400

		if model.Id != id {
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   err.Error(),
			}
			httpCode = http.StatusNotFound
		} else if err == ErrEmptyName || err == ErrEmptyCategory || err == ErrEmptyPrice || err == ErrEmptyStock {
			payload = fiber.Map{
				"success": false,
				"message": "ERR BAD REQUEST",
				"error":   "invaid id",
			}
			httpCode = http.StatusBadRequest
		} else {
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   "ada masalah pada server",
			}
			httpCode = http.StatusInternalServerError
		}

		return c.Status(httpCode).JSON(payload)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "DELETE DATA SUCCESS",
	})

}
