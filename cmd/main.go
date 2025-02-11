package main

import (
	"log"

	"github.com/pantunezmeli/bootcamp-wave15-g7/cmd/server"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/config"
)

func main() {
	serverConfig := &server.ConfigServerChi{
		ServerAddress: config.GetEnv("SERVER_ADDRESS", ":8080"),
	}

	app := server.NewServerChi(serverConfig)

	if err := app.Run(); err != nil {
		log.Fatal("Error running server: %v", err)
		return
	}
}
