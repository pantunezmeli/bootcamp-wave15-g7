package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	buyerstorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/buyer_storage"

	handler "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"
	buyerRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer"
	warehouse_rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse_repository"
	buyerService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer"
	warehouse_sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse_service"
	loader "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/warehouse_storage"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	ServerAddress  string
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath

		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	serverAddress  string
	loaderFilePath string
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// dependencies
	// - loader
	buyerSt := buyerstorage.NewBuyerJSONFile(a.loaderFilePath)
	warehouseSt := loader.NewWareHouseJSONFile(a.loaderFilePath)
	dbwarehouse, _ := warehouseSt.Load()

	// if err2 != nil {
	// 	return
	// }

	by_rp := buyerRepository.NewBuyerRepository(buyerSt)
	by_sv := buyerService.NewBuyerService(by_rp)
	by_hd := handler.NewBuyerHandler(by_sv)

	wh_rp := warehouse_rp.NewWareHouseRepository(dbwarehouse, warehouseSt)
	wh_sv := warehouse_sv.NewWareHouseService(wh_rp)
	wh_h := handler.NewWareHouseHandler(wh_sv)

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

		rt.Route("/products", func(rt chi.Router) {
			// Agrega tus rutas de products aquí
		})

		rt.Route("/employees", func(rt chi.Router) {
			// Agrega tus rutas de employees aquí
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
