package main

import (
	"log"

	database "github.com/nioWhiteHat/gaming-store-backend.git/internal/Database"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/config"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/server"
)

func main() {
	
	cfg := config.LoadDBConfig()
	pool,err := database.NewDBPool(cfg.DSNForApp)
	if err!=nil{
		log.Fatal(err)
	}
	addr := server.Config{
		Addr: ":8080",
	}
	app := server.Application{Db:pool, Config: addr }
	mux := app.Mount()

	err = app.Run(mux)
	if err!=nil{
		log.Fatalf("%s ",err)
	}

}