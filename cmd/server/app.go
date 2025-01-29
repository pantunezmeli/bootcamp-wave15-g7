package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"
	loaderfile "github.com/pantunezmeli/bootcamp-wave15-g7/internal/loaderFile"
	buyerRepository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/buyer"
	buyerService "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/buyer"
	//warehouse_h "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"
	//warehouse_rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/warehouse_repository"
	//warehouse_sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/warehouse_service"
	//loader "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/warehouse_storage"
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
	ld := loaderfile.NewBuyerJSONFile(a.loaderFilePath)
	// ldw := loader.NewWareHouseJSONFile(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}

	// - repository
	rp := buyerRepository.NewBuyerRepository(db, ld)
	// - service
	sv := buyerService.NewBuyerService(rp)
	// - handler
	hd := handler.NewBuyerHandler(sv)
	//  wh_rp := warehouse_rp.NewWareHouseRepository(db, ldw)
	//  wh_sv := warehouse_sv.NewWareHouseService(wh_rp)
	//  wh_h := warehouse_h.NewWareHouseHandler(wh_sv)

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	rt.Route("/api/v1", func(rt chi.Router) {
		rt.Route("/sellers", func(rt chi.Router) {
			// Agrega tus rutas de sellers aquí
		})

		rt.Route("/warehouses", func(rt chi.Router) {
			// Aquí se corrige el bloque para manejar correctamente las rutas
			// rt.Get("/", wh_h.Get())
			// rt.Get("/{id}", wh_h.GetById())
			// rt.Post("/", wh_h.Create())
			// rt.Patch("/{id}", wh_h.Update())
			// rt.Delete("/{id}", wh_h.Delete())
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
			rt.Get("/", hd.GetAll())
			rt.Get("/{id}", hd.GetBuyerById())
			rt.Post("/", hd.CreateBuyer())
			rt.Patch("/{id}", hd.UpdateBuyer())
		})
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
