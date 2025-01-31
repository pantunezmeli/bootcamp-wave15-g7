package main

import (
	"fmt"
	//"log"
	//"os"

	//"github.com/joho/godotenv"
	"github.com/pantunezmeli/bootcamp-wave15-g7/cmd/server"
)

func main() {

	cfg := &server.ConfigServerChi{
		ServerAddress: ":8080",
	}
	fmt.Printf("Running on localhost %s", cfg.ServerAddress)
	app := server.NewServerChi(cfg)

	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
