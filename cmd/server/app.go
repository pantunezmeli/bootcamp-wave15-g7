package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	sectionStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/section"

	//Buyer
	buyerRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer"
	buyerService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer"
	buyerStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/buyer_storage"

	// Warehouse
	warehouseRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse"
	warehouseService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse"
	warehouseStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/warehouse_storage"
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
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
	}

	return &ServerChi{
		serverAddress: defaultConfig.ServerAddress,
	}
}

type ServerChi struct {
	serverAddress string
}

func (a *ServerChi) Run() (err error) {

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
	warehouseStorage := warehouseStorage.NewWareHouseJSONFile(warehouseFilePath)
	warehouseRepository := warehouseRepository.NewWareHouseRepository(warehouseStorage)
	warehouseService := warehouseService.NewWareHouseService(warehouseRepository)
	warehouseHandler := handler.NewWareHouseHandler(warehouseService)

	// Product
	productStorage := productStorage.NewProductJSONFile(productFilePath)
	productRepository := productRepository.NewProductRepositoryMap(productStorage)
	productService := productService.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	// Sections
	sectionStorage := sectionStorage.NewSectionJSONFile(sectionFilePath)
	sectionRepository := sectionRepository.NewStRepository(sectionStorage)
	sectionService := sectionService.NewSectionService(sectionRepository)
	sectionHandler := handler.NewSectionDefault(sectionService)

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
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
