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

func (h *Handler) InitCategory(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)

		r.Get("/", h.GetCategories)
		r.Get("/{id}", h.GetCategory)
		r.Post("/create", h.CreateCategory)
		r.Put("/update/{id}", h.UpdateCategory)
		r.Delete("/delete/{id}", h.DeleteCategory)
	})
}

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.services.Category.GetAll()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
	}

	response.ResponseMessage(w, "get categories success", categories, http.StatusOK)
}

func (h *Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.services.Category.GetById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.ResponseMessage(w, "get category success", category, http.StatusOK)
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var requests request.CreateCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	category, err := h.services.Category.Create(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "create category success", category, http.StatusOK)
}

func (h *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var requests request.UpdateCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	category, err := h.services.Category.Update(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "update category success", category, http.StatusOK)
}

func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	err = h.services.Category.Delete(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	response.ResponseMessage(w, "delete category success", nil, http.StatusOK)
}
