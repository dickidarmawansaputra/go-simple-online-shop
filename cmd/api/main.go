package main

import (
	"log"

	"github.com/dickidarmawansaputra/go-simple-online-shop/external/database"
	"github.com/dickidarmawansaputra/go-simple-online-shop/internal/config"
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
}
