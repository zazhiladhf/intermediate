package files

import (
	"product-catalog/pkg/images"
	"product-catalog/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutesFile(router fiber.Router, client images.Cloudinary, cloud string, apiKey string, apiSecret string) {
	cloudSvc, _ := images.NewCloudinary(cloud, apiKey, apiSecret)
	handler := NewHandler(cloudSvc)

	filesRouter := router.Group("/v1/files")
	{
		filesRouter.Post("/upload", middleware.AuthMiddleware(), handler.Upload)
	}
}
