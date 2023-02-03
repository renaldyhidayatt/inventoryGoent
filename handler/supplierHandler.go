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

type supplierHandler struct {
	supplier services.SupplierService
}

func NewSupplierService(supplier services.SupplierService) *supplierHandler {
	return &supplierHandler{supplier: supplier}
}

func (h *supplierHandler) Results(w http.ResponseWriter, r *http.Request) {
	res, err := h.supplier.Results()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *supplierHandler) Result(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.supplier.Result(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *supplierHandler) Create(w http.ResponseWriter, r *http.Request) {
	var supplierRequest request.SupplierRequest
	err := json.NewDecoder(r.Body).Decode(&supplierRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(supplierRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	res, err := h.supplier.Create(supplierRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *supplierHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	var supplierRequest request.SupplierRequest
	err = json.NewDecoder(r.Body).Decode(&supplierRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(supplierRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.supplier.UpdateById(int(Id), supplierRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil update data", res, http.StatusCreated)
	}
}

func (h *supplierHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}
	res, err := h.supplier.DeleteById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil delete data", res, http.StatusCreated)
	}
}
