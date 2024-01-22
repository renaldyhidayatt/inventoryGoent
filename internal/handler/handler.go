package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/inventorygoent/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *chi.Mux {
	router := chi.NewRouter()

	h.initApi(router)

	fileServer := http.FileServer(http.Dir("./static"))

	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}

func (h *Handler) initApi(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	h.InitAuth("/auth", r)
	h.InitCategory("/category", r)
	h.InitCustomer("/customer", r)
	h.InitProduct("/product", r)
	h.InitProductMasuk("/productmasuk", r)
	h.InitProductKeluar("/productkeluar", r)
	h.InitSupplier("/supplier", r)
}
