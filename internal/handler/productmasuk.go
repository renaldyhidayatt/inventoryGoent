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

func (h *Handler) InitProductMasuk(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)
		r.Get("/", h.GetProductMasuks)
		r.Get("/{id}", h.GetProductMasuk)
		r.Post("/create", h.CreateProductMasuk)
		r.Put("/update/{id}", h.UpdateProductMasuk)
		r.Delete("/delete/{id}", h.DeleteProductMasuk)
	})
}

func (h *Handler) GetProductMasuks(w http.ResponseWriter, r *http.Request) {
	productMasuks, err := h.services.ProductMasuk.GetAll()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "get product masuks success", productMasuks, http.StatusOK)
}

func (h *Handler) GetProductMasuk(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	productMasuk, err := h.services.ProductMasuk.GetById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "get product masuk success", productMasuk, http.StatusOK)
}

func (h *Handler) CreateProductMasuk(w http.ResponseWriter, r *http.Request) {
	var requests request.CreateProductMasukRequest

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	productMasuk, err := h.services.ProductMasuk.Create(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "create product masuk success", productMasuk, http.StatusOK)
}

func (h *Handler) UpdateProductMasuk(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	var requests request.UpdateProductMasukRequest

	requests.ID = int(Id)

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	productMasuk, err := h.services.ProductMasuk.Update(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "update product masuk success", productMasuk, http.StatusOK)
}

func (h *Handler) DeleteProductMasuk(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	err = h.services.ProductMasuk.Delete(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "delete product masuk success", nil, http.StatusOK)

}
