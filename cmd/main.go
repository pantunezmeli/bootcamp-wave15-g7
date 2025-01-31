package main

import (
	"fmt"
	//"log"
	//"os"

	//"github.com/joho/godotenv"
	"github.com/pantunezmeli/bootcamp-wave15-g7/cmd/server"
)

func main() {

	// if err := godotenv.Load("../.env"); err != nil {
	// 	log.Fatal("could not be read env files")
	// }

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8081" // value default
	// }

	// path := os.Getenv("JSON")
	// if path == "" {
	// 	log.Fatal("could not be read path")
	// }

	fmt.Println("Project running in the port : 8080")
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
