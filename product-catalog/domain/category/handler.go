package category

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type categoryHandler struct {
	svc categoryService
}

func newHandler(svc categoryService) categoryHandler {
	return categoryHandler{
		svc: svc,
	}
}

func (h categoryHandler) GetListCategories(c *fiber.Ctx) error {
	listCategories, err := h.svc.getListCategories(c.UserContext())
	if err != nil {
		log.Println(listCategories)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success":    false,
			"message":    "internal server error",
			"error":      err.Error(),
			"error_code": "50001",
			// "payload": listCategories,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "get categories success",
		"payload": listCategories,
	})
}
