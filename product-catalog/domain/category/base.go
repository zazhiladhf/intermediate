package category

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegisterCategoriesRouter(router fiber.Router, dbSqlx *sqlx.DB) {
	repo := NewPostgreSqlxRepository(dbSqlx)
	svc := newService(repo)
	handler := newHandler(svc)

	v1 := router.Group("v1")
	v1.Get("/categories", handler.GetListCategories)

}
