package product

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router, db *sql.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRoute := router.Group("products")
	{
		productRoute.Get("", handler.GetListProduct)
		productRoute.Post("", handler.CreateProduct)
		productRoute.Get("/sku/:sku", handler.GetProductDetail)
	}
}
