package middlewares

import (
	"errors"
	"net/http"

	"github.com/renaldyhidayatt/inventorygoent/dto/response"
	"github.com/renaldyhidayatt/inventorygoent/security"
)

func MiddlewareAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform authentication

		_, err := security.Authorization(r)

		if err != nil {
			response.ResponseError(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return

		}
		next.ServeHTTP(w, r)
	})
}
