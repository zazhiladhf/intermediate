package files

import (
	"product-catalog/pkg/images"
	"product-catalog/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouteFiles(router fiber.Router, client images.Cloudinary, cloud string, apiKey string, apiSecret string) {
	cloudSvc, _ := images.NewCloudinary(cloud, apiKey, apiSecret)
	handler := NewHandler(cloudSvc)

	v1 := router.Group("v1")
	files := v1.Group("/files")
	{
		files.Post("/upload", middleware.AuthMiddleware(), handler.Upload)
		// files.Post("/download", Download)
	}
}
