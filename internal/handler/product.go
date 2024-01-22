package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/internal/middlewares"
	"github.com/renaldyhidayatt/inventorygoent/internal/upload"
)

func (h *Handler) InitProduct(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)
		r.Get("/", h.GetProducts)
		r.Get("/{id}", h.GetProduct)
		r.Post("/create", h.CreateProduct)
		r.Put("/update/{id}", h.UpdateProduct)
		r.Delete("/delete/{id}", h.DeleteProduct)
	})
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.services.Product.GetAll()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.ResponseMessage(w, "get products success", products, http.StatusOK)
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	product, err := h.services.Product.GetById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "get product success", product, http.StatusOK)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var requests request.CreateProductRequest

	fileupload := upload.NewLocationUpload("product", "")

	requests.Name = r.FormValue("name")
	requests.Image = fileupload.FileUpload(w, r)
	requests.Qty = r.FormValue("qty")
	categoryID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
	}
	requests.CategoryID = categoryID

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	product, err := h.services.Product.Create(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "create product success", product, http.StatusOK)

}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var requests request.UpdateProductRequest

	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	fileupload := upload.NewLocationUpload("product", "")

	requests.ID = int(Id)
	requests.Name = r.FormValue("name")
	requests.Image = fileupload.FileUpload(w, r)
	requests.Qty = r.FormValue("qty")
	categoryID, err := strconv.Atoi(r.FormValue("category_id"))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")

		return
	}

	requests.CategoryID = categoryID

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	product, err := h.services.Product.Update(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "update product success", product, http.StatusOK)

}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	err = h.services.Product.Delete(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "delete product success", nil, http.StatusOK)
}
