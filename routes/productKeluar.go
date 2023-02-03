package routes

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/handler"
	"github.com/renaldyhidayatt/inventorygoent/middlewares"
	"github.com/renaldyhidayatt/inventorygoent/repository"
	"github.com/renaldyhidayatt/inventorygoent/services"
)

func NewProductKeluarRoutes(prefix string, db *ent.Client, router *chi.Mux, context context.Context) {
	repository := repository.NewProductKeluarRepository(db, context)
	service := services.NewProductKeluarService(repository)
	handler := handler.NewProductKeluarHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/", handler.FindAll)
		r.Get("/{id:[0-9]+}", handler.FindById)
		r.Post("/", handler.Create)
		r.Put("/{id:[0-9]+}", handler.UpdateById)
		r.Delete("/{id:[0-9]+}", handler.DeleteById)
	})
}
