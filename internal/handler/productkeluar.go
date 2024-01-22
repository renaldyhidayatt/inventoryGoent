package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/internal/middlewares"
)

func (h *Handler) InitProductKeluar(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)
		r.Get("/", h.GetProductKeluars)
		r.Get("/{id}", h.GetProductKeluar)
		r.Post("/create", h.CreateProductKeluar)
		r.Put("/update/{id}", h.UpdateProductKeluar)
		r.Delete("/delete/{id}", h.DeleteProductKeluar)
	})
}

func (h *Handler) GetProductKeluars(w http.ResponseWriter, r *http.Request) {
	productKeluars, err := h.services.ProductKeluar.GetAll()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "get product keluars success", productKeluars, http.StatusOK)
}

func (h *Handler) GetProductKeluar(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	productKeluar, err := h.services.ProductKeluar.GetById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "get product keluar success", productKeluar, http.StatusOK)
}

func (h *Handler) CreateProductKeluar(w http.ResponseWriter, r *http.Request) {
	var requests request.CreateProductKeluarRequest

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	productKeluar, err := h.services.ProductKeluar.Create(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "create product keluar success", productKeluar, http.StatusOK)
}

func (h *Handler) UpdateProductKeluar(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	var requests request.UpdateProductKeluarRequest

	requests.ID = int(Id)

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	productKeluar, err := h.services.ProductKeluar.Update(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "update product keluar success", productKeluar, http.StatusOK)
}

func (h *Handler) DeleteProductKeluar(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	err = h.services.ProductKeluar.Delete(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "delete product keluar success", nil, http.StatusOK)
}
