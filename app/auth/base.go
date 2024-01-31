package auth

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router, db *sql.DB) {
	repo := newRepository(db)
	scv := newService(repo)
	handler := newHandler(scv)

	authRouter := router.Group("auth")
	{
		authRouter.Post("register", handler.register)
		authRouter.Post("login", handler.login)
	}
}
