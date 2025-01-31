package server

import (
	"net/http"

	product_ld "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/product_storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	erp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/employee"
	rep "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/section"
	SellerRepo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/seller"
	esv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/employee"
	sec "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/section"
	SellerService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/seller"
	employeeStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/employee_storage"
	sectionstorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/section"
	buyerStorageorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/buyer_storage"
	warehouseStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/warehouse_storage"
	sellerStorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/seller_storage"

	handler "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"
	product_hd "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"
	buyerRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer"
	product_rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	warehouse_rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse_repository"
	buyerService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer"
	product_sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product"
	warehouse_sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse_service"
)

const (
	PATH_PRODUCT_JSON_FILE = "../docs/db/product_data.json"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	ServerAddress           string
	BuyerLoaderFilePath     string
	WarehouseLoaderFilePath string
	EmployeeLoaderFilePath  string
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
		if cfg.EmployeeLoaderFilePath != "" {
			defaultConfig.EmployeeLoaderFilePath = cfg.EmployeeLoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:     defaultConfig.ServerAddress,
		buyerFilePath:     defaultConfig.BuyerLoaderFilePath,
		warehouseFilePath: defaultConfig.WarehouseLoaderFilePath,
		employeeFilPath:   defaultConfig.EmployeeLoaderFilePath,

	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	serverAddress     string
	buyerFilePath     string
	warehouseFilePath string
	employeeFilPath   string

}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {


	// - loader
	sectionSt := sectionstorage.NewSectionJSONFile("../docs/db/section_data.json")

	// - repository
	st_rp := rep.NewStRepository(sectionSt)
	// - service
	st_sv := sec.NewSectionService(st_rp)
	// - handler
	st_hd := handler.NewSectionDefault(st_sv)
	warehouseSt := warehouseStorage.NewWareHouseJSONFile(a.warehouseFilePath)
	productSt := product_ld.NewProductJSONFile(PATH_PRODUCT_JSON_FILE)


	// Sellers
	ld := sellerStorage.NewSellerJSONFile("../docs/db/seller_data.json")
	sellerRepository := SellerRepo.NewSellerStorage(*ld)
	sellerService := SellerService.NewSellerDefault(sellerRepository)
	sellerHandler := handler.NewSellerDefault(sellerService)


	// Employees
	employeeStorage := employeeStorage.NewEmployeeJSONFile(a.employeeFilPath)
	employeeRepository := erp.NewEmployeeMap(*employeeStorage)
	employeeService := esv.NewDefaultService(employeeRepository)
	employeeHandler := handler.NewDefaultHandler(employeeService)

	// Buyers
	buyerStorage := buyerStorageorage.NewBuyerJSONFile(a.buyerFilePath)
	by_rp := buyerRepository.NewBuyerRepository(buyerStorage)
	by_sv := buyerService.NewBuyerService(by_rp)
	by_hd := handler.NewBuyerHandler(by_sv)

	// Warehouses
	wh_rp := warehouse_rp.NewWareHouseRepository(warehouseSt)
	wh_sv := warehouse_sv.NewWareHouseService(wh_rp)
	wh_h := handler.NewWareHouseHandler(wh_sv)

	// Product
	rpProduct := product_rp.NewProductRepositoryMap(productSt)
	svProduct := product_sv.NewProductService(rpProduct)
	hdProduct := product_hd.NewProductHandler(svProduct)

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints

	rt.Route("/api/v1", func(r chi.Router) {
		r.Route("/sellers", func(r chi.Router) {
			r.Get("/", sellerHandler.GetAll())
			r.Get("/{id}", sellerHandler.GetById())
			r.Post("/", sellerHandler.Create())
			r.Delete("/{id}", sellerHandler.Delete())
			r.Patch("/{id}", sellerHandler.Update())

		})

		r.Route("/warehouses", func(rt chi.Router) {
			rt.Get("/", wh_h.Get())
			rt.Get("/{id}", wh_h.GetById())
			rt.Post("/", wh_h.Create())
			rt.Patch("/{id}", wh_h.Update())
			rt.Delete("/{id}", wh_h.Delete())
		})


		r.Route("/sections", func(rt chi.Router) {
			rt.Get("/", st_hd.Get())
			rt.Get("/{id}", st_hd.GetById())
			rt.Post("/", st_hd.Create())
			rt.Patch("/{id}", st_hd.Update())
			rt.Delete("/{id}", st_hd.Delete())		
		})

		r.Route("/products", func(rt chi.Router) {
			rt.Get("/", hdProduct.GetAll())
			rt.Get("/{id}", hdProduct.GetById())
			rt.Post("/", hdProduct.Create())
			rt.Patch("/{id}", hdProduct.Update())
			rt.Delete("/{id}", hdProduct.Delete())
		})

		r.Route("/employees", func(rt chi.Router) {
			rt.Get("/", employeeHandler.GetAll())
			rt.Get("/{id}", employeeHandler.GetById())
			rt.Post("/", employeeHandler.Add())
			rt.Patch("/{id}", employeeHandler.Update())
			rt.Delete("/{id}", employeeHandler.DeleteById())
		})

		r.Route("/buyers", func(rt chi.Router) {
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
