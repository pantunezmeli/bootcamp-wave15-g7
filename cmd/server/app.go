package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/config"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/db"
	handler "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"

	// Products
	productRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	productService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product"
	//productStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/product_storage"

	//ProductRecords
	productRecordsRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/productrecords"
	productRecordsService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product_records"

	// Sellers
	sellerRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/seller"
	sellerService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/seller"

	//Localities
	localityRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/locality"
	localityService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/locality"


	// Employees
	employeeRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/employee"
	employeeService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/employee"
	employeeStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/employee_storage"

	// Sections
	sectionRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/section"
	sectionService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/section"
	sectionStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/section"

	// Buyer
	buyerRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer"
	buyerService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer"

	// Purchase
	purchaseRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/purchase"
	purchaseService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/purchase"

	// Warehouse
	warehouseRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse"
	warehouseService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse"

	// Carrier
	carrierRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/carrier"
	carrierService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/carrier"
)

const (
	productFilePath   = "../docs/db/json/product_data.json"
	buyerFilePath     = "../docs/db/json/buyer_data.json"
	warehouseFilePath = "../docs/db/json/warehouse_data.json"
	employeeFilePath  = "../docs/db/json/employee_data.json"
	sellerFilePath    = "../docs/db/json/seller_data.json"
	sectionFilePath   = "../docs/db/json/section_data.json"
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
	sellerRepository := sellerRepository.NewSellerMySql(dbConn)
	sellerService := sellerService.NewSellerDefault(sellerRepository)
	sellerHandler := handler.NewSellerDefault(sellerService)

	// Localities
	localityRepository := localityRepository.NewLocalityMySql(dbConn)
	localityService := localityService.NewLocalityDefault(localityRepository)
	localityHandler := handler.NewLocalityDefault(localityService)

	// Employees
	employeeStorage := employeeStorage.NewEmployeeJSONFile(employeeFilePath)
	employeeRepository := employeeRepository.NewEmployeeMap(*employeeStorage)
	employeeService := employeeService.NewDefaultService(employeeRepository)
	employeeHandler := handler.NewDefaultHandler(employeeService)

	// Buyers
	//buyerStorage := buyerStorage.NewBuyerJSONFile(buyerFilePath)
	buyerRepository := buyerRepository.NewBuyerRepository(dbConn)
	buyerService := buyerService.NewBuyerService(buyerRepository)
	buyerHandler := handler.NewBuyerHandler(buyerService)

	// Purchases
	purchaseRepository := purchaseRepository.NewBuyerRepository(dbConn)
	purchaseService := purchaseService.NewBuyerService(purchaseRepository)
	purchaseHandler := handler.NewPurchaseHandler(purchaseService)

	// Warehouses
	warehouseRepository := warehouseRepository.NewWareHouseRepository(dbConn)
	warehouseService := warehouseService.NewWareHouseService(warehouseRepository)
	warehouseHandler := handler.NewWareHouseHandler(warehouseService)

	// Product
	//productStorage := productStorage.NewProductJSONFile(productFilePath)
	productRepository := productRepository.NewProductRepositoryMysql(dbConn)
	productService := productService.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	//ProductRecords
	productRecordsRepository := productRecordsRepository.NewProductRecordsRepository(dbConn)
	productRecordsService := productRecordsService.NewProductRecordsService(productRecordsRepository)
	productRecordsHandler := handler.NewHandlerProductRecords(productRecordsService)

	// Sections
	sectionStorage := sectionStorage.NewSectionJSONFile(sectionFilePath)
	sectionRepository := sectionRepository.NewStRepository(sectionStorage)
	sectionService := sectionService.NewSectionService(sectionRepository)
	sectionHandler := handler.NewSectionDefault(sectionService)

	// Carriers
	carrierRepository := carrierRepository.NewCarrierRepository(dbConn)
	carrierService := carrierService.NewCarrierService(carrierRepository)
	carrierHandler := handler.NewCarrierHandler(carrierService)

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

		r.Route("/localities", func(r chi.Router) {
			r.Get("/reportSellers", localityHandler.GetReportSellers())
			r.Post("/", localityHandler.Create())
			r.Get("/{id}", localityHandler.GetById())
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
			rt.Get("/reportRecords", productRecordsHandler.GetRecords())
		})

		r.Route("/productRecords", func(rt chi.Router) {
			rt.Post("/", productRecordsHandler.Create())
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

		r.Route("/reportPurchaseOrders", func(rt chi.Router) {
			rt.Get("/", purchaseHandler.Get())
			rt.Get("/{id}", purchaseHandler.GetById())
			rt.Post("/", purchaseHandler.Create())

		})

		r.Route("/carriers", func(rt chi.Router) {
			rt.Post("/", carrierHandler.Create())
			rt.Get("/reportCarries", carrierHandler.GetCarriesAmount())
		})
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
