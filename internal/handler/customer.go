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

func (h *Handler) InitCustomer(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)
		r.Get("/", h.GetCustomers)
		r.Get("/{id}", h.GetCustomer)
		r.Post("/create", h.CreateCustomer)
		r.Put("/update/{id}", h.UpdateCustomer)
		r.Delete("/delete/{id}", h.DeleteCustomer)
	})
}

func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.services.Customer.GetAll()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.ResponseMessage(w, "get customers success", customers, http.StatusOK)
}

func (h *Handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	customer, err := h.services.Customer.GetById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "get customer success", customer, http.StatusOK)
}

func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var requests request.CreateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	customer, err := h.services.Customer.Create(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "create customer success", customer, http.StatusOK)
}

func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var requests request.UpdateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	customer, err := h.services.Customer.Update(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "update customer success", customer, http.StatusOK)
}

func (h *Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	err = h.services.Customer.Delete(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "delete customer success", nil, http.StatusOK)
}
