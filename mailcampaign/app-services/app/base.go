package app

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// svc := NewService()
	handler := NewHandler()

	router.POST("/send", handler.HandlerRequest)

}
