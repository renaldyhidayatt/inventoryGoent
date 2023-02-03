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

type productKeluarHandler struct {
	productKeluar services.ProductKeluarService
}

func NewProductKeluarHandler(productkeluar services.ProductKeluarService) *productKeluarHandler {
	return &productKeluarHandler{productKeluar: productkeluar}
}

func (h *productKeluarHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.productKeluar.Results()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *productKeluarHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.productKeluar.Result(int(id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *productKeluarHandler) Create(w http.ResponseWriter, r *http.Request) {
	var productKeluarRequest request.ProductKeluarRequest
	err := json.NewDecoder(r.Body).Decode(&productKeluarRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(productKeluarRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	res, err := h.productKeluar.Create(productKeluarRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *productKeluarHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	var productKeluarRequest request.ProductKeluarRequest
	err = json.NewDecoder(r.Body).Decode(&productKeluarRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(productKeluarRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.productKeluar.UpdateById(int(Id), productKeluarRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil update data", res, http.StatusCreated)
	}
}

func (h *productKeluarHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}
	res, err := h.productKeluar.DeleteById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil delete data", res, http.StatusCreated)
	}
}
