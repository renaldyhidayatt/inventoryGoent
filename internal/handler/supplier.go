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

func (h *Handler) InitSupplier(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)
		r.Get("/", h.GetSuppliers)
		r.Get("/{id}", h.GetSupplier)
		r.Post("/create", h.CreateSupplier)
		r.Put("/update/{id}", h.UpdateSupplier)
		r.Delete("/delete/{id}", h.DeleteSupplier)
	})
}

func (h *Handler) GetSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := h.services.Supplier.GetAll()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return

	}

	response.ResponseMessage(w, "get suppliers success", suppliers, http.StatusOK)

}

func (h *Handler) GetSupplier(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	supplier, err := h.services.Supplier.GetByID(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "get supplier success", supplier, http.StatusOK)
}

func (h *Handler) CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var requests request.CreateSupplierRequest

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	supplier, err := h.services.Supplier.Create(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "create supplier success", supplier, http.StatusOK)
}

func (h *Handler) UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	var requests request.UpdateSupplierRequest

	requests.ID = int(Id)

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	supplier, err := h.services.Supplier.Update(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "update supplier success", supplier, http.StatusOK)
}

func (h *Handler) DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	err = h.services.Supplier.Delete(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "delete supplier success", nil, http.StatusOK)
}
