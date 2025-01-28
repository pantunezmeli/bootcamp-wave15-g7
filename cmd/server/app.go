package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	warehouse_h "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"
	loader "github.com/pantunezmeli/bootcamp-wave15-g7/internal/loader/warehouse_loader"
	warehouse_rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse_repository"
	warehouse_sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse_service"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
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
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the server
func (a *ServerChi) Run() (err error) {
	// dependencies
	// - loader
	ld := loader.NewWareHouseJSONFile(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}

	wh_rp := warehouse_rp.NewWareHouseRepository(db, ld)
	wh_sv := warehouse_sv.NewWareHouseService(wh_rp)
	wh_h := warehouse_h.NewWareHouseHandler(wh_sv)

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	rt.Route("/sellers", func(rt chi.Router) {

	})

	rt.Route("/warehouses", func(rt chi.Router) {
		rt.Get("/", wh_h.GetAll())

	})

	rt.Route("/sections", func(rt chi.Router) {

	})

	rt.Route("/products", func(rt chi.Router) {

	})

	rt.Route("/employees", func(rt chi.Router) {

	})

	rt.Route("/buyers", func(rt chi.Router) {

	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
