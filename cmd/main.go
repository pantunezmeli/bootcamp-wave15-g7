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
	fmt.Println("Proyecto corriendo...")
	cfg := &server.ConfigServerChi{
		ServerAddress:           ":8080",
		BuyerLoaderFilePath:     "../docs/db/buyer_data.json",
		WarehouseLoaderFilePath: "../docs/db/warehouse_data.json",
		EmployeeLoaderFilePath:  "../docs/db/employee_data.json",
	}
	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
