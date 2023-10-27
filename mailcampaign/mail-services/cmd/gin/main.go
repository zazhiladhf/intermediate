package main

import (
	"log"
	"mailcampaign/mail-services/app"
	"mailcampaign/mail-services/config"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// setup routing
	router := gin.New()
	router.Use(cors.Default())

	err := config.LoadConfig("./mail-services/config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	//handler with basic routing
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
			"version": "not set",
		})
	})

	app.RegisterRoutes(router)

	router.Run(config.Cfg.App.Port)
}
