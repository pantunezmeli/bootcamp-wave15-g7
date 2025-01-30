package main

import (
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/cmd/server"
)


func main() {
	// env
	// ...

	// app
	// - config
	cfg := &server.ConfigServerChi{
		ServerAddress:  ":8080",
		StorageFilePath: "",
	}
	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
