package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/config"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/db"
	handler "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"

	//Products
	productRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	productService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product"
	productStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/product_storage"

	//Sellers
	sellerRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/seller"
	sellerService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/seller"
	sellerStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/seller_storage"

	// Employees
	employeeRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/employee"
	employeeService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/employee"
	employeeStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/employee_storage"

	//Sections
	sectionRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/section"
	sectionService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/section"

	//Buyer
	buyerRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer"
	buyerService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer"
	buyerStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/buyer_storage"

	// Warehouse
	warehouseRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse"
	warehouseService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse"

	//Product Batches
	productBatchRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product_batch"
	productBatchService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product_batch"
)

const (
	productFilePath   = "../docs/db/product_data.json"
	buyerFilePath     = "../docs/db/buyer_data.json"
	warehouseFilePath = "../docs/db/warehouse_data.json"
	employeeFilePath  = "../docs/db/employee_data.json"
	sellerFilePath    = "../docs/db/seller_data.json"
	sectionFilePath   = "../docs/db/section_data.json"
)

type ConfigServerChi struct {
	ServerAddress string
	DB            *sql.DB
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.DB != nil {
			defaultConfig.DB = cfg.DB
		}
	}

	return &ServerChi{
		serverAddress: defaultConfig.ServerAddress,
		db:            defaultConfig.DB,
	}
}

type ServerChi struct {
	serverAddress string
	db            *sql.DB
}

func (a *ServerChi) Run() (err error) {

	// Set DB connection
	cfg := config.LoadConfig()

	// Connection
	dbConn := db.CreateConnectionToDB(cfg)
	defer dbConn.Close()

	// Sellers
	sellerStorage := sellerStorage.NewSellerJSONFile(sellerFilePath)
	sellerRepository := sellerRepository.NewSellerStorage(*sellerStorage)
	sellerService := sellerService.NewSellerDefault(sellerRepository)
	sellerHandler := handler.NewSellerDefault(sellerService)

	// Employees
	employeeStorage := employeeStorage.NewEmployeeJSONFile(employeeFilePath)
	employeeRepository := employeeRepository.NewEmployeeMap(*employeeStorage)
	employeeService := employeeService.NewDefaultService(employeeRepository)
	employeeHandler := handler.NewDefaultHandler(employeeService)

	// Buyers
	buyerStorage := buyerStorage.NewBuyerJSONFile(buyerFilePath)
	buyerRepository := buyerRepository.NewBuyerRepository(buyerStorage)
	buyerService := buyerService.NewBuyerService(buyerRepository)
	buyerHandler := handler.NewBuyerHandler(buyerService)

	// Warehouses
	warehouseRepository := warehouseRepository.NewWareHouseRepository(dbConn)
	warehouseService := warehouseService.NewWareHouseService(warehouseRepository)
	warehouseHandler := handler.NewWareHouseHandler(warehouseService)

	// Product
	productStorage := productStorage.NewProductJSONFile(productFilePath)
	productRepository := productRepository.NewProductRepositoryMap(productStorage)
	productService := productService.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	// Sections
	sectionRepository := sectionRepository.NewSectionRepository(dbConn)
	sectionService := sectionService.NewSectionService(sectionRepository)
	sectionHandler := handler.NewSectionDefault(sectionService)

	// Product Batches
	productBatchRepository := productBatchRepository.NewProductBatchRepository(dbConn)
	productBatchService := productBatchService.NewProductBatchService(productBatchRepository)
	productBatchHandler := handler.NewProductBatchHandler(productBatchService)

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints

	rt.Route("/api/v1", func(r chi.Router) {
		r.Route("/sellers", func(r chi.Router) {
			r.Get("/", sellerHandler.Get())
			r.Get("/{id}", sellerHandler.GetById())
			r.Post("/", sellerHandler.Create())
			r.Delete("/{id}", sellerHandler.Delete())
			r.Patch("/{id}", sellerHandler.Update())

		})

		r.Route("/warehouses", func(rt chi.Router) {
			rt.Get("/", warehouseHandler.Get())
			rt.Get("/{id}", warehouseHandler.GetById())
			rt.Post("/", warehouseHandler.Create())
			rt.Patch("/{id}", warehouseHandler.Update())
			rt.Delete("/{id}", warehouseHandler.Delete())
		})

		r.Route("/sections", func(rt chi.Router) {
			rt.Get("/", sectionHandler.Get())
			rt.Get("/{id}", sectionHandler.GetById())
			rt.Post("/", sectionHandler.Create())
			rt.Patch("/{id}", sectionHandler.Update())
			rt.Delete("/{id}", sectionHandler.Delete())
		})

		r.Route("/products", func(rt chi.Router) {
			rt.Get("/", productHandler.Get())
			rt.Get("/{id}", productHandler.GetById())
			rt.Post("/", productHandler.Create())
			rt.Patch("/{id}", productHandler.Update())
			rt.Delete("/{id}", productHandler.Delete())
		})

		r.Route("/employees", func(rt chi.Router) {
			rt.Get("/", employeeHandler.Get())
			rt.Get("/{id}", employeeHandler.GetById())
			rt.Post("/", employeeHandler.Create())
			rt.Patch("/{id}", employeeHandler.Update())
			rt.Delete("/{id}", employeeHandler.Delete())
		})

		r.Route("/buyers", func(rt chi.Router) {
			rt.Get("/", buyerHandler.Get())
			rt.Get("/{id}", buyerHandler.GetById())
			rt.Post("/", buyerHandler.Create())
			rt.Patch("/{id}", buyerHandler.Update())
			rt.Delete("/{id}", buyerHandler.Delete())
		})

		r.Route("/productBatches", func(rt chi.Router) {
			rt.Post("/", productBatchHandler.Create())
		})
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
