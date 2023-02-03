package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/dto/response"
	"github.com/renaldyhidayatt/inventorygoent/services"
)

type customerHandler struct {
	customer services.CustomerService
}

func NewCustomerHandler(customer services.CustomerService) *customerHandler {
	return &customerHandler{customer: customer}
}

func (h *customerHandler) Results(w http.ResponseWriter, r *http.Request) {
	res, err := h.customer.Results()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *customerHandler) Result(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.customer.Result(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *customerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var customerRequest request.CustomerRequest
	err := json.NewDecoder(r.Body).Decode(&customerRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(customerRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	res, err := h.customer.Create(customerRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *customerHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	var customerRequest request.CustomerRequest
	err = json.NewDecoder(r.Body).Decode(&customerRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(customerRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.customer.UpdateById(int(Id), customerRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil update data", res, http.StatusCreated)
	}
}

func (h *customerHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}
	res, err := h.customer.DeleteById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil delete data", res, http.StatusCreated)
	}
}
