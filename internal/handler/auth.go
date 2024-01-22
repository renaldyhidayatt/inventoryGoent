package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/request"
	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
)

func (h *Handler) InitAuth(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var requests request.AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	user, err := h.services.Auth.Login(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.ResponseToken(w, "login success", user, http.StatusOK)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var requests request.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if err := requests.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "invalid request")
		return
	}

	user, err := h.services.Auth.Register(requests)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.ResponseMessage(w, "register success", user, http.StatusOK)
}
