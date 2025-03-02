package middleware

import (
	"net/http"

	"github.com/mhaatha/go-rest-api/helper"
	"github.com/mhaatha/go-rest-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-KEY") == "RAHASIA" {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Data: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
