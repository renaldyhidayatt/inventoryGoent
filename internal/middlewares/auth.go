package middlewares

import (
	"net/http"

	"github.com/renaldyhidayatt/inventorygoent/internal/domain/response"
	"github.com/renaldyhidayatt/inventorygoent/pkg/auth"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.Authorization(r)

		if err != nil {
			response.ResponseError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		r = auth.SetContextUserId(r, token)

		next.ServeHTTP(w, r)
	})
}
