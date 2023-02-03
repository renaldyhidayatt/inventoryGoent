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

type categoryHandler struct {
	category services.CategoryService
}

func NewCategoryHandler(category services.CategoryService) *categoryHandler {
	return &categoryHandler{category: category}
}

func (h *categoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.category.Results()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *categoryHandler) FindById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.category.Result(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *categoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var categoryRequest request.CategoryRequest
	err := json.NewDecoder(r.Body).Decode(&categoryRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(categoryRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	res, err := h.category.Create(categoryRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *categoryHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	var categoryRequest request.CategoryRequest
	err = json.NewDecoder(r.Body).Decode(&categoryRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(categoryRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.category.UpdateById(int(Id), categoryRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil update data", res, http.StatusCreated)
	}
}

func (h *categoryHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}
	res, err := h.category.DeleteById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil delete data", res, http.StatusCreated)
	}
}
