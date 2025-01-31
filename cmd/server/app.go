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

type ConfigServerChi struct {
	ServerAddress string
	Path          string
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.Path != "" {
			defaultConfig.Path = cfg.Path

		}
	}

	return &ServerChi{
		serverAddress: defaultConfig.ServerAddress,
		path:          defaultConfig.Path,
	}
}

type ServerChi struct {
	serverAddress string
	path          string
}

func (a *ServerChi) Run() (err error) {

	buyerSt := buyerstorage.NewBuyerJSONFile(a.path)
	warehouseSt := loader.NewWareHouseJSONFile(a.path)
	dbwarehouse, _ := warehouseSt.Load()

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
			rt.Get("/{id}", by_hd.GetById())
			rt.Post("/", by_hd.Create())
			rt.Patch("/{id}", by_hd.Update())
			rt.Delete("/{id}", by_hd.Delete())
		})
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
