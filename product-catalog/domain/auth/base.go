package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Run(router fiber.Router, db *sqlx.DB) {
	repo := NewPostgreSqlxRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	v1 := router.Group("v1")
	auth := v1.Group("auth")
	{
		auth.Post("/register", handler.register)
		auth.Post("/signin", handler.login)
	}
	router.Get("", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	// auth := router.Group("auth")
	// {
	// 	auth.Post("/register", handler.signUp)
	// 	auth.Post("/signin", handler.signIn)
	// }
}
