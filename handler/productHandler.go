package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/dto/response"
	"github.com/renaldyhidayatt/inventorygoent/services"
	"github.com/renaldyhidayatt/inventorygoent/utils"
)

type productHandler struct {
	product services.ProductService
}

func NewProductHandler(product services.ProductService) *productHandler {
	return &productHandler{product: product}
}

func (h *productHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.product.Results()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *productHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.product.Result(int(id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *productHandler) Create(w http.ResponseWriter, r *http.Request) {
	var productRequest request.ProductRequest

	fileupload := utils.NewLocationUpload("product", "")

	productRequest.Name = r.FormValue("name")
	productRequest.Image = fileupload.FileUpload(w, r)
	productRequest.Qty = r.FormValue("qty")
	categoryID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}
	productRequest.CategoryID = categoryID

	validate := validator.New()
	err = validate.Struct(productRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.product.Create(productRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *productHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	var productRequest request.ProductRequest
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.product.Result(int(id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	fileupload := utils.NewLocationUpload("product", res.Image)

	productRequest.Name = r.FormValue("name")
	productRequest.Image = fileupload.FileUpload(w, r)
	productRequest.Qty = r.FormValue("qty")
	categoryID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}
	productRequest.CategoryID = categoryID

	resupdate, err := h.product.UpdateById(int(id), productRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	} else {
		response.ResponseMessage(w, "Berhasil update data", resupdate, http.StatusCreated)
	}
}

func (h *productHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.product.DeleteById(int(id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	} else {
		response.ResponseMessage(w, "Berhasil delete data", res, http.StatusCreated)
	}
}
