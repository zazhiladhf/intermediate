package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Run(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)
	auth := router.Group("auth")
	{
		auth.Post("/signup", handler.signUp)
		auth.Post("/signin", handler.signIn)
	}
}
