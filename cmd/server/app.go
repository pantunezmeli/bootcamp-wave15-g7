package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	ehd "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler/employee"
	erp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/employee"
	esv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/employee"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage"

	buyerstorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/buyer_storage"
	warehouseStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/warehouse_storage"

	handler "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"
	product_hd "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"
	product_ld "github.com/pantunezmeli/bootcamp-wave15-g7/internal/loader/product"
	buyerRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer"
	product_rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	warehouse_rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse_repository"
	buyerService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer"
	product_sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product"
	warehouse_sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse_service"
)

const (
	PATH_PRODUCT_JSON_FILE = "docs/db/product_data.json"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	ServerAddress           string
	BuyerLoaderFilePath     string
	WarehouseLoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress:           ":8080",
		BuyerLoaderFilePath:     "../docs/db/buyer_data.json",
		WarehouseLoaderFilePath: "../docs/db/warehouse_data.json",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.BuyerLoaderFilePath != "" {
			defaultConfig.BuyerLoaderFilePath = cfg.BuyerLoaderFilePath
		}
		if cfg.WarehouseLoaderFilePath != "" {
			defaultConfig.WarehouseLoaderFilePath = cfg.WarehouseLoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:     defaultConfig.ServerAddress,
		buyerFilePath:     defaultConfig.BuyerLoaderFilePath,
		warehouseFilePath: defaultConfig.WarehouseLoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	serverAddress     string
	buyerFilePath     string
	warehouseFilePath string
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {

	// - loader
	est := storage.NewEmployeeJSONFile(a.loaderFilePath)
	buyerSt := buyerstorage.NewBuyerJSONFile(a.buyerFilePath)
	warehouseSt := warehouseStorage.NewWareHouseJSONFile(a.warehouseFilePath)
	ldProduct := product_ld.NewProductJSONFile(PATH_PRODUCT_JSON_FILE)

	// Employee
	employeeRepository := erp.NewEmployeeMap(*est)
	employeeService := esv.NewDefaultService(employeeRepository)
	employeeHandler := ehd.NewDefaultHandler(employeeService)

	// Buyer
	by_rp := buyerRepository.NewBuyerRepository(buyerSt)
	by_sv := buyerService.NewBuyerService(by_rp)
	by_hd := handler.NewBuyerHandler(by_sv)

	// Warehouse
	wh_rp := warehouse_rp.NewWareHouseRepository(warehouseSt)
	wh_sv := warehouse_sv.NewWareHouseService(wh_rp)
	wh_h := handler.NewWareHouseHandler(wh_sv)

	// Product
	rpProduct := product_rp.NewProductRepositoryMap(ldProduct)
	svProduct := product_sv.NewProductService(rpProduct)
	hdProduct := product_hd.NewProductHandler(svProduct)

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	rt.Route("/api/v1", func(rt chi.Router) {
		rt.Route("/sellers", func(rt chi.Router) {
		})

		rt.Route("/warehouses", func(rt chi.Router) {
			rt.Get("/", wh_h.Get())
			rt.Get("/{id}", wh_h.GetById())
			rt.Post("/", wh_h.Create())
			rt.Patch("/{id}", wh_h.Update())
			rt.Delete("/{id}", wh_h.Delete())
		})

		rt.Route("/sections", func(rt chi.Router) {
			// Agrega tus rutas de sections aquí
		})

		rt.Route("/products", func(r chi.Router) {
			r.Get("/", hdProduct.GetAll())
			r.Get("/{id}", hdProduct.GetById())
			r.Post("/", hdProduct.Create())
			r.Patch("/{id}", hdProduct.Update())
			r.Delete("/{id}", hdProduct.Delete())
		})

		rt.Route("/employees", func(rt chi.Router) {
			rt.Get("/", employeeHandler.GetAll())
			rt.Get("/{id}", employeeHandler.GetById())
			rt.Post("/", employeeHandler.Add())
			rt.Patch("/{id}", employeeHandler.Update())
			rt.Delete("/{id}", employeeHandler.DeleteById())
		})

		rt.Route("/buyers", func(rt chi.Router) {
			rt.Get("/", by_hd.GetAll())
			rt.Get("/{id}", by_hd.GetBuyerById())
			rt.Post("/", by_hd.CreateBuyer())
			rt.Patch("/{id}", by_hd.UpdateBuyer())
			rt.Delete("/{id}", by_hd.DeleteBuyer())
		})
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
