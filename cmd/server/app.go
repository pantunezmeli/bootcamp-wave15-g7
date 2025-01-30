package server

import (
	product_hd "github.com/pantunezmeli/bootcamp-wave15-g7/internal/handler"
	product_ld "github.com/pantunezmeli/bootcamp-wave15-g7/internal/loader/product"
	product_rp "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/product"
	product_sv "github.com/pantunezmeli/bootcamp-wave15-g7/internal/service/product"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	PATH_PRODUCT_JSON_FILE = "docs/db/product_data.json"
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
	// ld := loader.NewVehicleJSONFile(a.loaderFilePath)
	// db, err := ld.Load()
	// if err != nil {
	// 	return
	// }

	ldProduct := product_ld.NewProductJSONFile(PATH_PRODUCT_JSON_FILE)
	rpProduct := product_rp.NewProductRepositoryMap(ldProduct)
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
		})

		r.Route("/warehouses", func(r chi.Router) {
		})

		r.Route("/sections", func(r chi.Router) {
		})

		r.Route("/products", func(r chi.Router) {
			r.Get("/", hdProduct.GetAll())
			r.Get("/{id}", hdProduct.GetById())
			r.Post("/", hdProduct.Create())
			r.Patch("/{id}", hdProduct.Update())
			r.Delete("/{id}", hdProduct.Delete())
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
