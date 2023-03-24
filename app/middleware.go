package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
	"github.com/stadio-app/stadio-backend/types"
	"github.com/stadio-app/stadio-backend/utils"
)

type FuncHandler func(http.Handler) http.Handler

func (app AppBase) GothMiddleware() FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = gothic.GetContextWithProvider(r, chi.URLParam(r, "provider"))
			next.ServeHTTP(w, r)
		})
	}
}

func (app AppBase) BaseMiddleware() FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
}

func (app AppBase) AuthMiddleware() FuncHandler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth_header := r.Header.Get(types.AuthHeader)
			ctx, err := app.BearerAuthentication(r.Context(), auth_header)
			if err != nil {
				utils.ErrorResponse(w, http.StatusUnauthorized, "unauthorized")
				return
			}
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
