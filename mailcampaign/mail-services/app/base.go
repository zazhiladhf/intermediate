package app

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	svc := NewService()
	handler := NewHandler(svc)

	router.POST("/send", handler.SendMail)

}
