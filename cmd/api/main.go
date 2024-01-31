package main

import (
	"log"

	"github.com/dickidarmawansaputra/go-simple-online-shop/app/auth"
	"github.com/dickidarmawansaputra/go-simple-online-shop/app/product"
	"github.com/dickidarmawansaputra/go-simple-online-shop/external/database"
	"github.com/dickidarmawansaputra/go-simple-online-shop/internal/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectMysql(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("DB connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	auth.Init(router, db)
	product.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
