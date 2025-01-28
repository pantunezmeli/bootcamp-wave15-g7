package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/loader"
	SellerRepo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/seller"
	SellerService"github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/seller"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"
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
	ld := loader.NewSellerJSONFile("../docs/db/seller_data.json")


	// - repository
	sellerRepository := SellerRepo.NewSellerStorage(*ld)



	// - service
	sellerService := SellerService.NewSellerDefault(sellerRepository)


	// - handler
	sellerHandler := handler.NewSellerDefault(sellerService)


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
		})
	
		r.Route("/warehouses", func(r chi.Router) {
		})
	
		r.Route("/sections", func(r chi.Router) {
		})
	
		r.Route("/products", func(r chi.Router) {
		})
	
		r.Route("/employees", func(r chi.Router) {
		})
	
		r.Route("/buyers", func(r chi.Router) {
		})
	})
	

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
