package main

import (
	"log"
	"mailcampaign/app-services/app"
	"mailcampaign/config"
	appMailService "mailcampaign/mail-services/app"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func AppServicesServer() http.Handler {
	// setup routing
	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	app.RegisterRoutes(router)

	return router
}

func MailServicesServer() http.Handler {
	// setup routing
	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	appMailService.RegisterRoutes(router)

	return router
}

func main() {

	serverAppService := &http.Server{
		Addr:         config.Cfg.App.Port,
		Handler:      AppServicesServer(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":4001",
		Handler:      MailServicesServer(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := serverAppService.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		err := server02.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
