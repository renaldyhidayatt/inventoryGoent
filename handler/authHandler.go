package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/renaldyhidayatt/inventorygoent/dto/request"
	"github.com/renaldyhidayatt/inventorygoent/dto/response"
	"github.com/renaldyhidayatt/inventorygoent/security"
	"github.com/renaldyhidayatt/inventorygoent/services"
)

type authHandler struct {
	auth services.AuthService
}

func NewAuthHandler(auth services.AuthService) *authHandler {
	return &authHandler{auth: auth}
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest request.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(registerRequest)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.auth.Register(registerRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var authRequest request.AuthRequest

	err := json.NewDecoder(r.Body).Decode(&authRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(authRequest)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.auth.Login(authRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if res.ID > 0 {
		hashPwd := res.Password
		pwd := authRequest.Password

		hash := security.VerifyPassword(hashPwd, pwd)

		if hash == nil {
			token, err := security.GenerateToken(res.Email)

			if err != nil {
				response.ResponseError(w, http.StatusInternalServerError, err)
				return
			}

			response.ResponseToken(w, "Login Berhasil", token, res, http.StatusOK)
		} else {
			response.ResponseError(w, http.StatusBadRequest, errors.New("password tidak sesuai"))
			return
		}
	}
}
